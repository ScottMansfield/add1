package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shopspring/decimal"
)

const port = 7001

var one = decimal.NewFromFloat(1)

func main() {
	http.HandleFunc("/", add1)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func add1(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		rw.Header().Add("Allow", "GET")
		http.Error(rw, "Method Not Allowed", 405)
		return
	}

	in := req.URL.Query().Get("num")

	if strings.TrimSpace(in) == "" {
		http.Error(rw, "Bad Request", 400)
		return
	}

	if num, err := decimal.NewFromString(in); err == nil {
		fmt.Fprint(rw, num.Add(one).String())
		return
	}

	// all else fails, it was a bad request
	http.Error(rw, "Bad Request", 400)
}
