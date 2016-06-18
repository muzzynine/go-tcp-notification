var React = require('react');
var _ = require('lodash');
var Panel = require('react-bootstrap').Panel;
var ButtonToolbar = require('react-bootstrap').ButtonToolbar;
var Button = require('react-bootstrap').Button;
var AdminAppStore = require('../stores/AdminApplicationStore');
var AdminAppAction = require('../actions/AdminAppActions');
var Table = require('react-bootstrap').Table;
var Button = require('react-bootstrap').Button;
var AddMessageModal = require('./AddMessageModal.react');
var DeleteMessageModal = require('./deleteMessageModal.react');
var UpdateMessageModal = require('./UpdateMessageModal.react');

var UserDescription = React.createClass({
    getInitialState : function(){
	return {
	    loaded : false,
	    messageList : [],
	    showAddModal : false,
	    showDeleteModal : false,
	    showUpdateModal : false
	}
    },

    componentWillReceiveProps : function(){
	AdminAppAction.getConnection(this.props.user.ID);
    },
    
    componentWillMount : function(){
	AdminAppStore.addConnectionListener(this.setUser);
    },

    componentWillUnmount : function(){
	AdminAppStore.removeConnectionListener(this.setUser);
    },

    setUser : function(err, messageList){
	console.log("setUser")
	this.setState({
	    loaded : true,
	    messageList : messageList.Msgs
	});
    },

    openAddModal : function(){
	this.setState({
	    showAddModal : true
	});
    },

    closeAddModal : function(){
	this.setState({
	    showAddModal : false
	});
    },

    openUpdateModal : function(){
	this.setState({
	    showUpdateModal : true
	});
    },

    closeUpdateModal : function(){
	this.setState({
	    showUpdateModal : false
	});
    },

    openDeleteModal : function(){
	this.setState({
	    showDeleteModal : true
	});
    },

    closeDeleteModal : function(){
	this.setState({
	    showDeleteModal : false
	});
    },
	
    render : function(){
	var show;
	var messageRows = []
	if(!this.props.user.ID){
	    show = (<h4> Select Connection </h4>);
	} else {
	    if(this.state.loaded){
		this.state.messageList.forEach(function(message){
		    console.log(message)
		    messageRows.push(<tr><td>{message.MsgId}</td><td>{message.Msg}</td></tr>);
		});
		show = (
			<div>
			<h4>{this.props.user.Name}</h4>
			<Table>
			<thead>
			<tr>
			<th>#</th>
			<th>Message</th>
			</tr>
			</thead>
			<tbody>
			{messageRows}
		    </tbody>
			</Table>
			<Button onClick={this.openAddModal}>Add</Button>
			<Button onClick={this.openUpdateModal}>Update</Button>
			<Button onClick={this.openDeleteModal}>Delete</Button>
			</div>
		);
	    } else {
		show = (<h4> Loading... </h4>);
	    }
	}
	
	
	return (
		<div>
		<AddMessageModal showModal={this.state.showAddModal} user={this.props.user} closeModal={this.closeAddModal} />
		<DeleteMessageModal showModal={this.state.showDeleteModal} user={this.props.user} closeModal={this.closeDeleteModal} />
		<UpdateMessageModal showModal={this.state.showUpdateModal} user={this.props.user} closeModal={this.closeUpdateModal} />
										       
		<Panel header="Connection Description">
		{show}
		</Panel>
		</div>
	);
    }
});


module.exports = UserDescription;






