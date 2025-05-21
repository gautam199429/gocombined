package main

import (
	coprocessor "coprocessor/internal"
	utility "coprocessor/utility"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8085
	}
	utility.CacheInit(24*time.Hour, 24*time.Hour)
	http.HandleFunc("/entitlements", coprocessor.RequestHandler)
	log.Printf("Starting on :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
