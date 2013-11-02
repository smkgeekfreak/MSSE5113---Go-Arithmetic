package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"skircher/WikiArith"
	"strings"
)

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

type ArithRequest struct {
	ArithMethod string
	ArithArgs   string
}

func doJSONArithmeticPOST(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		fmt.Fprintf(w, "Couldn't read request body: %s", err)
	} else {
		fmt.Println("Post req", body)
		dec := json.NewDecoder(strings.NewReader(string(body)))
		fmt.Println("Post dec", dec)
		var m ArithRequest
		if err := dec.Decode(&m); err != nil {
			fmt.Fprintf(w, "Couldn't decode JSON: %s", err)
		} else {

			//fmt.Println("Parsed error", args, err2)
			w.Header().Set("Content-Type", "application/json")
			//var args = 0.0
			args, _ := parseArgsF64(m.ArithArgs)
			var result float64
			switch m.ArithMethod {
			case "addition":
				//args, _ := parseArgsInt(m.ArithArgs)
				result = float64(WikiArith.Add(args))
			case "subtraction":
				//args, _ := parseArgsInt(m.ArithArgs)
				result = float64(WikiArith.Subtract(args))
			case "multiplication":
				//args, _ := parseArgsInt(m.ArithArgs)
				result = float64(WikiArith.Multiply(args))
			case "division":
				//args, _ := parseArgsF64(m.ArithArgs)
				result = float64(WikiArith.Divide(args))
			default:
			}

			fmt.Fprint(w, Response{"ArithMethod": m.ArithMethod,
				"ArithInArgs": m.ArithArgs, "ArithParsedArgs": args, "ArithResult": result})

		}
	}
}

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}
