package victorops

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
