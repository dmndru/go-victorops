package victorops

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (api *Client) get(path string, values url.Values, intf interface{}, debug bool) error {
	client := &http.Client{
		Timeout: 10,
	}

	req, err := http.NewRequest("GET", victorOpsAPI+path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-VO-Api-Id", api.config.id)
	req.Header.Add("X-VO-Api-Key", api.config.key)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &intf)
	if err != nil {
		return err
	}
	return nil
}
