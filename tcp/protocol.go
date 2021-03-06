package main

import (
	"bytes"
	"net"
	"encoding/binary"
	"errors"
)

var (
	ret = []byte("\r\n")
	PACKET_GETALL = "getAll"
	PACKET_SUB = "subscribe"
	PACKET_HB = "heartbeat"
	PACKET_ADD = "add"
	PACKET_DEL = "delete"
	PACKET_UPT = "update"
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

func (msg *SignboardMessage) getCommand() string{
	return msg.command
}

func (msg *SignboardMessage) getToken() string{
	return msg.token
}

func (msg *SignboardMessage) getHeartbeatTime() uint16{
	return msg.heartbeat
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
					token := string(buffer.Next(36));
					hbPeriod := binary.BigEndian.Uint16(buffer.Next(16))
					return &SignboardMessage{
						command : PACKET_SUB,
						token : token,
						heartbeat : hbPeriod,
					}, nil
					break

				case "z" :
					token := string(buffer.Next(36));

					return &SignboardMessage{
						command : PACKET_GETALL,
						token : token,
					}, nil
					break;

				case "h" :
					return &SignboardMessage{
						command : PACKET_HB,
						}, nil
					break
					
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
		buffer.Write([]byte("h"))
		break
	case PACKET_SUB :
		buffer.Write([]byte("s"))
		break
	case PACKET_GETALL :
		buffer.Write([]byte("z"))

		msgCount := make([]byte, 16)
		binary.BigEndian.PutUint16(msgCount, uint16(len(msg.messages)))

		buffer.Write(msgCount)

		for _, msg := range msg.messages {
			msgId := make([]byte, 16)
			binary.BigEndian.PutUint16(msgId, uint16(msg.ID))

			buffer.Write(msgId)

			msgValue := []byte(msg.Message)

			msgLength := make([]byte, 16)
			binary.BigEndian.PutUint16(msgLength, uint16(binary.Size(msgValue)))

			buffer.Write(msgLength)
			buffer.Write(msgValue)

		}
		break;
		
	case PACKET_ADD :
		buffer.Write([]byte("a"))

		msgId := make([]byte, 16)
		binary.BigEndian.PutUint16(msgId, msg.msgId)

		buffer.Write(msgId)

		msgValue := []byte(msg.msg)

		msgLength := make([]byte, 16)
		binary.BigEndian.PutUint16(msgLength, uint16(binary.Size(msgValue)))

		buffer.Write(msgLength)
		buffer.Write(msgValue)
		break
	case PACKET_DEL :
		buffer.Write([]byte("d"))

		msgId := make([]byte, 16)
		binary.BigEndian.PutUint16(msgId, msg.msgId)

		buffer.Write(msgId)
		break
	case PACKET_UPT :
		buffer.Write([]byte("u"))

		msgId := make([]byte, 16)
		binary.BigEndian.PutUint16(msgId, msg.msgId)

		buffer.Write(msgId)

		msgValue := []byte(msg.msg)

		msgLength := make([]byte, 16)
		binary.BigEndian.PutUint16(msgLength, uint16(binary.Size(msgValue)))

		buffer.Write(msgLength)
		buffer.Write(msgValue)
		break
		
		default :
		return nil, errors.New("Unsupported operation")
	}
	
	buffer.Write(ret)

	return buffer.Bytes(), nil
}
				
						

				

				
