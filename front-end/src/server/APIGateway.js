import axios from "axios";

export default class APIGateway {
  static createCampaign({ data, onSuccess = () => {}, onFailure = () => {} }) {
    axios
      .post("http://localhost:8080/campaigns/create", data)
      .then(response => {
        onSuccess(response);
      })
      .catch(error => {
        onFailure(error);
      });
  }

  static editCampaign({ data, onSuccess = () => {}, onFailure = () => {} }) {
    axios
      .post("http://localhost:8080/campaigns/edit", data)
      .then(response => {
        onSuccess(response);
      })
      .catch(error => {
        onFailure(error);
      });
  }

  static fetchCampaigns({ onSuccess = () => {}, onFailure = () => {} }) {
    axios
      .get("http://localhost:8080/campaigns")
      .then(response => {
        onSuccess(response);
      })
      .catch(error => {
        onFailure(error);
      });
  }
}
