package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mtmoses/httprouter"
)

func getclaim(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := mux.Vars(r)
	if params.ID == homeId {
		getClaimsByHomeID(dbConn*sql.DB, homeID)
	} else if params.ID == warrantyID {
		getClaimsByWarrantyID(dbConn*sql.DB, warrantyID)
	}
}
func main() {
	router := mux.NewRouter()

	router.GET("/claim/{id}", getclaim)

	log.Fatal(http.ListenAndServe(":8000", router))
}
