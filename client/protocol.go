package main

import (
	"log"
	"bytes"
	"net"
	"encoding/binary"
	"errors"
)

var (
	ret = []byte("\r\n")
	PACKET_SUB = "subscribe" // -> s
	PACKET_HB = "heartbeat" // -> h
	PACKET_ADD = "add"
	PACKET_DEL = "delete"
	PACKET_UPT = "update"
	PACKET_GETALL = "getAll"
)

type SignboardProtocol struct {
}

type SignboardMessage struct {
	command string
	token string
	heartbeat uint16
	msgId uint16
	msg string

	messages []Message
}

type Message struct {
	ID uint16
	Message string
}

func (msg *SignboardMessage) getCommand() string{
	return msg.command
}

func (msg *SignboardMessage) getToken() string{
	return msg.token
}

func (msg *SignboardMessage) getHeartbeatTime() uint16{
	return msg.heartbeat
}

func (msg *SignboardMessage) getMessageID() uint16{
	return msg.msgId
}

func (msg *SignboardMessage) getMessage() string{
	return msg.msg
}

func (msg *SignboardMessage) getAllMessage() []Message{
	return msg.messages
}



/* 
 * SignboardProtocol
 * --------------------------------------
 * | command(1) | token(8) / time(8)    |
 * --------------------------------------
 * command 
 *   - s : subscribe
 *   - h : heartbeat
 * (subscribe) token 
 * (heartbeat) time  
 */

func (proto *SignboardProtocol) Read(conn *net.TCPConn) (*SignboardMessage, error){
	buffer :=  bytes.NewBuffer([]byte{})

	for {
		temp := make([]byte, 1024)

		if length, err := conn.Read(temp) ; err != nil {
			return nil, err
		} else {
			buffer.Write(temp[:length])

			index := bytes.Index(buffer.Bytes(), ret)
			if index > -1 {
				command := string(buffer.Next(1))

				switch command {
				case "s" :
					return &SignboardMessage{
						command : PACKET_SUB,
						}, nil

				case "h" :
					return &SignboardMessage{
						command : PACKET_HB,
						}, nil

				case "z" :
					messages := []Message{}
					msgCount := binary.BigEndian.Uint16(buffer.Next(16))
					for i := 0 ; i < int(msgCount) ; i++ {
						msgId := binary.BigEndian.Uint16(buffer.Next(16))
						log.Print(msgId)
						msgLength := binary.BigEndian.Uint16(buffer.Next(16))
						msgValue := string(buffer.Next(int(msgLength)))
						messages = append(messages, Message{ID : msgId, Message : msgValue})
					}
					return &SignboardMessage{
						command : PACKET_GETALL,
						messages : messages,
					}, nil
					
					
						
						
				case "a" :
					msgId := binary.BigEndian.Uint16(buffer.Next(16))
					log.Print(msgId)
					msgLength := binary.BigEndian.Uint16(buffer.Next(16))
					msgValue := string(buffer.Next(int(msgLength)))

					return &SignboardMessage{
						command : PACKET_ADD,
						msgId : msgId,
						msg : msgValue,
					}, nil


				case "d" :
					msgId := binary.BigEndian.Uint16(buffer.Next(16))

					return &SignboardMessage{
						command : PACKET_DEL,
						msgId : msgId,
					}, nil

				case "u" :
					msgId := binary.BigEndian.Uint16(buffer.Next(16))
					msgLength := binary.BigEndian.Uint16(buffer.Next(16))
					msgValue := string(buffer.Next(int(msgLength)))

					return &SignboardMessage{
						command : PACKET_UPT,
						msgId : msgId,
						msg : msgValue,
					}, nil
					
					
					
					default :
					return nil, errors.New("Unsupported operation")
				}
			}
		}
	}
}

func Serialize(msg *SignboardMessage) ([]byte, error){
	buffer := bytes.NewBuffer([]byte{})

	switch msg.getCommand() {
	case PACKET_HB :
		//1byte
		buffer.Write([]byte("h"))
		
		break
	case PACKET_SUB :
		buffer.Write([]byte("s"))

		//16byte
		buffer.Write([]byte(msg.getToken()))
		if msg.getHeartbeatTime() >= (1 << 8) {
			return nil, errors.New("hb time overflow")
		}
		//16byte
		hbPeriod := make([]byte, 16)
		binary.BigEndian.PutUint16(hbPeriod, uint16(msg.getHeartbeatTime()))

		buffer.Write(hbPeriod);

		break
	case PACKET_GETALL :
		buffer.Write([]byte("z"))

		//16byte
		buffer.Write([]byte(msg.getToken()))
		break;

		default :
		return nil, errors.New("Unsupported operation")


	}

	
	
	buffer.Write(ret)

	return buffer.Bytes(), nil
}
				
						

				

				
