package main

import (
	"strconv"
	"net/http"
	"encoding/json"
	"log"
	nrpc "github.com/muzzynine/go-tcp-notification/rpc"

)

func main(){
	config := InitConfig()
	log.Print("Config initialized");
	err := BindDB(config.DBAddr, config.DBUser, config.DBPasswd, config.DBName)
	if err != nil {
		log.Print(err)
		log.Print("connect db failed")
		panic(err)
	}
	log.Print("DB connected successfully")
	
	fs := http.FileServer(http.Dir("public"))

	err = BindRPC(config.RpcAddr)
	if err != nil {
		log.Print("rpc bind failed")
		panic(err)
	}

	http.Handle("/", fs)
	http.HandleFunc("/connection", connectionHandle)
	http.HandleFunc("/connections", connectionsHandle)
	http.HandleFunc("/message", messageHandle)
	log.Print("http start")
	http.ListenAndServe(config.Port, nil)
}
/*
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request){
	}
}
*/

func connectionsHandle(w http.ResponseWriter, r *http.Request){
	if r.ParseForm() != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET" :
		connInfoArray, err := NRPC.GetConnections(&nrpc.GetConnectionsArgs{})
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		nodes, err := DB.getAllNode()
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Print("hello")
		
		connectionsResp := GetConnectionsResp{Connections : []GetConnectionsDescription{}}

		for _, node := range nodes {
			for _, info := range connInfoArray {
				if node.ID == info.ConnId {
					connectionsResp.Connections = append(connectionsResp.Connections, GetConnectionsDescription{ID : info.ConnId, Name : node.Name, IPAddr : info.ConnIPAddr, Status : true})
					break
				}
					
			}
			connectionsResp.Connections = append(connectionsResp.Connections, GetConnectionsDescription{ID : node.ID, Name : node.Name, IPAddr : "-", Status : false})
			
		}

		b, err := json.Marshal(connectionsResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
		default :
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}


func connectionHandle(w http.ResponseWriter, r *http.Request){
	log.Print("in connectionHandle")
	if r.ParseForm() != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET" :
		connId := r.Form.Get("connId")

		messages, err := DB.getNodeMessages(connId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		connectionResp := GetConnectionResp{Msgs : []GetConnectionMsg{}}

		for _, msg := range messages {
			connectionResp.Msgs = append(connectionResp.Msgs, GetConnectionMsg{MsgId : msg.ID, Msg : msg.Message})
		}

		b, err := json.Marshal(connectionResp)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
		
	case "POST" :
		decoder := json.NewDecoder(r.Body)
		var reqBody CreateNodeReq;
		err := decoder.Decode(&reqBody)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		node, err := DB.createNode(reqBody.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(CreateNodeResp{Id : node.ID, Name : node.Name})
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
		
		default :
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
		
	}
}

func messageHandle(w http.ResponseWriter, r *http.Request){
	if r.ParseForm() != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET" :
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("www"))
		return
	case "POST" :
		decoder := json.NewDecoder(r.Body)
		var reqBody AddMessageReq
		err := decoder.Decode(&reqBody)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Print(reqBody)

		message, err := DB.createMessage(reqBody.ConnId, reqBody.Msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := NRPC.AddMessage(&nrpc.AddMessageArgs{ConnId : reqBody.ConnId, MsgId : message.ID, Msg : message.Message}); err != nil {
			log.Print("rpc request failed")
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		return

	case "DELETE" :
		decoder := json.NewDecoder(r.Body)
		var reqBody DeleteMessageReq
		err := decoder.Decode(&reqBody)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		i, err := strconv.Atoi(reqBody.MsgId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		
		err = DB.deleteMessage(uint(i))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := NRPC.DeleteMessage(&nrpc.DeleteMessageArgs{ConnId : reqBody.ConnId, MsgId : uint(i)}); err != nil {
			log.Print("rpc request failed")
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		return

	case "PUT" :
		decoder := json.NewDecoder(r.Body)
		var reqBody UpdateMessageReq
		err := decoder.Decode(&reqBody)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Print(reqBody)

		i, err := strconv.Atoi(reqBody.MsgId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}


		message, err := DB.updateMessage(uint(i), reqBody.Msg)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := NRPC.UpdateMessage(&nrpc.UpdateMessageArgs{ConnId : reqBody.ConnId, MsgId : message.ID, Msg : message.Message}); err != nil {
			log.Print(err)
			log.Print("rpc request failed")
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
		
		default :
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
		
	}

}






