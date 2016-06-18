package main

import (
	"log"
	"net/rpc"
	nrpc "github.com/muzzynine/go-tcp-notification/rpc"
)

var NRPC *nrpc.NotificationRPC

func BindRPC(addr string) error {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("dialing failed addr : %s", addr)
		return err
	}
	NRPC = &nrpc.NotificationRPC{Client : client}

	return nil
}
	

