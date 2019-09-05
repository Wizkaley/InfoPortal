package main


import (
	"github.ibm.com/dash/dash_utils/dashtest"
	"testing"
)



func TestMain(m *testing.M){
	main()
	dashtest.ControlCoverage(m)
}