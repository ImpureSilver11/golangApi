package main

import(
  "log"
  "net/http"
  "path/filepath"
  "sync"
  "text/template"
  "fmt"
  "io/ioutil"
  // "os"
  // // "database/sql"
  // _ "github.com/lib/pq"
)

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
  
  raw, err := ioutil.ReadFile("./json/profile.json")
  if err != nil {
      fmt.Println(err.Error())
      // os.Exit(1)
  }else{
    fmt.Println(raw)
  }

  http.Handle("/", &templateHandler{filename: "chat.html"})

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
}