var React = require('react');
var ReactDOM = require('react-dom');
var Router = require('react-router');
var Grid = require('react-bootstrap').Grid;
var Row = require('react-bootstrap').Row;
var Col = require('react-bootstrap').Col;
var LoginView = require('./view/loginView.react');
var AdminView = require('./view/adminView.react');
var DashboardView = require('./view/dashboardView.react');
var SearchView = require('./view/SearchView.react');

var AdminAppStore = require('./stores/AdminApplicationStore');

function getUserState(){
    return AdminAppStore.getUser();
}

var App = React.createClass({
    componentDidMount: function(){
	AdminAppStore.addLoginSuccessListener(this._onLoginSuccess);
    },

    componentWillUnmount: function(){
	AdminAppStore.removeLoginSuccessListener(this._onLoginSuccess);
    },
    
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
    _onLoginSuccess : function(){
	this.setState({
	    user : AdminAppStore.getUser()
	});
    }
});

ReactDOM.render((
	<Router.Router>
	<Router.Route path="/" component={App}>
	<Router.Route path="login" component={LoginView} />
	<Router.Route path="control" user={AdminAppStore.getUser()} component={AdminView} >
	<Router.IndexRoute component={DashboardView} />
	<Router.Route path="search" component={SearchView} />
	</Router.Route>
	</Router.Route>
	</Router.Router>),
		document.getElementById('adminapp')
);















