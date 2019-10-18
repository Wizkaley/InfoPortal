package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

// StudentAggregates ...
func StudentAggregates(ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			res, err := json.Marshal("Bad Request")
			if err != nil {
				log.Println("Error while Marshalling Error Message")
			}
			sendErr(w, http.StatusBadRequest, res)
		}
		if r.Body != nil {
			//params := mux.Vars(r)
			//marks, _ := strconv.Atoi(params["marks"])
			//err := dao.StudentAggregates(int32(marks), ds)

			// if err != nil {
			// 	log.Printf("Could not apply Indexes: %v", err)
			// }
		}
	})
}
