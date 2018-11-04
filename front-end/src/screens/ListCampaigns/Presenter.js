import React from "react";
import qs from "query-string";

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
    const campaign = this.state.campaigns.find(c => c.id === id);

    this.props.history.push(`campaigns/edit?${qs.stringify(campaign)}`);
  };

  handleNewCampaign = () => {
    this.props.history.push("campaigns/new");
  };
}
