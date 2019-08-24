package main

import(
  "log"
  "net/http"
  "path/filepath"
  "sync"
  "text/template"
  "fmt"
  "io/ioutil"
  "os"
  "database/sql"
  _ "github.com/lib/pq"
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


func main() {
  // DB周り
  db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres_user password=password dbname=postgres_db sslmode=disable")
  if err != nil {
      fmt.Println(err)
  }
  // INSERT
  rows, err := db.Query("SELECT * FROM users")
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(rows)
  // JSONファイル取り出し
  // raw, err := ioutil.ReadFile("./json/profile/career.json")
  // raw, err := ioutil.ReadFile("./json/profile/skil.json")
  raw, err := ioutil.ReadFile("./json/profile/profile.json")
  if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
  }else{
    fmt.Println(raw)
  }
  defer db.Close()


// APIサーバーにする。
  http.Handle("/", &templateHandler{filename: "chat.html"})

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
}