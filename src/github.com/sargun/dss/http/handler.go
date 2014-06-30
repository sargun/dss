package dsshttp

import (
	"fmt"
	"github.com/sargun/dss/common"
)

var dataTypeMap = make(map[string]dsscommon.DataType)

func Register(dataType dsscommon.DataType) {
	dataTypeMap[dataType.Name] = dataType
}
func Start() {
	fmt.Println("Starting DSS HTTPd")
	for key, value := range dataTypeMap {
		fmt.Printf("%s -> %s\n", key, value)
	}
}
