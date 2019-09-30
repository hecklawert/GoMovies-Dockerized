/*
*    @author      Hëck Lawert
*    @githuh      https://github.com/hecklawert
*    @date        11/09/2019
*    @description GoMovies is an API-REST developed in GoLang+MongoDB.
*				  For more information please see the README.md
 */

package main

import (
	"log"
	"net/http"
)

const Port = ":8080" // Port server

func main() {
	// Define our router object
	router := NewRouter()

	// Print log
	log.Printf("[*] Port: %d\n", Port)
        log.Println("[*] Status: Succes")

	// Getting up the server
	server := http.ListenAndServe(Port, router)
	log.Fatal(server)
}
