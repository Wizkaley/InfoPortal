package main


import (
	"github.ibm.com/dash/dash_utils/dashtest"
	"testing"
	"time"
)




func TestMain(m *testing.M){
	go main()
	dashtest.ControlCoverage(m)
	time.Sleep(1*time.Second)
}