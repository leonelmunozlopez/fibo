package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Printf("HTTP Server on <localhost:8080>")
	http.HandleFunc("/", GetNumber)
	http.ListenAndServe(":8080", nil)
}

// Data for templating
type Data struct {
	Number string
	Result string
}

// GetNumber receives a number by query param <form>
func GetNumber(w http.ResponseWriter, r *http.Request) {
	var d Data

	n := r.URL.Query().Get("n")
	if n != "" {
		d.Number = n

		p, err := strconv.Atoi(n)
		if err != nil {
			log.Printf("Error converting string <%v> to int : %v", n, err)
			d.Result = "Ingrese un número correcto"
		} else {
			// Find fibonacci number by position
			found := fibo(p)
			d.Result = fmt.Sprintf("%v", found)
		}
	}

	t := template.New("index.html")
	t, err := t.ParseFiles("index.html")
	if err != nil {
		log.Printf("Error parse <index.html> : %v", err)
		w.Write([]byte("Error rendering!\n"))
		return
	}
	t.Execute(w, d)
}

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}