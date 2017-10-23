package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func cookieHandler() {
	count =+ 1

	cookie = &http.Cookie {
		Name: "count",
		Value: strconv.Itoa(count)
		Expires:
	}


	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "You have visited %d times.", count)
}

func main() {

	http.HandleFunc("/", cookieHandler)

	http.Handle("/favicon.ico", http.NotFoundHandler())
}