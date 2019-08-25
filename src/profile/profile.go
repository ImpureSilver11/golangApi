package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func returnProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, getProfile())
	}
}

func getProfile() string {
	raw, err := ioutil.ReadFile("./json/profile/profile.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(raw)
}
