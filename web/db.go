package main

import (
	"log"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
)

var DB *DBWrapper

var (
	ErrPrimaryKeyBlank = errors.New("Primary key is blank")
)


type Node struct {
	ID string `gorm:"primary_key"`
	Name string
}

type Message struct {
	ID uint `gorm:"primary_key"`
	Message string
	NodeID  string
}

type DBWrapper struct {
	DB *gorm.DB
}


func BindDB(addr string, id string, passwd string, dbName string) error{
	//	connect := id+":"+passwd+"@tcp("+addr+")/"+dbName
	connect := "root:root@tcp(localhost:8889)/signboard"
	log.Print(connect)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		return err
	}

	DB = &DBWrapper{DB : db};

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Node{}, &Message{})
	return nil
}

func (db *DBWrapper) createNode(name string) (Node, error){
	log.Print("create node " + name)
	node := Node{ID : uuid.NewV4().String(), Name : name}

	if err := db.DB.Create(&node).Error; err != nil {
		return Node{}, err
	}

	return node, nil
}

func (db *DBWrapper) getAllNode() ([]Node, error){
	var nodes []Node

	if err := db.DB.Find(&nodes).Error; err != nil {
		return nil, err
	}

	return nodes, nil
}

func (db *DBWrapper) getNodeMessages(connId string) ([]Message, error){
	var messages []Message

	node := Node{ID : connId}

	if err := db.DB.Model(&node).Related(&messages).Error; err != nil {
		return []Message{}, err
	}

	return messages, nil
}


func (db *DBWrapper) createMessage(connId string, message string) (Message, error){
	new := Message{Message : message, NodeID : connId}

	if err := db.DB.Create(&new).Error; err != nil {
		return Message{}, err
	}

	return new, nil
}

func (db *DBWrapper) deleteMessage(msgId uint) error {
	message := Message{ID : msgId}

	if err := db.DB.Delete(&message).Error; err != nil {
		return err
	}

	return nil
}

func (db *DBWrapper) updateMessage(msgId uint, msg string) (Message, error){
	var old Message

	if err := db.DB.Where("id = ?", msgId).First(&old).Error; err != nil {
		return Message{}, err
	}

	old.Message = msg

	if err := db.DB.Save(&old).Error; err != nil {
		return Message{}, err
	}

	return old, nil
}


	

	
