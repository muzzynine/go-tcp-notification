var React = require('react');
var Modal = require('react-bootstrap').Modal;
var FormGroup = require('react-bootstrap').FormGroup;
var ControlLabel = require('react-bootstrap').ControlLabel;
var FormControl = require('react-bootstrap').FormControl;
var Button = require('react-bootstrap').Button;
var AdminAppStore = require('../stores/AdminApplicationStore');
var AdminAppAction = require('../actions/AdminAppActions');

var AddMessageModal = React.createClass({
    getInitialState : function(){
	return {
	    status : "wait",
	    error : "",
	    inputValue : ""
	}
    },

    closeModal : function(){
	this.setState({
	    status : "wait",
	    error : "",
	    inputValue : ""
	});
	this.props.closeModal();
    },
    
    register : function(name){
	AdminAppStore.addMessageListener(this.registered)
	AdminAppAction.addMessage(this.props.user.ID, this.state.inputValue)
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
	    AdminAppAction.getConnection(this.props.user.Id);
	}
	AdminAppStore.removeAddMessageListener(this.registered)
    },

    handleChange : function(e){
	this.setState({
	    inputValue : e.target.value
	});
    },

    render : function(){
	var notYet = (
	    	<Modal show={this.props.showModal} onHide={this.closeModal}>
		<Modal.Header closeButton>
		<Modal.Title>Add Message</Modal.Title>
		</Modal.Header>
		<Modal.Body>
		
		<form>
		<FormGroup controlId="AddMessage">
		<ControlLabel>Input Message</ControlLabel>
		<FormControl type="text" value={this.state.inputValue} placeholder="New Message" onChange={this.handleChange} />
		</FormGroup>
		</form>

		</Modal.Body>
		<Modal.Footer>
		<Button bsStyle="primary" onClick={this.register}>Add</Button>
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


module.exports = AddMessageModal;
	
