package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func returnSkill() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, getProfile())
	}
}

func getSkill() string {
	raw, err := ioutil.ReadFile("./json/profile/skill.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(raw)
}
