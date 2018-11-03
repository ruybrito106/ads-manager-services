import React from "react";
import { Input, Card, Button, Dropdown } from "semantic-ui-react";
import DatePicker from "react-datepicker";

import "./View.css";

export default ({
  visitsGoal,
  startDate,
  endDate,
  name,
  places,
  ads,
  placeOptions,
  adOptions,
  onVisitsGoalChange,
  onStartDateChange,
  onEndDateChange,
  onNameChange,
  onPlacesChange,
  onAdsChange,
  onSubmit
}) => (
  <div className="new-campaign">
    <Card className="new-campaign__card">
      <Card.Content textAlign="left">
        <p>Name</p>
        <Input
          className="new-campaign__input"
          icon="pencil"
          iconPosition="left"
          fluid
          value={name}
          onChange={onNameChange}
          placeholder="Black friday Rappi!"
        />
        <p>Visits Goal</p>
        <Input
          className="new-campaign__input"
          icon="target"
          iconPosition="left"
          fluid
          value={visitsGoal}
          onChange={onVisitsGoalChange}
          placeholder="3.000"
        />
        <p>Places</p>
        <Dropdown
          className="new-campaign__input"
          fluid
          search
          options={placeOptions}
          selection
          multiple
          value={places}
          onChange={onPlacesChange}
          placeholder="Walmart, Shopping Recife, Gamerz..."
        />
        <p>Ads</p>
        <Dropdown
          className="new-campaign__input"
          fluid
          search
          options={adOptions}
          value={ads}
          onChange={onAdsChange}
          selection
          multiple
          placeholder="Manda FOODS!"
        />
        <p>Campaign Period</p>
        <div className="new-campaign__input">
          <DatePicker
            selected={startDate}
            selectsStart
            startDate={startDate}
            endDate={endDate}
            onChange={onStartDateChange}
          />
          <DatePicker
            selected={endDate}
            selectsEnd
            startDate={startDate}
            endDate={endDate}
            onChange={onEndDateChange}
          />
        </div>
        <Button
          className="new-campaign__input"
          primary
          fluid
          onClick={onSubmit}
        >
          Create Campaign
        </Button>
      </Card.Content>
    </Card>
  </div>
);
