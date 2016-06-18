package rpc

import (
	"errors"
	"net/rpc"
)

type NotificationRPC struct {
	Client *rpc.Client
}

var (
	ErrAddMessage = errors.New("AddMessage client push failed")
	ErrDeleteMessage = errors.New("DeleteMessage client push failed")
	ErrUpdateMessage = errors.New("UpdateMessage client push failed")
)

const (
	notificationRPC = "ConnectionRPC"
	NotificationChangeBoardSetting = "ConnectionRPC.ChangeBoardSetting"
	GetConnections = "ConnectionRPC.GetConnections"
	AddMessage = "ConnectionRPC.AddMessage"
	DeleteMessage = "ConnectionRPC.DeleteMessage"
	UpdateMessage = "ConnectionRPC.UpdateMessage"
)

type AddMessageArgs struct {
	ConnId string
	MsgId uint
	Msg string
}

type DeleteMessageArgs struct {
	ConnId string
	MsgId uint
}

type UpdateMessageArgs struct {
	ConnId string
	MsgId uint
	Msg string
}

type GetConnectionsArgs struct {
}

type GetConnectionsResp struct {
	Description []GetConnectionsDescription
}

type GetConnectionsDescription struct {
	ConnId string
	ConnIPAddr string
}

func (nrpc *NotificationRPC) AddMessage(args *AddMessageArgs) error {
	var reply int

	err := nrpc.Client.Call(AddMessage, args, &reply)
	if err != nil || reply == 0 {
		return ErrAddMessage
	}

	return nil
}

func (nrpc *NotificationRPC) DeleteMessage(args *DeleteMessageArgs) error {
	var reply int

	err := nrpc.Client.Call(DeleteMessage, args, &reply)
	if err != nil || reply == 0 {
		return ErrDeleteMessage
	}

	return nil
}

func (nrpc *NotificationRPC) UpdateMessage(args *UpdateMessageArgs) error {
	var reply int

	err := nrpc.Client.Call(UpdateMessage, args, &reply)
	if err != nil || reply == 0 {
		return ErrUpdateMessage
	}

	return nil
}	


func (nrpc *NotificationRPC) GetConnections(args *GetConnectionsArgs) ([]GetConnectionsDescription, error) {
	var reply GetConnectionsResp

	err := nrpc.Client.Call(GetConnections, args, &reply)
	if err != nil {
		return nil, err
	}
	
	return reply.Description, nil
}
	
	
