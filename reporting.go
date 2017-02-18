package victorops

import "time"

type OncallLog struct {
	TeamSlug string    `json:"teamSlug"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	UserLogs []UserLog `json:"userLogs"`
}

type UserLog struct {
	UserID        string       `json:"userId"`
	Total         HoursMinutes `json:"total"`
	AdjustedTotal HoursMinutes `json:"adjustedTotal"`
	Logs          []Log        `json:"log"`
}

type HoursMinutes struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type Log struct {
	On       string       `json:"on"`
	Off      string       `json:"off"`
	Duration HoursMinutes `json:"duration"`
}

// GetOncallLog returns a log of user shift changes for the specified team
func (client *Client) GetOncallLog(team string) (OncallLog, error) {
	var history OncallLog
	err := client.sendRequest("GET", "api-reporting/v1/team/"+team+"/oncall/log", nil, &history)
	return history, err
}

/*
{
  "teamSlug": "string",
  "start": "2017-02-15T20:54:08.831Z",
  "end": "2017-02-15T20:54:08.831Z",
  "userLogs": [
    {
      "userId": "string",
      "total": {
        "hours": 0,
        "minutes": 0
      },
      "adjustedTotal": {
        "hours": 0,
        "minutes": 0
      },
      "log": [
        {
          "on": "string",
          "off": "string",
          "duration": {
            "hours": 0,
            "minutes": 0
          }
        }
      ]
    }
  ]
}
*/
