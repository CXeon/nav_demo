package http

import (
	"context"
	"fmt"
	"github.com/CXeon/nav_demo/config"
	"net/http"
	"time"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 5 * time.Second
)

var httpServer *http.Server

func Start(conf *config.Config) error {
	addr := fmt.Sprintf("%s:%d", conf.Application.Host, conf.Application.Port)
	httpServer = &http.Server{
		Addr:         addr,
		Handler:      initRoutes(conf),
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	return httpServer.ListenAndServe()
}

func Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
	defer cancel()
	return httpServer.Shutdown(ctx)
}
