package main

import (
	"net/http"
	"github.com/labstack/echo"
	"strconv"
	"fmt"
)

func main() {
	e := echo.New()

	e.GET("/v1/:owner/:cm", getDick)

	e.Logger.Fatal(e.Start(":9000"))
}

type Dick struct {
	Object string `json:"object" xml:"object"`
	Owner string `json:"owner" xml:"owner"`
	Length string `json:"length" xml:"length"`
	Width string `json:"width" xml:"width"`
}

func getDick(context echo.Context) error {
	owner := context.Param("owner")
	cm, cmErr := strconv.ParseFloat(context.Param("cm"), 32)

	if cmErr != nil {
		dick := &Dick{
			Object: "dick",
                	Owner: owner,
	                Length: "Bad request!",
                	Width: "Bad request!",
        	}

		return context.JSON(http.StatusBadRequest, dick)
	}

	dick := &Dick{
		Object: "dick",
                Owner: owner,
		Length: fmt.Sprintf("%fcm", cm),
		Width: fmt.Sprintf("%fcm", cm/4),
	}

	return context.JSON(http.StatusAccepted, dick)
}

