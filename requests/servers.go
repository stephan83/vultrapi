package requests

import (
	"encoding/json"
	"errors"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/types"
	"io/ioutil"
	"fmt"
)

func GetServers(c Client, APIKey string) (servers ServerDict, err error) {
	resp, err := c.Get(fmt.Sprintf("/server/list?api_key=%s", APIKey))
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

	servers = ServerDict{}
	err = json.Unmarshal(body, &servers)

	return
}