package dps

// DPS main

import (
	logger "github.com/ShaoWenAcerLu/BCIR_main/utilities"
	"github.com/ShaoWenAcerLu/BCIR_policyManager"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	// initialize API Map
	for key, val := range policyManager.PolicyApiMap {
		DPSAPIMap[key] = val
	}
}

func StartDPS() {
	logger.LogInfo("DPS service starts...")
	router := mux.NewRouter()

	// policy
	router.HandleFunc(DPSAPIMap["CreatePolicy"], CreatePolicy).Methods("POST")
	router.HandleFunc(DPSAPIMap["UpdatePolicy"], UpdatePolicy).Methods("PUT")
	router.HandleFunc(DPSAPIMap["GetPolicy"], GetPolicy).Methods("GET")
	router.HandleFunc(DPSAPIMap["GetPolicies"], GetPolicies).Methods("GET")
	router.PathPrefix("/").Subrouter().StrictSlash(true).Path(DPSAPIMap["GetPolicies"]).Queries("status", "{key}").HandlerFunc(GetPolicies).Methods("GET")
	router.HandleFunc(DPSAPIMap["DeletePolicy"], DeletePolicy).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
