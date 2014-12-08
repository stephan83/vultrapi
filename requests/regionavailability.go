package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRegionAvailability(regionId int) (plans []int, err error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://api.vultr.com/v1/regions/availability?DCID=%d",
		regionId))
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
