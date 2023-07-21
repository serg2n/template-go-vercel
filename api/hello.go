package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	  if err != nil {
			 fmt.Printf("server: could not read request body: %s\n", err)
	  }
	fmt.Fprintf(w, "{<!DOCTYPE html><html lang="ru"><head><meta charset="UTF-8"></head><body>}")
	fmt.Fprintf(w,string(reqBody))
	fmt.Fprintf(w, " </body></html>")
}