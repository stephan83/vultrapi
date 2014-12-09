package requests

import (
	"encoding/json"
	"errors"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/types"
	"io/ioutil"
)

func GetOS(c Client) (OS OSDict, err error) {
	resp, err := c.Get("/os/list")
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

	OS = OSDict{}
	err = json.Unmarshal(body, &OS)

	return
}
