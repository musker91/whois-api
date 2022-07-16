package models

import (
	"fmt"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
	"time"
	"whois-api/server"

	"github.com/forease/gotld"
)

type WhoisRequestForm struct {
	Domain   string `form:"domain" json:"domain" xml:"domain" binding:"required"`
	OutType  string `form:"type" json:"type" xml:"type"`
	Standard bool   `form:"standard" json:"standard" xml:"standard"`
}

const (
	WhoisOutTextType = "text"
	WhoisOutJsonType = "json"
)

const (
	WhoisStateQuerySuccess = iota
	WhoisStateParseFail
	WhoisStateDomainUnregistered
	WhoisStateTldNotSupport
	WhoisStateQueryFail
	WhoisStateRequestParamsError
)

type WhoisInfo struct {
	RequestForm WhoisRequestForm
	Domain      string
	Tld         string
	Server      string
	TldSupport  bool
	State       int
	TextInfo    string
	JsonInfo    map[string]interface{}
}

const WHOIS_PORT = "43"

func (info *WhoisInfo) Whois() (err error) {
	tld, domain, err := gotld.GetTld(info.RequestForm.Domain)
	if err != nil {
		info.State = WhoisStateRequestParamsError
		return
	}
	info.Tld = tld.Tld
	info.Domain = domain
	info.Server, info.TldSupport = server.GetWhoisServer(info.Tld)
	// query
	err = info.whoisQuery()
	if err != nil {
		return
	}
	// pasre result
	err = info.matchWohis()
	if err != nil {
		return
	}
	// convert result type
	if info.RequestForm.OutType == WhoisOutJsonType {
		err = info.textInfoToJson()
		if err != nil {
			return
		}
	}
	return nil
}

func (info *WhoisInfo) whoisQuery() (err error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(info.Server, WHOIS_PORT), time.Second*30)
	if err != nil {
		return
	}
	defer conn.Close()

	conn.Write([]byte(info.Domain + "\r\n"))
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))

	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		if info.TldSupport {
			info.State = WhoisStateQueryFail
		} else {
			info.State = WhoisStateTldNotSupport
		}
		return
	}

	info.TextInfo = string(buffer)

	return
}

func (info *WhoisInfo) matchWohis() (err error) {
	textInfoSlice := strings.Split(info.TextInfo, "\n")
	if len(textInfoSlice) <= 1 {
		textInfoSlice = strings.Split(info.TextInfo, "\r\n")
	}
	newTextInfo := ""
	for _, line := range textInfoSlice {
		newTextInfo += line + "\n"
	}
	info.TextInfo = newTextInfo

	// check query is success
	matchedUpper, err1 := regexp.Match(fmt.Sprintf("Domain Name: %s",
		strings.ToUpper(info.Domain)), []byte(info.TextInfo))
	matchedLower, err2 := regexp.Match(fmt.Sprintf("Domain Name: %s",
		strings.ToLower(info.Domain)), []byte(info.TextInfo))
	if err1 != nil || err2 != nil {
		info.State = WhoisStateParseFail
		if err1 != nil {
			return err1
		} else {
			return err2
		}
	}
	if matchedUpper || matchedLower {
		info.State = WhoisStateQuerySuccess
		return nil
	} else if info.TldSupport {
		info.State = WhoisStateDomainUnregistered
	} else {
		info.State = WhoisStateTldNotSupport
	}
	return
}

func (info *WhoisInfo) textInfoToJson() (err error) {
	keyCount := make(map[string]int)
	textInfoSlice := strings.Split(info.TextInfo, "\n")
	for _, line := range textInfoSlice {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "DNSSEC") {
			break
		}
		lineSlice := strings.Split(line, ":")
		key := lineSlice[0]
		c, ok := keyCount[key]
		if ok {
			keyCount[key] = c + 1
		} else {
			keyCount[key] = 1
		}
	}
	whoisJsonInfo := make(map[string]interface{})
	for key, count := range keyCount {
		keySlice := strings.Split(key, " ")
		newKey := strings.ToLower(keySlice[0])
		for _, ks := range keySlice[1:] {
			newKey += ks
		}

		if count > 1 {
			//`Domain Name: (.+)`
			whoisJsonInfo[newKey] = matchText(fmt.Sprintf("%s: (.+)", key), info.TextInfo, "slice")
		} else {
			whoisJsonInfo[newKey] = matchText(fmt.Sprintf("%s: (.+)", key), info.TextInfo, "str")
		}
	}
	info.JsonInfo = whoisJsonInfo
	err = info.parseJsonInfo()
	if err != nil {
		return
	}
	return
}

