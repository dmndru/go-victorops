package victorops

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func (api *Client) get(path string, values url.Values, intf interface{}, debug bool) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
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
	if debug {
		log.Println(string(data))
	}
	err = json.Unmarshal(data, &intf)
	if err != nil {
		return err
	}
	return nil
}
