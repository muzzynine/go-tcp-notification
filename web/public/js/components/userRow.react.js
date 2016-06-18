var React = require('react');
var FontAwesome = require('react-fontawesome');

var UserRow = React.createClass({
    render : function(){
	var statusActive = "Active"
	var statusInactive = "Inactive"

	var status = this.props.user.Status ? statusActive : statusInactive;

	console.log(status)
	
	return (
		<tr onClick={this.onSelectRow}>
		<td>{this.props.user.ID}</td>
		<td>{this.props.user.Name}</td>
		<td>{this.props.user.IPAddr}</td>
	    <td>{status}</td>
		</tr>
	);
    },

    onSelectRow : function(){
	this.props.onSelectUser(this.props.user);
    }
});

module.exports = UserRow;
