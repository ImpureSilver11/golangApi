package users

import (
	"fmt"
	"net/http"
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func returnUsers() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, getUsers())
	}
}


func getUsers() string {
	// DB周り
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres_user password=password dbname=postgres_db sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// SELECT
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
	return string("hoge")
}