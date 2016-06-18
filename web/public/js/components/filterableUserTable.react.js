var React = require('react');
var Grid = require('react-bootstrap').Grid;
var Row = require('react-bootstrap').Row;
var Col = require('react-bootstrap').Col;

var UserInfoTable = require('./userInfoTable.react');

var FilterableUserTable = React.createClass({
    render : function(){
	return (
		<Grid fluid={false}>
		<Row>
		<Col md={6}>
		<UserInfoTable onSelectUser={this.props.onSelectUser}/>
		</Col>
		</Row>
		</Grid>
	);
    }
});

module.exports = FilterableUserTable;
