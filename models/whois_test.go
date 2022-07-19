package models

import "testing"

func TestWhois(t *testing.T) {
	form := &WhoisRequestForm{
		Domain:  "http://www.baidu.com",
		OutType: "json",
	}

	whoisInfo := &WhoisInfo{
		RequestForm: *form,
	}

	err := whoisInfo.Whois()
	if err != nil {
		t.Error("err", err)
	}
	t.Logf("Text Result: %v\n", whoisInfo.TextInfo)
	t.Logf("JSON Result: %#v\n", whoisInfo.JsonInfo)
}
