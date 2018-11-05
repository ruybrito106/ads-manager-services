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
        onPasswordConfirmChange={this.handlePasswordConfirmChange}
        onLogin={this.handleAlreadyRegistered}
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

  handlePasswordConfirmChange = (e, { value }) => {
    this.setState({ password_confirm: value });
  };

  handleAlreadyRegistered = (e, { value }) => {
    this.props.history.push("/login")
  };

  handleSubmit = () => {
    const { username, password, password_confirm } = this.state;

    if (password !== password_confirm) {
      this.props.history.push("/register")
    }

    APIGateway.registerUser({
      data: JSON.stringify({
        id: username,
        password: password
      }),
      onSuccess: () => this.props.history.push("/login"),
      onFailure: () => this.props.history.push("/register")
    });
  };
}
