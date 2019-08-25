package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func returnCareer() echo.HandlerFunc {
	return func(c echo.Context) error {
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
