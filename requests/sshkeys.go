package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/types"
	"io/ioutil"
)

func GetSSHKeys(c Client, APIKey string) (SSHKeys SSHKeyMap, err error) {
	resp, err := c.Get(fmt.Sprintf("/sshkey/list?api_key=%s", APIKey))
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

	SSHKeys = SSHKeyMap{}
	err = json.Unmarshal(body, &SSHKeys)

	return
}
