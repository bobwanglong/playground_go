package main

import "net/http"

type HttpHandlerDecoratie func(http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecoratie) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i]
		h = d(h)
	}
	return h
}

func main() {

}
