package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
)

/*
Inspiration posts:
- https://arstechnica.com/information-technology/2021/09/new-azure-active-directory-password-brute-forcing-flaw-has-no-fix/
- https://securecloud.blog/2019/12/26/reddit-thread-answer-azure-ad-autologon-endpoint/
- Error codes (https://docs.microsoft.com/en-us/azure/active-directory/develop/reference-aadsts-error-codes#aadsts-error-codes)
*/

type Target struct {
	guid   string
	domain string
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
	/*
		resp, err := http.Post("https://autologon.microsoftazuread-sso.com/" + t.domain + "winauth/trust/2005/usernamemixed?client-request-id=")

		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
		}
	*/
	fmt.Println("https://autologon.microsoftazuread-sso.com/" + t.domain + "/winauth/trust/2005/usernamemixed?client-request-id=" + t.guid)

	// XML Template to ensure consistency (Consider implementing this in code to avoid additional files.)
	xmlFile, err := os.Open("request.xml")

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}

	fmt.Println("Successfully opened request.xml")
	defer xmlFile.Close()
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Target Domain: ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	target := Target{
		domain: text,
		guid:   guidGeneration(),
	}

	query(&target)
}
