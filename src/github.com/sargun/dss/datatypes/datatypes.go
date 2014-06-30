package dss_datatypes

import (
	orset "github.com/sargun/dss/datatypes/orset"
	dsshttp "github.com/sargun/dss/http"
)

func Setup() {
	dsshttp.Register(orset.DataTypeSpec)
}