func (info *WhoisInfo) parseJsonInfo() (err error) {
	standardJsonData := make(map[string]interface{})

	// Assemble standard fixed field output
	standardJsonData["domainName"] = info.JsonInfo["domainName"]
	standardJsonData["domainStatus"] = info.JsonInfo["domainStatus"]
	standardJsonData["dnsNameServer"] = info.JsonInfo["nameServer"]

	if registrationTime, ok := info.JsonInfo["registrationTime"]; ok {
		standardJsonData["registrationTime"] = registrationTime
	} else if registrationTime2, ok := info.JsonInfo["creationDate"]; ok {
		standardJsonData["registrationTime"] = registrationTime2
	} else {
		standardJsonData["registrationTime"] = ""
	}

	if registryExpiryDate, ok := info.JsonInfo["registryExpiryDate"]; ok {
		standardJsonData["expirationTime"] = registryExpiryDate
	} else if registryExpiryDate2, ok := info.JsonInfo["expirationTime"]; ok {
		standardJsonData["expirationTime"] = registryExpiryDate2
	} else {
		standardJsonData["expirationTime"] = ""
	}

	if updatedDate, ok := info.JsonInfo["updatedDate"]; ok {
		standardJsonData["updatedDate"] = updatedDate
	} else {
		standardJsonData["updatedDate"] = ""
	}

	if registrarWhoisServer, ok := info.JsonInfo["registrarWHOISServer"]; ok {
		standardJsonData["registrarWHOISServer"] = registrarWhoisServer
	} else {
		standardJsonData["registrarWHOISServer"] = ""
	}

	if registrar, ok := info.JsonInfo["registrar"]; ok {
		standardJsonData["registrar"] = registrar
	} else if registrar2, ok := info.JsonInfo["sponsoringRegistrar"]; ok {
		standardJsonData["registrar"] = registrar2
	} else {
		standardJsonData["registrar"] = ""
	}

	if registrant, ok := info.JsonInfo["registrant"]; ok {
		standardJsonData["registrant"] = registrant
	} else if registrant2, ok := info.JsonInfo["registrantOrganization"]; ok {
		standardJsonData["registrant"] = registrant2
	} else {
		standardJsonData["registrant"] = ""
	}

	if contactEmail, ok := info.JsonInfo["registrarAbuseContactEmail"]; ok {
		standardJsonData["contactEmail"] = contactEmail
	} else if contactEmail2, ok := info.JsonInfo["registrantContactEmail"]; ok {
		standardJsonData["contactEmail"] = contactEmail2
	} else {
		standardJsonData["contactEmail"] = ""
	}

	if contactPhone, ok := info.JsonInfo["registrarAbuseContactPhone"]; ok {
		standardJsonData["contactPhone"] = contactPhone
	} else if contactPhone2, ok := info.JsonInfo["registrantContactPhone"]; ok {
		standardJsonData["contactPhone"] = contactPhone2
	} else {
		standardJsonData["contactPhone"] = ""
	}

	info.JsonInfo = standardJsonData
	return
}

func matchText(pattern string, text string, ty string) (data interface{}) {
	re := regexp.MustCompile(pattern)
	submatch := re.FindAllStringSubmatch(text, -1)
	if len(submatch) == 0 {
		if ty == "str" {
			data = ""
		} else {
			data = make([]string, 0)
		}
	} else {
		if ty == "str" {
			data = strings.TrimSpace(submatch[0][1])
		} else {
			rslice := make([]string, 0, len(submatch))
			for _, match := range submatch {
				rslice = append(rslice, strings.TrimSpace(match[1]))
			}
			data = rslice
		}
	}
	return
}
