package requests

import (
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"io/ioutil"
	"net/url"
)

func PostDestroySnapshot(c Client, APIKey string,
	snapshotId string) (err error) {

	values := url.Values{
		"SNAPSHOTID": {snapshotId},
	}

	resp, err := c.PostForm(fmt.Sprintf("/snapshot/destroy?api_key=%s",
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

	return
}
