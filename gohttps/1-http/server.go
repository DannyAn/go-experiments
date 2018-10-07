package main

import (
	"fmt"
	 http "net/http"
)

func handler(w  http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
}

func main() {
	 http.HandleFunc("/", handler)
	 err := http.ListenAndServe(":8080", nil)
	 if err!=nil {
		fmt.Println("Error:",err);
	}

}
