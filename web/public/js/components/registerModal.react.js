var React = require('react');
var Modal = require('react-bootstrap').Modal;
var FormGroup = require('react-bootstrap').FormGroup;
var ControlLabel = require('react-bootstrap').ControlLabel;
var FormControl = require('react-bootstrap').FormControl;
var Button = require('react-bootstrap').Button;
var AdminAppStore = require('../stores/AdminApplicationStore');
var AdminAppAction = require('../actions/AdminAppActions');

var RegisterModal = React.createClass({
    getInitialState : function(){
	return {
	    status : "wait",
	    id : "",
	    name : "",
	    error : "",
	    inputValue : ""
	}
    },

    closeModal : function(){
	this.setState({
	    status : "wait",
	    id : "",
	    name : "",
	    error : "",
	    inputValue : ""
	});
	this.props.closeModal();
    },
    
    register : function(name){
	AdminAppStore.addRegisterListener(this.registered)
	AdminAppAction.register(this.state.inputValue)
    },

    registered : function(err, nodeInfo){
 	if(err){
	    //error handling
	    this.setState({
		status : "error",
		error : err.toString()
	    });
	} else {
	    console.log(nodeInfo)
	    this.setState({
		status : "success",
		id : nodeInfo.Id,
		name : nodeInfo.Name
	    });
	}
	AdminAppStore.removeRegisterListener(this.registered)
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
		<Modal.Title>Device Register</Modal.Title>
		</Modal.Header>
		<Modal.Body>
		
		<form>
		<FormGroup controlId="registerName">
		<ControlLabel>Input your device name</ControlLabel>
		<FormControl type="text" value={this.state.inputValue} placeholder="Enter device name" onChange={this.handleChange} />
		</FormGroup>
		</form>

		</Modal.Body>
		<Modal.Footer>
		<Button bsStyle="primary" onClick={this.register}>Register</Button>
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
		<p> ID : {this.state.id} NAME : {this.state.name} </p>
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


module.exports = RegisterModal;
	
