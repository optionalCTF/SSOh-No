package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/google/uuid"
	"github.com/optionalCTF/AzAutoLogon/pkg/az"
)

/*
Inspiration posts:
- https://arstechnica.com/information-technology/2021/09/new-azure-active-directory-password-brute-forcing-flaw-has-no-fix/
- https://securecloud.blog/2019/12/26/reddit-thread-answer-azure-ad-autologon-endpoint/
- Error codes (https://docs.microsoft.com/en-us/azure/active-directory/develop/reference-aadsts-error-codes#aadsts-error-codes)
*/

func init() {

	parser := argparse.NewParser("SSOh-No", "Enumerate and abuse a sub-par Azure SSO endpoint.")

	email := parser.String("e", "email", &argparse.Options{Required: false, Help: "Email address to query. Example: user@domain.com"})
	password := parser.String("p", "password", &argparse.Options{Required: false, Help: "Password to spray. Example: Password123!"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *email != "" && *password != "" {
		target := az.Target{
			User:     *email,
			Domain:   strings.Split(*email, "@")[1],
			Guid:     uuid.New().String(),
			Password: *password,
		}
		az.Query(&target)
	} else if *email != "" {
		target := az.Target{
			User:     *email,
			Domain:   strings.Split(*email, "@")[1],
			Guid:     uuid.New().String(),
			Password: *password,
		}
		az.Query(&target)
	} else {
		fmt.Print(parser.Usage(err))
	}
}

func main() {

}
