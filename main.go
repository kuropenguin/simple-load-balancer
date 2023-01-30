package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backend []*Backend
	current uint64
}

func main() {
	u, _ := url.Parse("http://localhost:8080")

	rp := httputil.NewSingleHostReverseProxy(u)

	http.HandlerFunc(rp.ServeHTTP)
}

func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}
