package dss_datatypes_orset

import (
	"fmt"
	"github.com/sargun/dss/common"
)

var DataTypeSpec = dsscommon.DataType{
	"orset",
	map[string]dsscommon.DataTypeMethod{
		"new": dsscommon.DataTypeMethod{new},
	},
}

func init() {
	fmt.Println("Orset initializing")
	//	new(dsscommon.DatatypeMethod)
}

func new(newRq NewRq) OrSet {
	var _ = newRq
	foo := OrSet{}
	return foo
}
