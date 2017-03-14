package victorops

import "encoding/json"

// Team struct describes a team
type Team struct {
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	MemberCount   int    `json:"memberCount"`
	Version       int    `json:"version"`
	IsDefaultTeam bool   `json:"isDefaultTeam"`
	SelfURL       string `json:"_selfUrl"`
	MembersURL    string `json:"_membersUrl"`
}

// TeamListResp is a list of teams for your organization
type TeamListResp struct {
	Teams   []Team `json:"teams"`
	SelfURL string `json:"_selfUrl"`
}

// TeamMember member of a team
type TeamMember struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Version   int    `json:"version"`
	Verified  string `json:"verified"`
	SelfURL   string `json:"_selfUrl"`
}

// TeamMembersResp contains response to GetTeamMembers method
type TeamMembersResp struct {
	Members []TeamMember `json:"teams"`
	SelfURL string       `json:"_selfUrl"`
}

// GetTeams Get a list of teams for your organization
func (client *Client) GetTeams() (TeamListResp, error) {
	var resp TeamListResp
	err := client.sendRequest("GET", "api-public/v1/team", nil, &resp)
	return resp, err
}

// AddTeam Add a team to your organization
func (client *Client) AddTeam(name string) (Team, error) {
	var team Team
	values := map[string]string{"name": name}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return team, err
	}
	err = client.sendRequest("POST", "api-public/v1/team", jsonData, &team)
	return team, err
}

// RemoveTeam Remove a team from your organization
func (client *Client) RemoveTeam(name string) error {
	err := client.sendRequest("DELETE", "api-public/v1/team/"+name, nil, nil)
	return err
}

// GetTeam Get the information for the specified team
func (client *Client) GetTeam(name string) (Team, error) {
	var team Team
	err := client.sendRequest("GET", "api-public/v1/team"+name, nil, &team)
	return team, err
}

// UpdateTeam Update the designated team
func (client *Client) UpdateTeam(name, newName string) (Team, error) {
	var team Team
	values := map[string]string{"name": newName}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return team, err
	}
	err = client.sendRequest("PUT", "api-public/v1/team"+name, jsonData, &team)
	return team, err
}
