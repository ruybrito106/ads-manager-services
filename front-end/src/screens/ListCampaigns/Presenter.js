import React from "react";

import View from "./View";
import Model from "./Model";
import APIGateway from "../../server/APIGateway";

export default class Presenter extends React.Component {
  state = Model;

  componentDidMount = () => {
    APIGateway.fetchCampaigns({
      onSuccess: ({ data }) => this.setState({ campaigns: data })
    });
  };

  render() {
    return (
      <View
        {...this.state}
        onCampaignClick={this.handleCampaignClick}
        onNewCampaign={this.handleNewCampaign}
      />
    );
  }

  handleCampaignClick = id => {
    const {
      name,
      start_ts,
      end_ts,
      status,
      visits_goal
    } = this.state.campaigns.find(c => c.id === id);

    this.props.history.push(
      `campaigns/edit?id=${id}&name=${name}&start_ts=${start_ts}&end_ts=${end_ts}&status=${status}&visits_goal=${visits_goal}`
    );
  };

  handleNewCampaign = () => {
    this.props.history.push("campaigns/new");
  };
}
