package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/get", getEmpInfo).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/post", postEmpInfo).Methods(http.MethodPost, http.MethodOptions)
	//router.Use(mux.CORSMethodMiddleware(router))
	return router
}

func enbaleCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Headers", "authentication, content-type")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
}

func getEmpInfo(w http.ResponseWriter, r *http.Request) {
	enbaleCORS(&w)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("get info from backend")
}

func postEmpInfo(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(w).Header().Set("Access-Control-Allow-Headers", "authentication, content-type")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Add("Content-Type", "application/json")
	fmt.Println(r.GetBody())
	json.NewEncoder(w).Encode("posts")
}
