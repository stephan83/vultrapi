package requests

import (
	"encoding/json"
	"errors"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/types"
	"io/ioutil"
)

func GetRegions(c Client) (regions RegionMap, err error) {
	resp, err := c.Get("/regions/list")
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

	regions = RegionMap{}
	err = json.Unmarshal(body, &regions)

	return
}
