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
  onNameChange,
  onPlacesChange,
  onAdsChange,
  onSubmit
}) => (
  <div className="edit-campaign">
    <Card className="edit-campaign__card">
      <Card.Content textAlign="left">
        <p>Name</p>
        <Input
          className="edit-campaign__input"
          icon="pencil"
          iconPosition="left"
          fluid
          value={name}
          onChange={onNameChange}
          placeholder="Black friday Rappi!"
        />
        <p>Visits Goal</p>
        <Input
          className="edit-campaign__input"
          icon="target"
          iconPosition="left"
          fluid
          disabled
          value={visitsGoal}
          placeholder="3.000"
        />
        <p>Places</p>
        <Dropdown
          className="edit-campaign__input"
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
          className="edit-campaign__input"
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
        <div className="edit-campaign__input">
          <DatePicker
            className="edit-campaign__date--disabled"
            selected={startDate}
            selectsStart
            disabled
            startDate={startDate}
            endDate={endDate}
          />
          <DatePicker
            className="edit-campaign__date--disabled"
            selected={endDate}
            selectsEnd
            disabled
            startDate={startDate}
            endDate={endDate}
          />
        </div>
        <Button
          className="edit-campaign__input"
          primary
          fluid
          onClick={onSubmit}
        >
          Edit Campaign
        </Button>
      </Card.Content>
    </Card>
  </div>
);
