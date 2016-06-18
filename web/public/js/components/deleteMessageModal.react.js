var React = require('react');
var Modal = require('react-bootstrap').Modal;
var FormGroup = require('react-bootstrap').FormGroup;
var ControlLabel = require('react-bootstrap').ControlLabel;
var FormControl = require('react-bootstrap').FormControl;
var Button = require('react-bootstrap').Button;
var AdminAppStore = require('../stores/AdminApplicationStore');
var AdminAppAction = require('../actions/AdminAppActions');

var DeleteMessageModal = React.createClass({
    getInitialState : function(){
	return {
	    status : "wait",
	    error : "",
	    inputId : ""
	}
    },

    closeModal : function(){
	this.setState({
	    status : "wait",
	    error : "",
	    inputId : ""
	});
	this.props.closeModal();
    },
    
    register : function(name){
	AdminAppStore.addDeleteMessageListener(this.registered)
	AdminAppAction.deleteMessage(this.props.user.ID, this.state.inputId)
    },

    registered : function(err){
 	if(err){
	    //error handling
	    this.setState({
		status : "error",
		error : err.toString()
	    });
	} else {
	    this.setState({
		status : "success",
	    });
	}
	AdminAppStore.removeDeleteMessageListener(this.registered)
    },

    handleChange : function(e){
	this.setState({
	    inputId : e.target.value
	});
    },

    render : function(){
	var notYet = (
	    	<Modal show={this.props.showModal} onHide={this.closeModal}>
		<Modal.Header closeButton>
		<Modal.Title>Delete Message</Modal.Title>
		</Modal.Header>
		<Modal.Body>
		
		<form>
		<FormGroup controlId="DeleteMessage">
		<ControlLabel>Input message id</ControlLabel>
		<FormControl type="text" value={this.state.inputId} placeholder="msg id" onChange={this.handleChange} />
		</FormGroup>
		</form>

		</Modal.Body>
		<Modal.Footer>
		<Button bsStyle="primary" onClick={this.register}>Delete</Button>
		<Button onClick={this.closeModal}>Close</Button>
		</Modal.Footer>
		</Modal>
	);

	var yet = (
	    	<Modal show={this.props.showModal} onHide={this.closeModal}>
	    	<Modal.Header closeButton>
		<Modal.Title>Success</Modal.Title>
		</Modal.Header>
		<Modal.Body>
		<p> Success </p>
		</Modal.Body>
		<Modal.Footer>
		<Button onClick={this.closeModal}>Close</Button>
		</Modal.Footer>
		</Modal>
	);

	var error = (
	    	<Modal show={this.props.showModal} onHide={this.closeModal}>
	    	<Modal.Header closeButton>
		<Modal.Title>Failed</Modal.Title>
		</Modal.Header>
		<Modal.Body>
		<p> {this.state.error} </p>
		</Modal.Body>
		<Modal.Footer>
		<Button onClick={this.closeModal}>Close</Button>
		</Modal.Footer>
		</Modal>
	);
	    

	var show = this.state.status === "wait" ? notYet
	    : this.state.status === "success" ? yet : error;
	    
	return (
	    <div>
		{show}
	    </div>
	);
    }
});


module.exports = DeleteMessageModal;
	
