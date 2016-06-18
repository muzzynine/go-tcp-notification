package main

import(
	"log"
	"io/ioutil"
)

type SignboardController struct {
	crtChan chan *SignboardMessage
	messageStore map[uint16]string
}

func InitController() *SignboardController{
	contorller := &SignboardController{crtChan : make(chan *SignboardMessage), messageStore : make(map[uint16]string)}
	return contorller;
}

func (controller *SignboardController) SaveMessage(msgId uint16, msg string){
	controller.messageStore[msgId] = msg
}

func (controller *SignboardController) RemoveMessage(msgId uint16){
	delete(controller.messageStore, msgId)
}

func (controller *SignboardController) PrintOutDevice(){
	message := ""
	for _, value := range controller.messageStore {
		message += value
		message += "            "
	}

	d1 := []byte(message)
	log.Print(message)
	err := ioutil.WriteFile("/dev/adafruit_r160", d1, 0644);
	if err != nil {
		log.Print(err)
	}
}

func (controller *SignboardController) GetChannel() chan *SignboardMessage {
	return controller.crtChan
}

func (controller *SignboardController) StartJob(){
	for{
		select {
		case msg := <- controller.crtChan :
			if(msg.getCommand() == PACKET_GETALL){
				for _, msg := range msg.getAllMessage() {
					controller.SaveMessage(msg.ID, msg.Message)
				}
				controller.PrintOutDevice()
			} else if(msg.getCommand() == PACKET_ADD){
				controller.SaveMessage(msg.getMessageID(), msg.getMessage())
				controller.PrintOutDevice()
			} else if(msg.getCommand() == PACKET_DEL){
				controller.RemoveMessage(msg.getMessageID())
				controller.PrintOutDevice();
			} else if(msg.getCommand() == PACKET_UPT){
				controller.SaveMessage(msg.getMessageID(), msg.getMessage())
				controller.PrintOutDevice();
				
			} else {
			}
		}
	}
}
