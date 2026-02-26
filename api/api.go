package api

import (
	"fmt"
	"net/http"

	"github.com/Scrin/bambulab-exporter/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() <-chan error {
	http.Handle("/metrics", promhttp.Handler())

	errChan := make(chan error)
	go func() {
		addr := fmt.Sprintf(":%d", config.Port)
		errChan <- http.ListenAndServe(addr, nil)
	}()
	return errChan
}
