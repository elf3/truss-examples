package handlers

import (
	"core/echo-service/svc"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func InterruptHandler(errc chan<- error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)

	// Place whatever shutdown handling you want here
	if Register != nil{
		Register.Deregister()
	}

	errc <- terminateError
}

func SetConfig(cfg svc.Config) svc.Config {
	globalConfig = cfg
	return cfg
}
