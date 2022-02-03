package az

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Target struct {
	User   string
	Domain string
	Guid   string
}

func Query(t *Target) {
	tar := "https://autologon.microsoftazuread-sso.com/" + t.Domain + "/winauth/trust/2005/usernamemixed?client-request-id=" + t.Guid

	var body = strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?>
	<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:a="http://www.w3.org/2005/08/addressing" xmlns:u="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	  <s:Header>
		<a:Action s:mustUnderstand="1">http://schemas.xmlsoap.org/ws/2005/02/trust/RST/Issue</a:Action>
		<a:MessageID>urn:uuid:` + uuid.New().String() + `</a:MessageID>
		<a:ReplyTo>
		  <a:Address>http://www.w3.org/2005/08/addressing/anonymous</a:Address>
		</a:ReplyTo>
		<a:To s:mustUnderstand="1">` + tar + `</a:To>
		<o:Security xmlns:o="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" s:mustUnderstand="1">
		  <u:Timestamp u:Id="_0">
			<u:Created>` + time.Now().Format(time.RFC3339Nano) + `</u:Created>
			<u:Expires>` + time.Now().Add(time.Minute*10).Format(time.RFC3339Nano) + `</u:Expires>
		  </u:Timestamp>
		  <o:UsernameToken u:Id="uuid-ec4527b8-bbb0-4cbb-88cf-abe27fe60977">
			<o:Username>` + t.User + ` </o:Username>
			<o:Password>Pword</o:Password>
		  </o:UsernameToken>
		</o:Security>
	  </s:Header>
	  <s:Body>
		<trust:RequestSecurityToken xmlns:trust="http://schemas.xmlsoap.org/ws/2005/02/trust">
		  <wsp:AppliesTo xmlns:wsp="http://schemas.xmlsoap.org/ws/2004/09/policy">
			<a:EndpointReference>
			  <a:Address>urn:federation:MicrosoftOnline</a:Address>
			</a:EndpointReference>
		  </wsp:AppliesTo>
		  <trust:KeyType>http://schemas.xmlsoap.org/ws/2005/05/identity/NoProofKey</trust:KeyType>
		  <trust:RequestType>http://schemas.xmlsoap.org/ws/2005/02/trust/Issue</trust:RequestType>
		</trust:RequestSecurityToken>
	  </s:Body>
	</s:Envelope>
	`)

	req, err := http.NewRequest("POST", tar, body)

	if err != nil {
		fmt.Printf("Something went wrong! %s", err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Something went wrong! %s", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(data))
	res.Body.Close()
}
