package main

import (
	"fmt"

	"github.com/mignnoni/simulator-go/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[0])
}
