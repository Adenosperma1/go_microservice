package main

import (
	"io"
	"log"
	"net/http"
)

func main(){


//handlers
http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
	log.Println("hi")
	d, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Printf("Data______________ %s", d)
})

http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request){
	log.Println("bye")
})
	


http.ListenAndServe(":9090", nil) // this will run in the browser localhost:9090

}

