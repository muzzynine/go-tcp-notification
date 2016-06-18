package main

import (
	"log"
)

func main(){

	config := InitConfig()
	log.Printf("Initialize client config")

	controller := InitController();
	go controller.StartJob()

	if err := InitTCP(config, controller.GetChannel()) ; err != nil {
		log.Print(err)
		panic(err)
	}

	sigWatcher := InitSignal()
	sigWatcher.HandleSignal()
}

	
