package main

import (
	"fmt"
	"log"
	//"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func executeCode(w http.ResponseWriter, r *http.Request){
	// get the body of our POST request
	// return the string response
	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "%+v", string(reqBody))
	//fmt.Println("Endpoint Hit: executeCode")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(){
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/execute", executeCode).Methods("POST")

	log.Fatal(http.ListenAndServe(":7000", myRouter))
}

func main(){
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}