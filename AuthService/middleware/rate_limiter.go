package middleware

import (
	"net/http"
	"time"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Every(1*time.Minute),5)

func RateLimitMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow(){
			http.Error(w,"too many request",http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w,r)
	})
}