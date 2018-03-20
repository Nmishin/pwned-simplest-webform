package main

import (
    "fmt"
	hibp "github.com/mattevans/pwned-passwords"
	"os"
)

type Passwords struct {
    Password string
    Errors  map[string]string
}

func (pswd *Passwords) Validate() bool {
    pswd.Errors = make(map[string]string)
    client := hibp.NewClient()
    pwned, err := client.Pwned.Compromised(pswd.Password)
    if err != nil {
        fmt.Println("Pwned failed")
        os.Exit(1)
    }
    if pwned {
        pswd.Errors["Password"] = "Oh dear! You should avoid using that password"
    }
    return len(pswd.Errors) == 0
}
