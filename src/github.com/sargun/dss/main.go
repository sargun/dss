package main

import (
	"fmt"
	"github.com/sargun/dss/datatypes"
	"github.com/sargun/dss/http"
)

func main() {
	dss_datatypes.Setup()
	fmt.Println("Hello world")
	dsshttp.Start()
}
