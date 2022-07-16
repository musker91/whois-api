package server

import (
	"fmt"
	"net"
)

func GetWhoisServer(tld string) (tldServer string, support bool) {
	value, ok := server[fmt.Sprintf(".%s", tld)]
	if ok != true {
		ip := net.ParseIP(tld)
		if ip != nil {
			return IP_WHOIS_SERVER, true
		} else {
			return fmt.Sprintf("%s.%s", tld, DEFAULT_WHOIS_SERVER), false
		}
	}
	return value, true
}
