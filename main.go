package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Getting APi ready
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All users Endpoint hit")
}
func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Updated Endpoint Hit")
}

//Request and Response  Handler
func handleRequests() {
	//initializing Mux Router
	myrouter := mux.NewRouter().StrictSlash(true)
	//Api Call Backs
	myrouter.HandleFunc("/users", AllUsers).Methods("GET")
	myrouter.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	myrouter.HandleFunc("/users/{name}/{email}", updateUser).Methods("UPDATE")
	myrouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	//Runnning Server And Checking for Fatal Error
	log.Fatal(http.ListenAndServe(":8081", myrouter))
}
func main() {
    fmt.Println("Go ORM Tutorial")

    // Handle Subsequent requests
    handleRequests()
}
