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
	fmt.Fprintf(w,"%s",string(reqBody))
}