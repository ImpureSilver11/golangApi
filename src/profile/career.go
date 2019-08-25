package profile

import (
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"github.com/labstack/echo"
)

func returnCareer() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, getCareer())
	}
}

func getCareer() string {
	raw, err := ioutil.ReadFile("./json/profile/career.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(raw)
}