package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/app2", func(w http.ResponseWriter, req *http.Request) {
		addrs, _ := net.InterfaceAddrs()
		for i, addr := range addrs {
			fmt.Fprintf(w, "%d %v\n", i, addr)
		}
	})
	println("ready")
	http.ListenAndServe(":9092", nil)
}
