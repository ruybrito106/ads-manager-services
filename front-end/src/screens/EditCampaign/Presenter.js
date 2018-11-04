import React from "react";
import qs from "query-string";
import moment from "moment";

import View from "./View";
import Model from "./Model";
import places from "../../data/places";
import ads from "../../data/ads";
import APIGateway from "../../server/APIGateway";

export default class Presenter extends React.Component {
  state = Model;

  componentDidMount = () => {
    const {
      id,
      name,
      start_ts,
      end_ts,
      status,
      visits_goal,
      places,
      ads
    } = qs.parse(this.props.location.search);

    this.setState({
      id,
      name,
      startDate: moment.unix(start_ts),
      endDate: moment.unix(end_ts),
      status,
      visitsGoal: visits_goal,
      places: Array.isArray(places) ? places : [places],
      ads: Array.isArray(ads) ? ads : [ads]
    });
  };

  render() {
    return (
      <View
        {...this.state}
        onNameChange={this.handleNameChange}
        onPlacesChange={this.handlePlacesChange}
        onAdsChange={this.handleAdsChange}
        onSubmit={this.handleSubmit}
        onCancel={this.handleCancel}
        placeOptions={places}
        adOptions={ads}
      />
    );
  }

  handleNameChange = (e, { value }) => {
    this.setState({ name: value });
  };

  handlePlacesChange = (e, { value }) => {
    this.setState({ places: value });
  };

  handleAdsChange = (e, { value }) => {
    this.setState({ ads: value });
  };

  handleSubmit = () => {
    const {
      id,
      name,
      visitsGoal,
      startDate,
      endDate,
      places,
      ads
    } = this.state;

    APIGateway.editCampaign({
      data: JSON.stringify({
        end_ts: endDate.unix(),
        start_ts: startDate.unix(),
        visits_goal: parseInt(visitsGoal),
        status: "active",
        name,
        places,
        ads,
        id: parseInt(id)
      }),
      onSuccess: () => this.props.history.push("/campaigns"),
      onFailure: () => this.props.history.push("/campaigns")
    });
  };

  handleCancel = () => {
    this.props.history.push("/campaigns");
  };
}
