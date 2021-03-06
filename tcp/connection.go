package main

import (
	"net"
	"sync"
	"errors"
)

type Connection struct {
	id string
	conn *net.TCPConn
}

type ConnectionStore struct {
	Connections map[string]*Connection
	mutex *sync.Mutex
}

func NewConnectionStore() *ConnectionStore {
	connectionMap := make(map[string]*Connection)
	store := &ConnectionStore{Connections : connectionMap, mutex: &sync.Mutex{}}

	return store
}



func (conn *Connection) GetId() string {
	return conn.id
}

func (conn *Connection) GetIPAddr() string {
	return conn.conn.RemoteAddr().String()
}

func (conn *Connection) GetConn() *net.TCPConn {
	return conn.conn
}


func (store *ConnectionStore) Lock(){
	store.mutex.Lock()
}

func (store *ConnectionStore) Unlock(){
	store.mutex.Unlock()
}


func (store *ConnectionStore) Add(key string, conn *net.TCPConn)(*Connection, error){
	store.Lock()
	if c, ok := store.Connections[key]; ok {
		store.Unlock()
		return c, nil
	} else {
		c = &Connection{id : key, conn : conn}
		store.Connections[key] = c
		store.Unlock()
		return c, nil
	}
}


func (store *ConnectionStore) Get(key string) (*Connection, bool) {
	store.Lock()
	if connection, ok := store.Connections[key]; ok {
		store.Unlock()
		return connection, true
	}
	store.Unlock()
	return nil, false
}

func (store *ConnectionStore) Delete(key string) (*Connection, error){
	store.Lock()
	if connection, ok := store.Connections[key]; ok {
		delete(store.Connections, key)
		connection.Close()
		store.Unlock()
		return connection, nil
	} else {
		store.Unlock()
		return nil, errors.New("not exist key")
	}
}

func (store *ConnectionStore) Close() {
	for key := range store.Connections {
		if connection, ok := store.Get(key); ok {
			delete(store.Connections, key)
			connection.Close()
		}
	}
}

func (store *ConnectionStore) GetAllConnection() map[string]*Connection {
	return store.Connections
}


func (connection *Connection) Close() error{
	return connection.conn.Close()
}

func (connection *Connection) PushMessage(payload []byte) error {
	if _, err := connection.conn.Write(payload); err != nil {
		return err
	}
	return nil
}



	

