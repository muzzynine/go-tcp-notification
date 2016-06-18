var React = require('react');
var UserInfoTableHeader = require('./userInfoTableHeader.react');
var UserRow = require('./userRow.react');
var Grid = require('react-bootstrap').Grid;
var Row = require('react-bootstrap').Row;
var Col = require('react-bootstrap').Col;
var Table = require('react-bootstrap').Table;
var AdminAppStore = require('../stores/AdminApplicationStore');
var AdminAppAction = require('../actions/AdminAppActions');

var UserInfoTable = React.createClass({
    getInitialState : function(){
	return {
	    customerList : [],
	    error : {
		errorState : false,
		errorMessage : ''
	    }
	}
    },

    componentWillMount : function(){
	AdminAppStore.addCustomerListListener(this.changeCustomerState);
    },

    componentWillUnmount : function(){
	AdminAppStore.removeCustomerListListner(this.changeCustomerState);
    },

    changeCustomerState : function(err){
	if(err){
	    this.setState({
   		error : {
		    errorState : true,
		    errorMessage : err
		}
	    });
	    console.log(err);
	} else {
	    this.setState({
		customerList : AdminAppStore.getCustomerList().customerList,
		error : {
		    errorState : false,
		    errorMessage : ''
		}
	    });
	}
    },
    
    render : function(){
	var userRows = [];
	var self = this;

	this.state.customerList.forEach(function(user){
	    userRows.push(<UserRow user={user} key={user.id} onSelectUser={self._onClickUser} />);
	});

	return (
		<Table striped bordered condensed hover responsive>
		<thead>
		<UserInfoTableHeader />
		</thead>
		<tbody>{userRows}</tbody>
		</Table>
	);
    },

    _onClickUser : function(user){
	this.props.onSelectUser(user);
    }
	
});

module.exports = UserInfoTable;
		

	
