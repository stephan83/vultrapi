package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"io/ioutil"
	"net/url"
	"strconv"
)

func PostCreateSnapshot(c Client, APIKey string, serverId int, description string) (snapshotId string, err error) {

	values := url.Values{
		"SUBID": {strconv.Itoa(serverId)},
	}

	if description != "" {
		values["description"] = []string{description}
	}

	resp, err := c.PostForm(fmt.Sprintf("/snapshot/create?api_key=%s",
		APIKey), values)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode >= 400 {
		err = errors.New(string(body))
		return
	}

	js := struct {
		Id string `json:"SNAPSHOTID"`
	}{}
	err = json.Unmarshal(body, &js)
	snapshotId = js.Id

	return
}
