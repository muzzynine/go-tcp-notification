package main

import (
	"log"
	"net"
	"sync"
	"errors"
	"time"
)

var (
	ErrReadConn = errors.New("connection read error")
	ErrWriteConn = errors.New("connection write error")
	ErrCloseConn = errors.New("connection closed")
	ErrHeartbeatTimeout = errors.New("heartbeat timeout")
	ErrSubscribeFailed = errors.New("subscribe failed")
)


type TCPClient struct {
	config *Config
	mutex *sync.Mutex
	conn *net.TCPConn
	rcvChan chan *SignboardMessage
	sndChan chan *SignboardMessage
	hbChan chan *SignboardMessage
	crtChan chan *SignboardMessage
	errChan chan error
	protocol *SignboardProtocol
}

func InitTCP(config *Config, crtChan chan *SignboardMessage) error {
	tcpClient := &TCPClient{config : config, crtChan : crtChan}

	if err := tcpClient.Connect(); err != nil {
		return err
	}
	return nil
}

func (client *TCPClient) Connect() error {
	client.errChan = make(chan error, 2)
	client.rcvChan = make(chan *SignboardMessage)
	client.sndChan = make(chan *SignboardMessage)
	client.hbChan = make(chan *SignboardMessage)
	
	tcpAddr, err := net.ResolveTCPAddr("tcp", client.config.Addr)

	if err != nil {
		log.Print(err)
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Print(err)
		return err
	}
	
	client.conn = conn

	go client.Handle()

	return nil
}

func (client *TCPClient) Close() error {
	client.errChan <- ErrCloseConn
	
	if err := client.conn.Close(); err != nil {
		log.Print("close failed")
		return err
	}

	time.Sleep(time.Duration(10 * time.Millisecond))

	close(client.rcvChan)
	close(client.sndChan)
	close(client.errChan)
	close(client.hbChan)


	return nil
}

func (client *TCPClient) Handle() {
	go client.Read()


	if err := client.conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Print("SetReadDeadLine failed")
		client.Close()
		return
	}
	
	payload, err := Serialize(&SignboardMessage{command : PACKET_SUB, token : client.config.ConnId, heartbeat : client.config.HeartbeatTimeout})

	if err != nil {
		log.Print("subscribe serialization failed")
		client.Close()
		return
	}

	if _, err := client.conn.Write(payload); err != nil {
		log.Print("subscribe registration wrtie failed")
		client.Close()
		return
	}

	for {
		message, err := client.protocol.Read(client.conn)
		if err != nil {
			log.Print("subscribe read failed")
			client.Close()
			return
		}

		if message.getCommand() == PACKET_SUB {
			//registartion response validation
			//do

			if err := client.conn.SetReadDeadline(time.Time{}); err != nil {
				log.Print("SetReadDeadLine failed")
				log.Print(err)
				client.Close()
				return
			}

			payload, err := Serialize(&SignboardMessage{command : PACKET_GETALL, token : client.config.ConnId})

			if err != nil {
				log.Print("get all serialization failed")
				client.Close()
				return
			}

			if _, err := client.conn.Write(payload); err != nil {
				log.Print("get all registration wrtie failed")
				client.Close()
				return
			}



			go client.Read()
			go client.Heartbeat()
			break;
		}
	}

	
	for {
		select {
		case err := <- client.errChan :
			if(err == ErrHeartbeatTimeout){
				log.Print("Heartbeat timeout")
				client.Close()
				return;
			}
		case rcvMsg := <- client.rcvChan :
			if(rcvMsg.getCommand() == PACKET_HB){
				client.hbChan <- rcvMsg
			} else {
				client.crtChan <- rcvMsg
			}
		case sndMsg := <- client.sndChan :
			log.Print("message Write")
			payload, err := Serialize(sndMsg)
			if err != nil {
				log.Print("Serialize failed")
				continue
			}
			client.mutex.Lock()
			if _, err := client.conn.Write(payload); err != nil {
				client.mutex.Unlock()
				log.Print("Write failed")
				continue
			}
			client.mutex.Unlock()
		}
	}
}

func (client *TCPClient) Write(message *SignboardMessage) {
	log.Print("in Write")
	client.sndChan <- message
}


func (client *TCPClient) Read() {
	for {
		/* Non-blocking
		 * errChan으로부터 메시지가 있을 경우 처리하고, 없을 경우 Read 진행 */
		select {
		case err := <- client.errChan :
			if(err == ErrCloseConn) {
				log.Printf("stop read goroutine")
				return
			}

		default :
			if message, err := client.protocol.Read(client.conn) ; err != nil {
				log.Print(err)
				log.Print("message parse failed")
				client.errChan <- err
			} else {
				log.Print("message arrived")
				client.rcvChan <- message
			}
		}
	}
}
		
func (client *TCPClient) Heartbeat(){
	period := time.Tick(time.Duration(client.config.HeartbeatPeriod) * time.Second)
	timeout := time.After(time.Duration(client.config.HeartbeatTimeout) * time.Second)

	count := 0

	for {
		log.Print("heartbeat")
		if(count >= client.config.HeartbeatLimit){
			client.errChan <- ErrHeartbeatTimeout
			return
		}
			
		select {
		case <- period :

			payload, err := Serialize(&SignboardMessage{command : PACKET_HB})
			if err != nil {
				log.Print("Serialize failed")
				continue
			}

//			client.mutex.Lock()
			if _, err := client.conn.Write(payload); err != nil {
				log.Print("Write failed")
//				client.mutex.Unlock()
				continue
			}
//			client.mutex.Unlock()

			break;

		case <- timeout :
			count++
			timeout = time.After(time.Duration(client.config.HeartbeatTimeout) * time.Second)
			break;

		case  <- client.hbChan :
			count = 0
			timeout = time.After(time.Duration(client.config.HeartbeatTimeout) * time.Second)

		}
	}	
}


	


	
	
