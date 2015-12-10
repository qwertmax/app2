package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Main of app2\n")

		addrs, _ := net.InterfaceAddrs()
		for i, addr := range addrs {
			fmt.Fprintf(w, "%d %v\n", i, addr)
		}
	})
	println("ready")
	http.ListenAndServe(":9092", nil)
}
