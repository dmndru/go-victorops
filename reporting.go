package victorops

import "time"

// OncallLog a log of user shift changes for the team
type OncallLog struct {
	TeamSlug string    `json:"teamSlug"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	UserLogs []UserLog `json:"userLogs"`
}

// UserLog oncall log of a user
type UserLog struct {
	UserID        string       `json:"userId"`
	Total         HoursMinutes `json:"total"`
	AdjustedTotal HoursMinutes `json:"adjustedTotal"`
	Logs          []Log        `json:"log"`
}

// HoursMinutes type contains hours and minutes
type HoursMinutes struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

// Log duration of oncall
type Log struct {
	On       string       `json:"on"`
	Off      string       `json:"off"`
	Duration HoursMinutes `json:"duration"`
}

// IncidentsSearchParams are params of method GetIncidentsHistory needed for searching incidents
type IncidentsSearchParams struct {
	Offset         int       `json:"offset,omitempty"`         // The offset within the set of matching incidents
	Limit          int       `json:"limit,omitempty"`          // The maximum number of matching incidents to return (100 max)
	EntityID       string    `json:"entityId,omitempty"`       // The entity ID involved This is the unique identifier for the entity causing the incident.
	IncidentNumber string    `json:"incidentNumber,omitempty"` // The incident number as shown in VictorOps Multiple values and ranges are allowed: 4,5,20:50
	StartedAfter   time.Time `json:"startedAfter,omitempty"`   // Return incidents started after this timestamp Specify the timestamp in ISO8601 format
	StartedBefore  time.Time `json:"startedBefore,omitempty"`  // Find incidents started before this timestamp Specify the timestamp in ISO8601 format
	Host           string    `json:"host,omitempty"`           // The host involved in the incident Multiple values can be separated with commas
	Service        string    `json:"service,omitempty"`        // The service involved in the incident (if any) Multiple values can be separated with commas
	CurrentPhase   string    `json:"currentPhase,omitempty"`   // The current phase of the incident "resolved", "triggered" or "acknowledged". Multiple values can be separated with commas
	RoutingKey     string    `json:"routingKey,omitempty"`     // The original routing of the incident
}

// IncidentsHistory response of the GetIncidentsHistory method
type IncidentsHistory struct {
	Offset    int              `json:"offset"`    // The offset passed in the request
	Limit     int              `json:"limit"`     // The limit value passed in the request
	Total     int              `json:"total"`     // The total number of incidents available for this search
	Incidents []IncidentDetail `json:"incidents"` // An array of incident objects matching the search
}

// IncidentDetail details of an incident
type IncidentDetail struct {
	IncidentNumber string  `json:"incidentNumber"` // The VictorOps incident number
	StartTime      string  `json:"startTime"`      // The time the incident started
	CurrentPhase   string  `json:"currentPhase"`   // The current phase of the incident "resolved", "triggered" or "acknowledged"
	AlertCount     int     `json:"alertCount"`     // The number of alerts received for this incident
	LastAlertTime  string  `json:"lastAlertTime"`  // The time of the last alert received for the incident
	LastAlertID    string  `json:"lastAlertID"`    // The unique id of the last alert for the incident
	AckUserID      string  `json:"ackUserId"`      // The VictorOps user id of the user that acknowledged the incident
	AckTime        string  `json:"ackTime"`        // The time of the last acknowledgment of the incident
	EntityID       string  `json:"entityId"`       // The unique identification of the entity being monitored that caused the incident
	Host           string  `json:"host"`           // The host on which the incident occurred
	Service        string  `json:"service"`        // The service name causing the incident (if any)
	EndTime        string  `json:"endTime"`        // The time the incident ended
	User           AckUser `json:"ackUser"`
	EntityType     string  `json:"entityType"`     // The type of entity causing the incident (host/service)
	EntityDispName string  `json:"entityDispName"` // The display name of the entity causing the incident
	Teams          string  `json:"teams"`          // The teams that were paged for the incident (comma separated)
}

// IncidentDetailV2 details of an incident used in reporting API V2
type IncidentDetailV2 struct {
	IncidentNumber string               `json:"incidentNumber"` // The VictorOps incident number
	StartTime      string               `json:"startTime"`      // The time the incident started
	CurrentPhase   string               `json:"currentPhase"`   // The current phase of the incident "resolved", "triggered" or "acknowledged"
	AlertCount     int                  `json:"alertCount"`     // The number of alerts received for this incident
	LastAlertTime  string               `json:"lastAlertTime"`  // The time of the last alert received for the incident
	LastAlertID    string               `json:"lastAlertID"`    // The unique id of the last alert for the incident
	EntityID       string               `json:"entityId"`       // The unique identification of the entity being monitored that caused the incident
	Host           string               `json:"host"`           // The host on which the incident occurred
	Service        string               `json:"service"`        // The service name causing the incident (if any)
	PagedUsers     []string             `json:"pagedUsers"`     // The users that were paged for the incident
	PagedTeams     []string             `json:"pagedTeams"`     // The teams that were paged for the incident
	Transitions    []IncidentTransition `json:"transitions"`    // Transitions of the incident state over time
}

// IncidentTransition Transitions of the incident state over time
type IncidentTransition struct {
	Name     string    `json:"name"`     // The transition name
	At       time.Time `json:"at"`       // The time of the transition
	By       string    `json:"by"`       // The user that caused the transition (if any)
	Message  string    `json:"message"`  // The message entered by that user (if any)
	Manually bool      `json:"manually"` // If the incident transition was caused by a person
	AlertID  string    `json:"alertId"`  // The unique alert ID that caused the transition
	AlertURL string    `json:"alertUrl"` // A URL to retrieve the details of the alert that caused the transition
}

// AckUser user acked an incident
type AckUser struct {
	Email string `json:"email"`
	First string `json:"first"`
	Last  string `json:"last"`
}

// GetOncallLog returns a log of user shift changes for the specified team
func (client *Client) GetOncallLog(team string) (OncallLog, error) {
	var history OncallLog
	err := client.sendRequest("GET", "api-reporting/v1/team/"+team+"/oncall/log", nil, &history)
	return history, err
}

// GetIncidentsHistoryV1 retrieves incident history for your company, searching over date ranges and with filtering options.
// This is historical data, and may be up to 15 minutes behind real-time incident data.
// By default, only resolved incidents will be returned
func (client *Client) GetIncidentsHistoryV1(filter IncidentsSearchParams) ([]IncidentsHistory, error) {
	var hist []IncidentsHistory
	return hist, nil
}

// GetIncidentsHistoryV2 retrieves incident history for your company, searching over date ranges and with filtering options.
// This API may be called a maximum of once a minute.
// Incident requests are paginated with a offset and limit query string parameters.
// The query for incidents is run and offset records are skipped, after which limit records will be returned.
// The default offset is 0 and the default limit is 20. The maximum value allowed for limit is 100.
// On return, the total number of records available for that query will be returned in the payload as 'total'.
func (client *Client) GetIncidentsHistoryV2(filter IncidentsSearchParams) ([]IncidentDetailV2, error) {
	var hist []IncidentDetailV2
	return hist, nil
}
