import React from "react";

import View from "./View";
import Model from "./Model";
import APIGateway from "../../server/APIGateway";

export default class Presenter extends React.Component {
  state = Model;

  render() {
    return (
      <View
        {...this.state}
        onUsernameChange={this.handleUsernameChange}
        onPasswordChange={this.handlePasswordChange}
        onSubmit={this.handleSubmit}
      />
    );
  }

  handleUsernameChange = (e, { value }) => {
    this.setState({ username: value });
  };

  handlePasswordChange = (e, { value }) => {
    this.setState({ password: value });
  };

  handleSubmit = () => {
    const { username, password } = this.state;

    APIGateway.loginUser({
      data: JSON.stringify({
        id: username,
        password: password
      }),
      onSuccess: () => this.props.history.push("/campaigns"),
      onFailure: () => this.props.history.push("/login")
    });
  };
}
