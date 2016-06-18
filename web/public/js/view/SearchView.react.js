var React = require('react');
var Grid = require('react-bootstrap').Grid;
var Row = require('react-bootstrap').Row;
var Col = require('react-bootstrap').Col;
var PageHeader = require('react-bootstrap').PageHeader;
var FilterableUserTable = require('../components/filterableUserTable.react');
var UserDescription = require('../components/UserDescription.react');
var ButtonToolbar = require('react-bootstrap').ButtonToolbar;
var Button = require('react-bootstrap').Button;
//var AdminAppAction = require('../action/AdminAppActions');
var RegisterModal = require('../components/registerModal.react');

var SearchView = React.createClass({

    getInitialState : function(){
	return {
	    selectedUser : {},
	    showModal : false
	};
    },

    openModal : function(){
	this.setState({
	    showModal : true
	})
    },

    closeModal : function(){
	this.setState({
	    showModal : false
	})
    },

    render : function(){
	return (
		<Grid fluid={false}>
		<RegisterModal showModal={this.state.showModal} closeModal={this.closeModal} />
		<Row>
		<Col md={9}>
		<PageHeader>Signboard Management <small>Output Control</small></PageHeader>
		</Col>
		<Col md={1}>
		<ButtonToolbar>
		<Button
	    onClick={this.openModal}
		>Register</Button>
		</ButtonToolbar>
		</Col>
		</Row>
		<Row>
		<Col md={6}>
		<FilterableUserTable onSelectUser={this._onSelectUser} />
		</Col>
		<Col md={4}>
		<UserDescription user={this.state.selectedUser} />
		</Col>
		</Row>
		</Grid>
		

	);
    },

    _onSelectUser : function(user){
	this.setState({
	    selectedUser : user
	});
    }
});


module.exports = SearchView;
