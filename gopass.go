package main

import (
  "github.com/bmizerany/pat"
  "github.com/kabukky/httpscerts"
  "html/template"
  "log"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  render(w, "index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
    pswd := &Passwords{
        Password: r.FormValue("password"),
    }
    if pswd.Validate() == false {
        render(w, "index.html", pswd)
    return
    }
}

func render(w http.ResponseWriter, filename string, data interface{}) {
  tmpl, err := template.ParseFiles(filename)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  if err := tmpl.Execute(w, data); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func CheckAndGenCerts() {
  err := httpscerts.Check("cert.pem", "key.pem")
  if err != nil {
    log.Println("Regenerate certificates")
    err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8443")
      if err != nil {
        log.Fatal("ERROR: Can't generate certificates")
      }
   }
}

func main() {
  CheckAndGenCerts()
  mux := pat.New()
  mux.Get("/", http.HandlerFunc(index))
  mux.Post("/", http.HandlerFunc(send))

  log.Println("Listening 127.0.0.1:8443...")
  err := http.ListenAndServeTLS("127.0.0.1:8443", "cert.pem", "key.pem", mux)
  if err != nil {
    log.Fatal(err)
  }
}
