package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type visitor struct {
	limiter    *rate.Limiter
	lastSeen   time.Time
	lastLogged time.Time
}

var visitors = make(map[string]*visitor)
var mu sync.RWMutex

func init() {
	go cleanupVisitors()
}

func getVisitor(ip string) *visitor {
	mu.RLock()

	v, exists := visitors[ip]
	if !exists {
		mu.RUnlock()
		mu.Lock()
		limiter := rate.NewLimiter(1, 3)
		v = &visitor{limiter, time.Now(), time.Now().Add(-time.Minute)}
		visitors[ip] = v
		mu.Unlock()

		return v
	}

	v.lastSeen = time.Now()
	mu.RUnlock()

	return v
}

// Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 5*time.Minute {
				delete(visitors, ip)
			}
		}

		mu.Unlock()
	}
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)

			return
		}

		visitor := getVisitor(ip)

		if !visitor.limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)

			if time.Since(visitor.lastLogged) > time.Minute {
				log.Printf("too many requests from %v", ip)
				mu.Lock()
				visitor.lastLogged = time.Now()
				mu.Unlock()
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}
