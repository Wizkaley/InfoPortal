package utils

import (
	"testing"

	mongodal "RESTApp/utils/mongodal"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	url := "localhost:27017"

	db, _ := Init(url)
	defer db.Close()

	//assert.Equalf(t, sess, err, "Expected %v but bot %v", sess, err)
}
func TestInitFail(t *testing.T) {
	db, err := Init("lhost:893024")
	defer db.Close()
	//assert.EqualError(t, err, "Error while getting session : no reachable servers")
	//errors.New("s:no reachable servers")
	assert.Error(t, err, "Expected and Got Error")

	// func TestMain(m *testing.M) {
	// 	m.Run("testInit", testInit(t))
	// 	m.Run("testInitFail", TestInitFail(t))
}

func TestEnsureIndex(t *testing.T) {
	s, _ := Init("localhost:27017")
	defer s.Close()
	val := ensureInd(s, "testing", "test")

	var items []string
	indices, _ := val.DB("testing").C("test").Indexes()
	for _, item := range indices {
		items = append(items, item.Name)
	}
	//fmt.Println(items)
	assert.Equal(t, items[1], "studentMarks_1")
}

