package main

import (
	"fmt"
	"net/http"
	"time"
)

func RequestDurationLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		// Log it in format: date time method url from ip:port - duration: dur

		fmt.Printf(

			"\n%s %s %s from %s - duration: %s\n",
			start.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
