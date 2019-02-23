package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	log.Printf("HTTP Server on <localhost:8080>")
	http.HandleFunc("/", GetNumber)

	// Open browser!
	go open("http://localhost:8080/")
	panic(http.ListenAndServe(":8080", nil))
}

// Data for templating
type Data struct {
	Number string
	Result string
}

// GetNumber receives a number by query param <form>
func GetNumber(w http.ResponseWriter, r *http.Request) {
	var d Data

	t := template.New("index.html")
	t, err := t.ParseFiles("index.html")
	if err != nil {
		log.Printf("Error parse <index.html> : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error rendering!\n"))
		return
	}

	n := r.URL.Query().Get("n")
	if n != "" {
		d.Number = n

		p, err := strconv.Atoi(n)
		if err != nil {
			log.Printf("Error converting string <%v> to int : %v", n, err)
			d.Result = "Ingrese un número correcto"

			w.WriteHeader(http.StatusBadRequest)
			t.Execute(w, d)
			return

		}

		// Find fibonacci number by position
		log.Printf("Checking for <%v> position...", p)
		found := fibo(p - 1)

		log.Printf("Results : %v", found)
		d.Result = fmt.Sprintf("%v", found)
	}

	w.WriteHeader(http.StatusOK)
	t.Execute(w, d)
}

// remember : to include first place "0", call function with (n-1)
func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

// open, auto-launch Browser
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}

	case "darwin":
		cmd = "open"

	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
