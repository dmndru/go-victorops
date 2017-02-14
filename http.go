package victorops

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (api *Client) sendRequest(method, path string, values []byte, intf interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(method, victorOpsAPI+path, bytes.NewBuffer(values))
	if err != nil {
		return err
	}
	req.Header.Add("X-VO-Api-Id", api.config.id)
	req.Header.Add("X-VO-Api-Key", api.config.key)

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if api.debug {
		log.Println(string(data))
	}

	if resp.StatusCode != 200 {
		switch resp.StatusCode {
		case 400:
			err = errors.New("Problem with the request arguments")
		case 401:
			err = errors.New("Authentication parameters missing")
		case 403:
			err = errors.New("Authentication failed or rate-limit reached")
		case 404:
			err = errors.New("Object not found")
		case 421:
			err = errors.New("You have reached your user limit")
		case 422:
			err = errors.New("Username or email is unavailable, or you have reached your user limit")
		case 500:
			err = errors.New("Internal server error")
		default:
			err = errors.New("Unknown response code: " + strconv.Itoa(resp.StatusCode))
		}
		return err
	}

	err = json.Unmarshal(data, &intf)
	if err != nil {
		return err
	}
	return nil
}
