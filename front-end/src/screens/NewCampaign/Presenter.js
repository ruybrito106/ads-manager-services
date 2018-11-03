import React from "react";

import View from "./View";
import Model from "./Model";
import places from "./data/places";
import ads from "./data/ads";
import APIGateway from "../../server/APIGateway";

export default class Presenter extends React.Component {
  state = Model;

  render() {
    return (
      <View
        {...this.state}
        onVisitsGoalChange={this.handleVisitsGoalChange}
        onNameChange={this.handleNameChange}
        onPlacesChange={this.handlePlacesChange}
        onAdsChange={this.handleAdsChange}
        onStartDateChange={this.handleStartDateChange}
        onEndDateChange={this.handleEndDateChange}
        onSubmit={this.handleSubmit}
        placeOptions={places}
        adOptions={ads}
      />
    );
  }

  handleVisitsGoalChange = (e, { value }) => {
    this.setState({ visitsGoal: value });
  };

  handleNameChange = (e, { value }) => {
    this.setState({ name: value });
  };

  handlePlacesChange = (e, { value }) => {
    this.setState({ places: value });
  };

  handleAdsChange = (e, { value }) => {
    this.setState({ ads: value });
  };

  handleStartDateChange = value => {
    this.setState({ startDate: value });
  };

  handleEndDateChange = value => {
    this.setState({ endDate: value });
  };

  handleSubmit = () => {
    const { name, visitsGoal, startDate, endDate } = this.state;

    APIGateway.createCampaign({
      data: JSON.stringify({
        end_ts: endDate.unix(),
        start_ts: startDate.unix(),
        visits_goal: parseInt(visitsGoal),
        status: "active",
        name
      })
    });
  };
}
