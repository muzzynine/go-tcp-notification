package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
)

type SignalWatcher struct {
	c chan os.Signal
}
	

func InitSignal() *SignalWatcher {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP);

	return &SignalWatcher{c : c}
}

func (watcher *SignalWatcher) HandleSignal(){
	for{
		signal := <- watcher.c
		log.Printf("Signal received : %s", signal.String())

		switch signal{
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP :
			return
		case syscall.SIGHUP :

		default:
			return
		}
	}
}

			
			
