var Dispatcher = require('../services/AppDispatcher');
var AdminAppConstants = require('../constants/AdminAppConstants');
var request = require('superagent');
var Promise = require('bluebird');

var HOST_ADDR = "http://localhost:8080"

var response = {
    accessToken : "access"
};

var AdminAppAction = {
    getConnections : function(){
	var baseUrl = HOST_ADDR + "/connections";
	request.get(baseUrl).end(function(err, res){
	    if(err){
		Dispatcher.dispatch({
		    type : AdminAppConstants.FIND_USERS,
		    success : false,
		    errorMessage : err.toString()
		});
	    } else {
		Dispatcher.dispatch({
		    type : AdminAppConstants.FIND_USERS,
		    success : true,
		    connectionList : res.body
		});
	    }
	});
    },

    getConnection : function(connId){
	request.get(HOST_ADDR + "/connection?connId=" + connId).end(function(err, res){	    
	    if(err){
		console.log(err)
		Dispatcher.dispatch({
		    type : AdminAppConstants.FIND_USER,
		    success : false,
		    errorMessage : err.toString()
		});

	    } else {
		console.log(res.body)
		Dispatcher.dispatch({
		    type : AdminAppConstants.FIND_USER,
		    success : true,
		    messageList : res.body
		});
	    }
	});
    },

    addMessage : function(connId, message){
	request.post(HOST_ADDR + "/message")
	    .send({
		connId : connId,
		msg : message
	    })
	    .end(function(err, res){
		if(err){
		    Dispatcher.dispatch({
			type : AdminAppConstants.ADD_MSG,
			success : false,
			errorMessage : err.toString()
		    });
		} else {
		    Dispatcher.dispatch({
			type : AdminAppConstants.ADD_MSG,
			success : true
		    });
		}
	    })
    },

    deleteMessage : function(connId, msgId){
	request.delete(HOST_ADDR + "/message")
	    .send({
		connId : connId,
		msgId : msgId
	    })
	    .end(function(err, res){
		if(err){
		    Dispatcher.dispatch({
			type : AdminAppConstants.DEL_MSG,
			success : false,
			errorMessage : err.toString()
		    });
		} else {
		    Dispatcher.dispatch({
			type : AdminAppConstants.DEL_MSG,
			success : true
		    });
		}
	    })
    },

    updateMessage : function(connId, msgId, message){
	request.put(HOST_ADDR + "/message")
	    .send({
		connId : connId,
		msgId : msgId,
		msg : message
	    })
	    .end(function(err, res){
		if(err){
		    Dispatcher.dispatch({
			type : AdminAppConstants.UPT_MSG,
			success : false,
			errorMessage : err.toString()
		    });
		} else {
		    Dispatcher.dispatch({
			type : AdminAppConstants.UPT_MSG,
			success : true
		    });
		}
	    })
    },

	    

    forceSignOutUser : function(uid){
	request.delete(HOST_ADDR + "/user/" + uid).end(function(err, res){
	    if(err){
		Dispatcher.dispatch({
		    type : AdminAppConstants.FORCE_SIGN_OUT_USER,
		    success : false,
		    errorMessage : "force Sign out failed"
		});
	    } else {
		Dispatcher.dispatch({
		    type : AdminAppConstants.FORCE_SIGN_OUT_USER,
		    success: true,
		    forcedUser : uid
		});
	    }
	});
    },

    register : function(name){
	request.post(HOST_ADDR + "/connection")
	    .send({
		name : name
	    })
	    .end(function(err, res){
		if(err){
		    Dispatcher.dispatch({
			type : AdminAppConstants.REGISTER,
			success : false,
			errorMessage : err.toString()
		    });
		} else {
		    Dispatcher.dispatch({
			type : AdminAppConstants.REGISTER,
			success : true,
			nodeInfo : res.body
		    });
		}
	    })
    },
		

    temporaryBan : function(uid, reason, start, duration){
	request.post(HOST_ADDR + "/ban/" + uid)
	    .send({
		reason : reason,
		startDate : start,
		duration : duration
	    })
	    .end(function(err, res){
		if(err){
		    Dispatcher.dispatch({
			type : AdminAppConstants.TEMPORARY_BAN_USER,
			success : false,
			errorMessage : "temporary ban failed"
		    });
		} else {
		    Dispatcher.dispatch({
			type : AdminAppConstants.TEMPORARY_BAN_USER,
			success : true,
			banInfo : {
			    uid : uid,
			    reason : reason,
			    start : start,
			    duration : duration
			}
		    });
		}
	    });
    }
}

module.exports = AdminAppAction;

