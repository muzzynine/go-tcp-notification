package main

import (
	"log"
)

var GlobalStore *ConnectionStore

func main() {

	config := InitConfig()
	log.Printf("Initialize server config..")
	
	GlobalStore = NewConnectionStore()
	defer func(){
		GlobalStore.Close()
		log.Printf("Server terminated. all connection closed")
	}()

	if err := BindDB(config.DBAddr, config.DBUser, config.DBPasswd, config.DBName); err != nil {
		panic(err)
	}

	log.Print("DB connected");

	rpc := &ConnectionRPC{config : config}
	rpc.Start()

	log.Print("RPC started");
	
	tcp := &TCPServer{config : config, protocol : &SignboardProtocol{}}

	if err := tcp.Start() ; err != nil {
		panic(err)
	}

	log.Printf("Server ready. waiting system signal...")
	
	sigWatcher := InitSignal()
	sigWatcher.HandleSignal()
}

