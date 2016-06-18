package main

import (
	"net"
	"net/rpc"
	"log"
	nrpc "github.com/muzzynine/go-tcp-notification/rpc"
)

type ConnectionRPC struct {
	config *Config
}

func (crpc *ConnectionRPC)Start() {
	err := rpc.Register(crpc)
	if err != nil {
		log.Print(err);
	}

	go listen(crpc.config.RPCAddr)
}

func listen(rpcAddr string){
	addr, err := net.ResolveTCPAddr("tcp", rpcAddr)

	if err != nil {
		log.Print("cannot resolve tcp address")
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		log.Print("RPC Listen failed")
		panic(err)
	}

	defer func(){
		if err := listener.Close(); err != nil {
			log.Print("listener close error")
		}
	}()

	log.Printf("RPC clinet listening..")
	rpc.Accept(listener)
	log.Printf("RPC accepted")
}


func (c *ConnectionRPC) GetConnections(args *nrpc.GetConnectionsArgs, ret *nrpc.GetConnectionsResp) error {
	connections := GlobalStore.GetAllConnection()

	s := []nrpc.GetConnectionsDescription{}

	for _, conn := range connections {
		log.Print(conn.GetId())
		s = append(s, nrpc.GetConnectionsDescription{ConnId : conn.GetId(), ConnIPAddr : conn.GetIPAddr()})
	}

	ret.Description = s

	return nil
}

func (c *ConnectionRPC) AddMessage(args *nrpc.AddMessageArgs, ret *int) error {
	connection, ok := GlobalStore.Get(args.ConnId)
	if !ok {
		*ret = 0
		return nil
	}

	rawConn := connection.GetConn()

	payload, err := Serialize(&SignboardMessage{command : PACKET_ADD, msgId : uint16(args.MsgId), msg : args.Msg})

	if err != nil {
		*ret = 0
		return nil
	}
		
	_, err = rawConn.Write(payload)
	
	if err != nil {
		*ret = 0
		return nil
	}
	*ret = 1
	return nil
}

func (c *ConnectionRPC) DeleteMessage(args *nrpc.DeleteMessageArgs, ret *int) error {
	connection, ok := GlobalStore.Get(args.ConnId)

	if !ok {
		*ret = 0
		return nil
	}

	rawConn := connection.GetConn()

	payload, err := Serialize(&SignboardMessage{command : PACKET_DEL, msgId : uint16(args.MsgId)})

	if err != nil {
		*ret = 0
		return nil
	}
		
	_, err = rawConn.Write(payload)
	
	if err != nil {
		*ret = 0
		return nil
	}
	*ret = 1
	return nil
}

func (c *ConnectionRPC) UpdateMessage(args *nrpc.UpdateMessageArgs, ret *int) error {
	connection, ok := GlobalStore.Get(args.ConnId)

	if !ok {
		*ret = 0
		return nil
	}

	rawConn := connection.GetConn()

	payload, err := Serialize(&SignboardMessage{command : PACKET_UPT, msgId : uint16(args.MsgId), msg : args.Msg})

	if err != nil {
		log.Print(err)
		*ret = 0
		return nil
	}
		
	_, err = rawConn.Write(payload)
	
	if err != nil {
		log.Print(err)
		*ret = 0
		return nil
	}
	*ret = 1
	return nil
}



//func (c *ConnectionRPC) getAllConnection(args *

func (c *ConnectionRPC) Ping(args int, ret *int) error {
	*ret = args
	return nil
}

	
	

