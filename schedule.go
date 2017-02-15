package victorops

// Schedule current schedule
type Schedule struct {
	Oncall         string     `json:"oncall"`
	OverrideOncall string     `json:"overrideoncall"`
	PolicyType     string     `json:"policyType"`
	RotationName   string     `json:"rotationName"`
	ShiftName      string     `json:"shiftName"`
	ShiftRoll      int        `json:"shiftRoll"`
	Rolls          []Roll     `json:"rolls"`
	Overrides      []Override `json:"overrides"`
}

// Roll is shcedule roll
type Roll struct {
	Change int    `json:"change"`
	Until  int    `json:"until"`
	Oncall string `json:"oncall"`
	IsRoll bool   `json:"isroll"`
}

// Override describes oncall overrides
type Override struct {
	Orig  string `json:"orig"`
	Over  string `json:"over"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

// TeamSchedule is api response struct
type TeamSchedule struct {
	Team      string     `json:"team"`
	Schedules []Schedule `json:"schedule"`
}

// ResponseResult is result of an api request
type ResponseResult struct {
	Result string `json:"result"`
}

//UserOncallSchedule returns the on-call schedule for a user for all teams, including on-call overrides
func (client *Client) UserOncallSchedule(nick string) ([]TeamSchedule, error) {
	var schd []TeamSchedule
	return schd, nil
}

// TeamOncallSchedule Get the on-call schedule for a user for all teams, including on-call overrides
func (client *Client) TeamOncallSchedule(team string) (TeamSchedule, error) {
	return TeamSchedule{}, nil
}

// CreateOncallOverride replaces a currently on-call user on the team with another
func (client *Client) CreateOncallOverride(fromUser, toUser, team string) (ResponseResult, error) {
	return ResponseResult{}, nil
}
