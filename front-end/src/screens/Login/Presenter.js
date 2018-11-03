import React from "react";

import View from "./View";
import Model from "./Model";

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

  handleSubmit = () => {};
}
