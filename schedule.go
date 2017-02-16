package victorops

// Schedule current schedule
type Schedule struct {
	Oncall         string `json:"oncall,omitempty"`
	OverrideOncall string `json:"overrideoncall,omitempty"`
	PolicyType     string `json:"policyType,omitempty"`
	RotationName   string `json:"rotationName,omitempty"`
	ShiftName      string `json:"shiftName,omitempty"`
	ShiftRoll      int    `json:"shiftRoll,omitempty"`
	Rolls          []Roll `json:"rolls,omitempty"`
}

// Roll is shcedule roll
type Roll struct {
	Change int    `json:"change,omitempty"`
	Until  int    `json:"until,omitempty"`
	Oncall string `json:"oncall,omitempty"`
	IsRoll bool   `json:"isroll,omitempty"`
}

// Override describes oncall overrides
type Override struct {
	Orig  string `json:"orig,omitempty"`
	Over  string `json:"over,omitempty"`
	Start int    `json:"start,omitempty"`
	End   int    `json:"end,omitempty"`
}

// TeamSchedule is api response struct
type TeamSchedule struct {
	Team      string     `json:"team"`
	Schedules []Schedule `json:"schedule"`
	Overrides []Override `json:"overrides"`
}

// ResponseResult is result of an api request
type ResponseResult struct {
	Result string `json:"result"`
}

//UserOncallSchedule returns the on-call schedule for a user for all teams, including on-call overrides
func (client *Client) UserOncallSchedule(nick string) ([]TeamSchedule, error) {
	var schd []TeamSchedule
	err := client.sendRequest("GET", "api-public/v1/user/"+nick+"/oncall/schedule", nil, &schd)
	if err != nil {
		return schd, err
	}
	return schd, nil
}

// TeamOncallSchedule Get the on-call schedule for a user for all teams, including on-call overrides
func (client *Client) TeamOncallSchedule(team string) (TeamSchedule, error) {
	var schd TeamSchedule
	err := client.sendRequest("GET", "api-public/v1/team/"+team+"/oncall/schedule", nil, &schd)
	if err != nil {
		return schd, err
	}
	return schd, nil
}

// CreateOncallOverride replaces a currently on-call user on the team with another
func (client *Client) CreateOncallOverride(fromUser, toUser, team string) (ResponseResult, error) {
	return ResponseResult{}, nil
}
