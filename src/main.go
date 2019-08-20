package main

import(
  "log"
  "net/http"
  "path/filepath"
  "sync"
  "text/template"
  "fmt"
)
// "database/sql"
// _ "github.com/lib/pq"

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("tmp", t.filename)))
  })
  t.templ.Execute(w, nil)
}


// APIサーバーにする。
// TODO:フレームワークどれにするか
// DB周り
func main() {
  fmt.Println("hello  ʕ◔ϖ◔ʔ")
  http.Handle("/", &templateHandler{filename: "chat.html"})

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
}