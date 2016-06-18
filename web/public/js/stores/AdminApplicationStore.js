var Dispatcher = require('../services/AppDispatcher');
var EventEmitter = require('events').EventEmitter;
var assign = require('object-assign');
var AdminAppConstants = require('../constants/AdminAppConstants');

var LOGIN_EVENT = "login";
var LOGIN_SUCCESS = "login_success";
var FIND_USERS = "find_users";
var FIND_USER = "find_user";
var REGISTER ="register";
var UPT_MSG = "update_message";
var ADD_MSG = "add_message";
var DEL_MSG = "delete_message";

var getConnectionsInfo = {
    connectionList : []
};

var AdminAppStore = assign({}, EventEmitter.prototype, {
    getConnectionList : function(){
	return getConnectionsInfo
    },

    emitConnectionListChange : function(err){
	this.emit(FIND_USERS, err);
    },

    addConnectionListListener : function(fn){
	this.on(FIND_USERS, fn);
    },

    removeConnectionListListner : function(fn){
	this.removeListener(FIND_USERS, fn);
    },

    addRegisterListener : function(fn){
	this.on(REGISTER, fn);
    },

    removeRegisterListener : function(fn){
	this.removeListener(REGISTER, fn)
    },

    emitRegistered : function(err, info){
	this.emit(REGISTER, err, info)
    },

    addConnectionListener : function(fn){
	this.on(FIND_USER, fn);
    },

    removeConnectionListener : function(fn){
	this.removeListener(FIND_USER, fn)
    },

    emitGetConnection : function(err, info){
	this.emit(FIND_USER, err, info)
    },

    addMessageListener : function(fn){
	this.on(ADD_MSG, fn);
    },

    removeMessageListener : function(fn){
	this.removeListener(ADD_MSG, fn);
    },

    emitAddMessage : function(err){
	this.emit(ADD_MSG, err);
    },

    addUpdateMessageListener : function(fn){
	this.on(UPT_MSG, fn);
    },

    removeUpdateMessageListener : function(fn){
	this.removeListener(UPT_MSG, fn);
    },

    emitUpdateMessage : function(err){
	this.emit(UPT_MSG, err);
    },

    addDeleteMessageListener : function(fn){
	this.on(DEL_MSG, fn);
    },

    removeDeleteMessageListener : function(fn){
	this.removeListener(DEL_MSG, fn);
    },

    emitDeleteMessage : function(err){
	this.emit(DEL_MSG, err);
    },

    
});

    
Dispatcher.register(function(action){
    switch(action.type){
 
    case AdminAppConstants.FIND_USER :
	if(action.success){
	    AdminAppStore.emitGetConnection(null, action.messageList);
	} else {
	    AdminAppStore.emitGetConnection(action.errorMessage);
	}
	break;

    case AdminAppConstants.FIND_USERS :
	if(action.success){
	    console.log(action)
	    getConnectionsInfo.connectionList = action.connectionList.Connections;
	    AdminAppStore.emitConnectionListChange();
	} else {
	    AdminAppStore.emitConnectionListChange(action.errorMessage);
	}
	break;

    case AdminAppConstants.REGISTER :
	if(action.success){
	    console.log("success")
	    AdminAppStore.emitRegistered(null, action.nodeInfo);
	} else {
	    console.log(action)
	    AdminAppStore.emitRegistered(action.errorMessage);
	}
	break;

    case AdminAppConstants.ADD_MSG :
	if(action.success){
	    AdminAppStore.emitAddMessage();
	} else {
	    AdminAppStore.emitAddMessage(action.errorMessage);
	}
	break;

    case AdminAppConstants.DEL_MSG :
	if(action.success){
	    AdminAppStore.emitDeleteMessage();
	} else {
	    AdminAppStore.emitDeleteMessage(action.errorMessage);
	}
	break;
	
    case AdminAppConstants.UPT_MSG :
	if(action.success){
	    AdminAppStore.emitUpdateMessage();
	} else {
	    AdminAppStore.emitUpdateMessage(action.errorMessage);
	}
	break;

	
    default :
	console.log("no op");
    }
});

module.exports = AdminAppStore;
		






