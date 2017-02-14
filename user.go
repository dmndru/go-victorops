package victorops

import "encoding/json"

// User contains all the information of a user
type User struct {
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	UserName            string `json:"username"`
	Email               string `json:"email"`
	CreatedAt           string `json:"createdAt"`
	PasswordLastUpdated string `json:"passwordLastUpdated"`
	Verified            bool   `json:"verified"`
	SelfURL             string `json:"_selfUrl"`
}

// UsersResponse response to GetUsers
type usersResponse struct {
	Members [][]User `json:"users"`
	SelfURL string   `json:"_selfUrl"`
}

// UserInfo params for creating/modifying a user
type UserInfo struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Admin           bool   `json:"admin,omitempty"`
	ExpirationHours int    `json:"expirationHours,omitempty"`
}

// NewUser initialize struct for creating a user
func NewUser(name, surname, nick, email string, isAdmin bool, expire int) UserInfo {
	return UserInfo{name, surname, nick, email, isAdmin, expire}
}

// GetUsers Get a list of users for your organization
func (client *Client) GetUsers() ([][]User, error) {
	var (
		users [][]User
	)
	response := &usersResponse{}
	err := client.sendRequest("GET", "api-public/v1/user", nil, response)
	if err != nil {
		return users, err
	}
	return response.Members, nil
}

// AddUser Add a user to your organization
func (client *Client) AddUser(info UserInfo) (User, error) {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return User{}, err
	}
	response := &User{}
	err = client.sendRequest("POST", "api-public/v1/user", jsonData, response)
	if err != nil {
		return User{}, err
	}
	return *response, nil
}

// RemoveUser Remove a user from your organization
func (client *Client) RemoveUser(username, replacement string) error {
	values := map[string]string{"replacement": replacement}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}
	err = client.sendRequest("DELETE", "api-public/v1/user/"+username, jsonData, nil)
	return err
}

// GetUserInfo Get the information for the specified user
func (client *Client) GetUserInfo(nick string) (User, error) {
	response := &User{}
	err := client.sendRequest("GET", "api-public/v1/user/"+nick, nil, response)
	if err != nil {
		return User{}, err
	}
	return *response, nil
}

// UpdateUser Update the designated user
func (client *Client) UpdateUser(nick string, info UserInfo) (User, error) {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return User{}, err
	}
	response := &User{}
	err = client.sendRequest("PUT", "api-public/v1/user/"+nick, jsonData, response)
	if err != nil {
		return User{}, err
	}
	return *response, nil
}
