import React from "react";
import { Input, Card, Button, Image } from "semantic-ui-react";

import "./View.css";
import Icon from "../../assets/login.svg";

export default ({
  username,
  password,
  error,
  onUsernameChange,
  onPasswordChange,
  onSubmit
}) => (
  <div className="login">
    <Card className="login__card">
      <Card.Content textAlign="center">
        <Image className="login__image" size="small" src={Icon} />
        <Input
          className="login__input"
          icon="user"
          iconPosition="left"
          fluid
          placeholder="Username"
          value={username}
          onChange={onUsernameChange}
          error={error}
        />
        <Input
          className="login__input"
          icon="key"
          iconPosition="left"
          fluid
          placeholder="Password"
          value={password}
          onChange={onPasswordChange}
          error={error}
        />
        <Button className="login__input" primary fluid onClick={onSubmit}>
          Login
        </Button>
      </Card.Content>
    </Card>
  </div>
);
