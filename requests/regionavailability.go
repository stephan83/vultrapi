package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"io/ioutil"
)

func GetRegionAvailability(c Client, regionId int) (plans []int, err error) {
	resp, err := c.Get(fmt.Sprintf("/regions/availability?DCID=%d", regionId))
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

	plans = []int{}
	err = json.Unmarshal(body, &plans)

	return
}
