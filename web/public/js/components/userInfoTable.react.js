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
	    connectionList : [],
	    error : {
		errorState : false,
		errorMessage : ''
	    }
	}
    },

    componentWillMount : function(){
	AdminAppStore.addConnectionListListener(this.changeConnectionState);
	AdminAppAction.getConnections();
    },

    componentWillUnmount : function(){
	AdminAppStore.removeConnectionListListner(this.changeConnectionState);
    },

    changeConnectionState : function(err){
	if(err){
	    this.setState({
   		error : {
		    errorState : true,
		    errorMessage : err
		}
	    });
	} else {
	    this.setState({
		connectionList : AdminAppStore.getConnectionList().connectionList,
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

	this.state.connectionList.forEach(function(connection){
	    userRows.push(<UserRow user={connection} key={connection.ID} onSelectUser={self._onClickUser} />);
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
		

	
