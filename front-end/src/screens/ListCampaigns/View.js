import React from "react";
import { Card, Button, List } from "semantic-ui-react";
import moment from "moment";

import "./View.css";

export default ({ campaigns, onCampaignClick, onNewCampaign }) => (
  <div className="list-campaigns">
    <Card className="list-campaigns__card">
      <Card.Content textAlign="left">
        <List
          className="list-campaigns__list"
          animated
          verticalAlign="middle"
          relaxed="very"
          selection
          size="medium"
        >
          {campaigns.map(({ id, name, start_ts, end_ts }) => (
            <List.Item key={id} onClick={() => onCampaignClick(id)}>
              <List.Icon
                name="money bill alternate outline"
                size="large"
                verticalAlign="middle"
              />
              <List.Content>
                <List.Header>{name}</List.Header>
                <List.Description>
                  {moment.unix(start_ts).format("DD MMM")} -{" "}
                  {moment.unix(end_ts).format("DD MMM")}
                </List.Description>
                <List.Description />
              </List.Content>
            </List.Item>
          ))}
        </List>
        <Button
          className="list-campaign__element"
          primary
          fluid
          onClick={onNewCampaign}
        >
          New Campaign
        </Button>
      </Card.Content>
    </Card>
  </div>
);
