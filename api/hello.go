package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	  if err != nil {
			 fmt.Printf("server: could not read request body: %s\n", err)
	  }
	  fmt.Printf("server: request body: %s\n", reqBody)

	  fmt.Fprintf(w, `{"message": "hello!"}`)
	fmt.Fprintf(w, reqBody)
}
