package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/v1/:name/:cm", calcDick)

	e.Logger.Fatal(e.Start(":9000"))
}

type dick struct {
	Type   string `json:"@type" xml:"@type"`
	Length string `json:"length" xml:"length"`
	Width  string `json:"width" xml:"width"`
}

// person is a struct for JSON/XML payloads responding
// Name of properties should be capitalized
type person struct {
	Type        string `json:"@type" xml:"@type"`
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
	Dick        dick   `json:"dick" xml:"dick"`
}

func calcDick(context echo.Context) error {
	name := strings.Title(strings.ToLower(context.Param("name")))
	cm, cmErr := strconv.ParseFloat(context.Param("cm"), 32)

	person := &person{
		Type: "Person",
		Name: name,
	}

	if cmErr != nil {
		person.Description = "You don't have any dick."
		return context.JSON(http.StatusLengthRequired, person)
	}

	typeDick := "dick"
	stringLength := fmt.Sprintf("%.2f", cm)
	stringWidth := fmt.Sprintf("%.2f", cm/4)
	dick := dick{
		Type:   strings.Title(typeDick),
		Length: stringLength,
		Width:  stringWidth,
	}

	person.Description =
		name + " has a " + typeDick + " with " +
			stringLength + "cm length and " +
			stringWidth + " cm width."
	person.Dick = dick

	return context.JSON(http.StatusOK, person)
}
