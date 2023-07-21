package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	  if err != nil {
			 fmt.Printf("server: could not read request body: %s\n", err)
	  }
	fmt.Fprintf(w, `{"message": "hello!"}<br />`)
	buf2 := utf8.AppendRune(reqBody, 0x10000)
	fmt.Fprintf(w, string(buf2))
}
