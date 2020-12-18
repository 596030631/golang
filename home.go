package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello")
	Log.Println("ok")
}

func Home() {
	http.HandleFunc("/", HomeHandler)
	e := http.ListenAndServe("127.0.0.1:7004", nil)
	if e != nil {
		Log.Fatalln("Start Error:", e)
	}
}
