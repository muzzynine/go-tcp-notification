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
	//connect := id+":"+passwd+"@tcp("+addr+")/"+dbName
	connect := "root:root@tcp(localhost:8889)/signboard"
	log.Print(connect);
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		return err
	}

	DB = &DBWrapper{DB : db};

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Node{}, &Message{})
	return nil
}

func (db *DBWrapper) nodeAuth(connId string) bool{
	var node Node

	if err := db.DB.Where("ID = ?", connId).First(&node).Error; err != nil {
		log.Print(err)
		return false
	}

	return true
}

func (db *DBWrapper) createNode(name string) (Node, error){
	log.Print("create node " + name)
	node := Node{ID : uuid.NewV4().String(), Name : name}

	if err := db.DB.Create(&node).Error; err != nil {
		return Node{}, err
	}

	return node, nil
}

func (db *DBWrapper) getNodeMessages(connId string) ([]Message, error){
	var messages []Message

	node := Node{ID : connId}

	if err := db.DB.Model(&node).Related(&messages).Error; err != nil {
		return []Message{}, err
	}

	return messages, nil
}

func (db *DBWrapper) getAllNode() ([]Node, error){
	var nodes []Node

	if err := db.DB.Find(&nodes).Error; err != nil {
		return nil, err
	}

	return nodes, nil
}

	

	
