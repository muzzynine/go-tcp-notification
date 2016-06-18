var React = require('react');

var UserRow = React.createClass({
    render : function(){
	return (
		<tr onClick={this.onSelectRow}>
		<td>{this.props.user.id}</td>
		<td>{this.props.user.auth_id}</td>
		<td>{this.props.user.nickname}</td>
	    <td>{this.props.user.sex}</td>
		</tr>
	);
    },

    onSelectRow : function(){
	this.props.onSelectUser(this.props.user);
    }
});

module.exports = UserRow;
