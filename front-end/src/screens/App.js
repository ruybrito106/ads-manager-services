import React, { Component } from "react";
import {
  BrowserRouter as Router,
  Route,
  Redirect,
  Switch
} from "react-router-dom";

import Login from "./Login/Presenter";
import Register from "./Register/Presenter";
import NewCampaign from "./NewCampaign/Presenter";
import EditCampaign from "./EditCampaign/Presenter";
import ListCampaigns from "./ListCampaigns/Presenter";

import "./App.css";

class App extends Component {
  render() {
    return (
      <Router>
        <Switch>
          <Route exact path="/login" component={Login} />
          <Route exact path="/register" component={Register} />
          <Route exact path="/campaigns" component={ListCampaigns} />
          <Route exact path="/campaigns/new" component={NewCampaign} />
          <Route exact path="/campaigns/edit" component={EditCampaign} />
          <Redirect to="/register" />
        </Switch>
      </Router>
    );
  }
}

export default App;
