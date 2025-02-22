package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/githiago-f/lb-mini/internals/algorithm"
	"github.com/githiago-f/lb-mini/internals/app"
	"github.com/githiago-f/lb-mini/internals/config"
)

func main() {
	fmt.Printf("%s\n\nThis is %s@%s made by %s\n", app.LOGO, app.NAME, app.VERSION, app.AUTHOR)

	config := config.Load()

	port := fmt.Sprintf(":%d", config.Loadbalancer.Listen)
	app.Logger.Infof("Listening at *.*.*.*%s\n", port)

	http.HandleFunc("/", forwardRequest(config))
	app.Logger.Fatal(http.ListenAndServe(port, nil))
}

type Handler func(res http.ResponseWriter, req *http.Request)

func forwardRequest(config *config.Config) Handler {
	return func(res http.ResponseWriter, req *http.Request) {
		url := algorithm.GetServer(req.RemoteAddr, config)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(res, req)
	}
}
