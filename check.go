package main

import (
    "fmt"
	hibp "github.com/mattevans/pwned-passwords"
	"os"
)

type Passwords struct {
    Password string
    Messages  map[string]string
}

func (pswd *Passwords) Validate() bool {
    pswd.Messages = make(map[string]string)
    client := hibp.NewClient()
    pwned, err := client.Pwned.Compromised(pswd.Password)
    if err != nil {
        fmt.Println("Pwned failed")
        os.Exit(1)
    }
    if pwned {
        pswd.Messages["Password"] = "Oh dear! You should avoid using that password"
    } else {
		pswd.Messages["Password"] = "Password is OK!"
	}
    return len(pswd.Messages) == 0
}
