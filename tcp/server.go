package main

import (
	"net"
	"log"
	"time"
)

type TCPServer struct {
	config *Config
	protocol *SignboardProtocol
	listener *net.TCPListener
}

func (server *TCPServer) Start() error{
	addr, err := net.ResolveTCPAddr("tcp", server.config.TCPAddr)
	if err != nil {
		log.Print("cannot resolve tcp address")
		return err;
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Print("cannot listen tcp establish")
		return err;
	}

	server.listener = listener

	go accept(listener, server.config, server.protocol)

	return nil
}

func (server *TCPServer) Close() error{
	log.Printf("tcp : %s close", server.config.TCPAddr)
	if err := server.listener.Close() ; err != nil {
		log.Print("close failed")
		return err
	}
	return nil
}


func accept(listener *net.TCPListener, config *Config, protocol *SignboardProtocol){
	for{
		conn, err := listener.AcceptTCP()

		if err != nil {
			log.Print("accept error")
			continue
		}

		if err := conn.SetReadBuffer(config.RcvBufferSize); err != nil {
			log.Print("set read buffer error")
			conn.Close()
			continue
		}

		if err := conn.SetWriteBuffer(config.WrtBufferSize); err != nil {
			log.Print("set write buffer error")
			conn.Close()
			continue
		}

		if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
			log.Print("set read deadline error")
			conn.Close()
			continue
		}

		message, err := protocol.Read(conn);

		if err != nil {
			log.Print("init read failed")
			conn.Close()
			continue
		}

		if message.getCommand() != PACKET_SUB {
			log.Print("no subscribe command")
			conn.Close()
			continue
		}

		log.Print("Received subscribe message")

		if ok := DB.nodeAuth(message.getToken()); !ok {
			log.Print("auth failed")
			conn.Close()
			continue
		}


		config.ConnId = message.getToken()
		config.HeartbeatTimeout = message.getHeartbeatTime()

		if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
			log.Print("set read deadline error")
			conn.Close()
			continue
		}


		payload, err := Serialize(&SignboardMessage{command : PACKET_SUB})

		if err != nil {
			log.Print("init message serialize failed")
			conn.Close()
			continue
		}

		if _, err := conn.Write(payload); err != nil {
			log.Print("subscribe response failed")
			conn.Close()
			continue
		}

		log.Printf("subscribe message writed")

		go handleConn(conn, protocol, config);
	}

}


func handleConn(conn *net.TCPConn, protocol *SignboardProtocol, config *Config){
	addr := conn.RemoteAddr().String()
	log.Printf("handle connection %s", addr)

	if _, ok := GlobalStore.Get(config.ConnId) ; !ok {
		GlobalStore.Add(config.ConnId, conn)
	}

	defer func(){
		GlobalStore.Delete(config.ConnId)
	}()

	
	for {
		if message, err := protocol.Read(conn) ; err != nil {
			return
		} else {
			switch message.getCommand() {

			case PACKET_HB :
				heartbeatHandle(conn, message, config)
				break

			case PACKET_GETALL :
				getAllMessageHandle(conn, message)
				break

				default :
			}
		}
	}
}

func subscribeHandle(conn *net.TCPConn, msg *SignboardMessage){
	/*
	if msg.getCount() < 2 {
		log.Fatalf("wrong argument")
		conn.Write([]byte("-p\r\n"))
		return
	}
*/

	payload, err := Serialize(&SignboardMessage{command : PACKET_SUB})

	if err != nil {
		log.Print("serialize failed")
		return
	}

	if _, err := conn.Write(payload); err != nil {
		log.Print("write subscribe failed")
		return
	}

}

func getAllMessageHandle(conn *net.TCPConn, msg *SignboardMessage){
	connId := msg.getToken()

	messages, err := DB.getNodeMessages(connId);

	if err != nil {
		log.Print("DB get failed");
		return
	}
	
	payload, err := Serialize(&SignboardMessage{command : PACKET_GETALL, messages : messages})

	if err != nil {
		log.Print("serialize failed")
		return
	}

	if _, err := conn.Write(payload); err != nil {
		log.Print("write subscribe failed")
		return
	}

}


func heartbeatHandle(conn *net.TCPConn, msg *SignboardMessage, config *Config){

	if err := conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(config.HeartbeatTimeout))); err != nil {
		log.Print("Set read deadline failed")
		return
	}

	payload, err := Serialize(&SignboardMessage{command : PACKET_HB})
	
	if err != nil {
		log.Print("init message serialize failed")
		return
	}

	if _, err := conn.Write(payload); err != nil {
		log.Print("subscribe response failed")
		return
	}

	log.Printf("subscribe message writed")
	
}


		
		

		
	
	

	
	
