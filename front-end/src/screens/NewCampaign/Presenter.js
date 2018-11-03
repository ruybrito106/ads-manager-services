import React from "react";

import View from "./View";
import Model from "./Model";
import places from "./data/places";
import ads from "./data/ads";

export default class Presenter extends React.Component {
  state = Model;

  render() {
    return (
      <View
        {...this.state}
        onVisitsGoalChange={this.handleVisitsGoalChange}
        onCpvChange={this.handleCpvChange}
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

  handleCpvChange = (e, { value }) => {
    this.setState({ cpv: value });
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

  handleSubmit = () => {};
}
