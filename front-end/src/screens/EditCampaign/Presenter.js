import React from "react";
import qs from "query-string";
import moment from "moment";

import View from "./View";
import Model from "./Model";
import places from "../../data/places";
import ads from "../../data/ads";

export default class Presenter extends React.Component {
  state = Model;

  componentDidMount = () => {
    const { id, name, start_ts, end_ts, status, visits_goal } = qs.parse(
      this.props.location.search
    );

    this.setState({
      id,
      name,
      startDate: moment.unix(start_ts),
      endDate: moment.unix(end_ts),
      status,
      visitsGoal: visits_goal
    });
  };

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
    this.props.history.push("/campaigns");
  };
}
