package main

import (
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
)

var accessCount = new(int64)

func index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	access := atomic.AddInt64(accessCount, 1)
	data := struct {
		Now         time.Time
		OS          string
		AccessCount int
	}{
		Now:         time.Now(),
		OS:          runtime.GOOS,
		AccessCount: int(access),
	}

	if err := getTemplate(ctx, "index").Execute(w, data); err != nil {
		panic(err)
	}
}
