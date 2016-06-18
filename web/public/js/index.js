var React = require('react');
var ReactDOM = require('react-dom');
var Router = require('react-router');
var Grid = require('react-bootstrap').Grid;
var Row = require('react-bootstrap').Row;
var Col = require('react-bootstrap').Col;
var SearchView = require('./view/SearchView.react');

var AdminAppStore = require('./stores/AdminApplicationStore');

var App = React.createClass({
    render : function(){
	return (
		<Grid fluid={false}>
		<Row>
		<Col md={12}>
		{this.props.children}
	    </Col>
		</Row>
		</Grid>
	);
    },
});

ReactDOM.render((
	<Router.Router>
	<Router.Route path="/" component={SearchView}>
	</Router.Route>
	</Router.Router>),
		document.getElementById('adminapp')
);















