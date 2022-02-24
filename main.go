package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/akamensky/argparse"
	"github.com/optionalCTF/SSOh-no/pkg/az"
	service "github.com/optionalCTF/SSOh-no/pkg/svc"
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
	userList := parser.String("U", "Userlist", &argparse.Options{Required: false, Help: "Specify userlist to enumerate"})

	var wg sync.WaitGroup

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *email != "" && *password != "" {
		az.Query(*email, strings.Split(*email, "@")[1], *password, &wg)
	} else if *email != "" {
		az.Query(*email, strings.Split(*email, "@")[1], "", &wg)
	} else if *userList != "" {
		users, err := service.ReadFile(*userList)
		if err != nil {
			fmt.Printf("readLines: %s", err)
		}
		wg.Add(len(users))
		for _, line := range users {
			go az.Query(line, strings.Split(line, "@")[1], "", &wg)
		}
		wg.Wait()
	} else {
		fmt.Print(parser.Usage(err))
	}
}

func main() {

}
