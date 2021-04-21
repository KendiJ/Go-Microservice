package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main () {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello")
		d, err := ioutil.ReadAll(r.Body)
		if err !=  nil {
			http.Error(rw, "Ooopsy", http.StatusBadRequest)
			return
		}

		// log.Printf("Data %s\n", d) //Able to print data to the log

		//write back to the user.
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("See you later Self:)")
	})


	http.ListenAndServe(":9090", nil)
}