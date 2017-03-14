package victorops

import (
	"encoding/json"
	"time"
)

// Schedule current schedule
type Schedule struct {
	Oncall         string    `json:"oncall,omitempty"`
	OverrideOncall string    `json:"overrideoncall,omitempty"`
	PolicyType     string    `json:"policyType,omitempty"`
	RotationName   string    `json:"rotationName,omitempty"`
	ShiftName      string    `json:"shiftName,omitempty"`
	ShiftRoll      time.Time `json:"shiftRoll,omitempty"`
	Rolls          []Roll    `json:"rolls,omitempty"`
}

// Roll is shcedule roll
type Roll struct {
	Start  time.Time `json:"change,omitempty"`
	End    time.Time `json:"until,omitempty"`
	Oncall string    `json:"oncall,omitempty"`
	IsRoll bool      `json:"isroll,omitempty"`
}

// Override describes oncall overrides
type Override struct {
	Orig  string    `json:"orig,omitempty"`
	Over  string    `json:"over,omitempty"`
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
}

// TeamSchedule is api response struct
type TeamSchedule struct {
	Team      string     `json:"team"`
	Schedules []Schedule `json:"schedule"`
	Overrides []Override `json:"overrides"`
}

// TeamOverride is a pyaload for creating override request
type teamOverride struct {
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
}

// ResponseResult is result of an api request
type ResponseResult struct {
	Result string `json:"result"`
}

//UserOncallSchedule returns the on-call schedule for a user for all teams, including on-call overrides
func (client *Client) UserOncallSchedule(nick string) ([]TeamSchedule, error) {
	var schd []TeamSchedule
	err := client.sendRequest("GET", "api-public/v1/user/"+nick+"/oncall/schedule", nil, &schd)
	return schd, err
}

// TeamOncallSchedule Get the on-call schedule for a user for all teams, including on-call overrides
func (client *Client) TeamOncallSchedule(team string) (TeamSchedule, error) {
	var schd TeamSchedule
	err := client.sendRequest("GET", "api-public/v1/team/"+team+"/oncall/schedule", nil, &schd)
	return schd, err
}

// CreateOncallOverride replaces a currently on-call user on the team with another
func (client *Client) CreateOncallOverride(fromUser, toUser, team string) (ResponseResult, error) {
	var resp ResponseResult
	values := teamOverride{fromUser, toUser}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}
	err = client.sendRequest("PATCH", "api-public/v1/team/"+team+"/oncall/user", jsonData, &resp)
	return resp, err
}
