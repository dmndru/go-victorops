package victorops

import "net/url"

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
	firstName       string
	lastName        string
	username        string
	email           string
	admin           bool
	expirationHours int
}

// GetUsers Get a list of users for your organization
func (client *Client) GetUsers() ([][]User, error) {
	var (
		users [][]User
	)
	response := &usersResponse{}
	err := client.get("api-public/v1/user", url.Values{}, response, true)
	if err != nil {
		return users, err
	}
	return response.Members, nil
}

// AddUser Add a user to your organization
func (client *Client) AddUser(info UserInfo) (User, error) {
	return User{}, nil
}

// RemoveUser Remove a user from your organization
func (client *Client) RemoveUser(user, replacement string) error {
	return nil
}

// GetUserInfo Get the information for the specified user
func (client *Client) GetUserInfo(user string) (User, error) {
	return User{}, nil
}

// UpdateUser Update the designated user
func (client *Client) UpdateUser(info UserInfo) (User, error) {
	return User{}, nil
}
