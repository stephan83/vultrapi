package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"io/ioutil"
	"net/url"
)

func PostCreateSSHKey(c Client, APIKey, name, sshKey string) (SSHKeyId string, err error) {

	values := url.Values{
		"name":    {name},
		"ssh_key": {sshKey},
	}

	resp, err := c.PostForm(fmt.Sprintf("/sshkey/create?api_key=%s",
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
		Id string `json:"SSHKEYID"`
	}{}
	err = json.Unmarshal(body, &js)
	SSHKeyId = js.Id

	return
}
