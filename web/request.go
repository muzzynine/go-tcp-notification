package main


type CreateNodeResp struct {
	Id string `json:id`
	Name string `json:name`
}

type CreateNodeReq struct {
	Name string `json:name`
}

type GetConnectionsResp struct {
	Connections []GetConnectionsDescription
}

type GetConnectionResp struct {
	Msgs []GetConnectionMsg
}

type GetConnectionMsg struct {
	MsgId uint `json:msgId`
	Msg string `json:msg`
}

type GetConnectionsDescription struct {
	ID string `json:id`
	Name string `json:name`
	IPAddr string `json:ipAddr`
	Status bool `json:status`
}

type AddMessageReq struct {
	ConnId string `json:connId`
	Msg string `json:msg`
}

type DeleteMessageReq struct {
	ConnId string `json:connId`
	MsgId string `json:msgId`
}

type UpdateMessageReq struct {
	ConnId string `json:connId`
	MsgId string `json:msgId`
	Msg string `json:msg`
}

	
