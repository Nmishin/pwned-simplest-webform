package main

import (
  "github.com/bmizerany/pat"
  "html/template"
  "log"
  "net/http"
)

func main() {
  mux := pat.New()
  mux.Get("/", http.HandlerFunc(index))
  mux.Post("/", http.HandlerFunc(send))
  mux.Get("/confirmation", http.HandlerFunc(confirmation))

  log.Println("Listening...")
  http.ListenAndServe(":8080", mux)
}

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
    } else {
        http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
    }
}

func confirmation(w http.ResponseWriter, r *http.Request) {
  render(w, "confirmation.html", nil)
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

