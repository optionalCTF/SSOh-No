package main

import (
	"flag"
	"strings"

	"github.com/google/uuid"
	"github.com/optionalCTF/AzAutoLogon/pkg/az"
)

/*
Inspiration posts:
- https://arstechnica.com/information-technology/2021/09/new-azure-active-directory-password-brute-forcing-flaw-has-no-fix/
- https://securecloud.blog/2019/12/26/reddit-thread-answer-azure-ad-autologon-endpoint/
- Error codes (https://docs.microsoft.com/en-us/azure/active-directory/develop/reference-aadsts-error-codes#aadsts-error-codes)
*/

type Flag struct {
	set   bool
	value string
}

var (
	email    Flag
	password Flag
)

func (fl *Flag) Set(x string) error {
	fl.value = x
	fl.set = true
	return nil
}

func (fl *Flag) String() string {
	return fl.value
}

func init() {
	flag.Var(&email, "email", "Example: user@domain.com")
	flag.Var(&password, "password", "Example: Password123!")
}

func main() {
	flag.Parse()

	if email.set {
		target := az.Target{
			User:   email.value,
			Domain: strings.Split(email.value, "@")[1],
			Guid:   uuid.New().String(),
		}
		az.Query(&target)
	}
	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter Target User: ")
		tarUser, _ := reader.ReadString('\n')
		// convert CRLF to LF
		tarUser = strings.Replace(tarUser, "\n", "", -1)
		tarDomain := strings.Split(tarUser, "@")
	*/

}
