import React from "react";
import { Input, Card, Button, Image, Divider } from "semantic-ui-react";

import "./View.css";
import Icon from "../../assets/login.svg";

export default ({
  username,
  password,
  password_confirm,
  error,
  onUsernameChange,
  onPasswordChange,
  onPasswordConfirmChange,
  onLogin,
  onSubmit
}) => (
  <div className="register">
    <Card className="register__card">
      <Card.Content textAlign="center">
        <Image className="register__image" size="small" src={Icon} />
        <Input
          className="register__input"
          icon="user"
          iconPosition="left"
          fluid
          placeholder="Username"
          value={username}
          onChange={onUsernameChange}
          error={error}
        />
        <Input
          className="register__input"
          icon="key"
          iconPosition="left"
          fluid
          placeholder="Password"
          value={password}
          onChange={onPasswordChange}
          error={error}
        />
        <Input
          className="register__input"
          icon="key"
          iconPosition="left"
          fluid
          placeholder="Confirm Password"
          value={password_confirm}
          onChange={onPasswordConfirmChange}
          error={error}
        />
        <Button className="register__input" primary fluid onClick={onSubmit}>
          Register
        </Button>
        <Divider horizontal>Or</Divider>
        <Button className="login__input" primary fluid onClick={onLogin}>
          Login
        </Button>
      </Card.Content>
    </Card>
  </div>
);
