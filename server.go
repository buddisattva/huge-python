package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/v1/:owner/:cm", getDick)

	e.Logger.Fatal(e.Start(":9000"))
}

// Dick is a struct for JSON/XML payloads responding
// Name of properties should be capitalized
type Dick struct {
	Object string `json:"object" xml:"object"`
	Owner  string `json:"owner" xml:"owner"`
	Length string `json:"length" xml:"length"`
	Width  string `json:"width" xml:"width"`
}

func getDick(context echo.Context) error {
	owner := context.Param("owner")
	cm, cmErr := strconv.ParseFloat(context.Param("cm"), 32)

	dick := &Dick{
		Object: "dick",
		Owner:  owner,
	}

	if cmErr != nil {
		dick.Length = "Bad request! Illegal input length!"
		dick.Width = "Bad request! Illegal input length!"
		return context.JSON(http.StatusBadRequest, dick)
	}

	dick.Length = fmt.Sprintf("%fcm", cm)
	dick.Width = fmt.Sprintf("%fcm", cm/4)

	return context.JSON(http.StatusAccepted, dick)
}
