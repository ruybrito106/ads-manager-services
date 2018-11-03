import axios from "axios";

export default class APIGateway {
  static createCampaign({ data, onSuccess = () => {}, onFailure = () => {} }) {
      console.log(data)
    axios
      .post("http://localhost:8080/campaigns/create", data)
      .then(response => {
        onSuccess(response);
      })
      .catch(error => {
        onFailure(error);
      });
  }
}
