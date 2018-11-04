import React, { Component } from "react";
import {
  BrowserRouter as Router,
  Route,
  Redirect,
  Switch
} from "react-router-dom";

import Login from "./Login/Presenter";
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
          <Route exact path="/campaigns" component={ListCampaigns} />
          <Route exact path="/campaigns/new" component={NewCampaign} />
          <Route exact path="/campaigns/edit" component={EditCampaign} />
          <Redirect to="/login" />
        </Switch>
      </Router>
    );
  }
}

export default App;
