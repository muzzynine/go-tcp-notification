var React = require('react');

var UserInfoTableHeader = React.createClass({
    render : function(){
	
	return (
		<tr>
		<th>ConnectionId</th>
		<th>Name</th>
		<th>IPAddr</th>
		<th>Status</th>
		</tr>
	);
    }
});

module.exports = UserInfoTableHeader;
