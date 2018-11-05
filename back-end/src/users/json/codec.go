package json

import (
	"encoding/json"

	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

func UserToJSON(user *users.User) ([]byte, error) {
	return json.Marshal(&user)
}

func UserFromJSON(encoded []byte) (*users.User, error) {
	user := users.User{}
	if err := json.Unmarshal(encoded, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
