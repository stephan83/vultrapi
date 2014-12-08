package requests

import (
	"encoding/json"
	"errors"
	. "github.com/stephan83/vultrapi/types"
	"io/ioutil"
	"net/http"
)

func GetOS() (OS OSDict, err error) {
	resp, err := http.Get("https://api.vultr.com/v1/os/list")
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
