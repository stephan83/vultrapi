package requests

import (
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"io/ioutil"
	"net/url"
	"strconv"
)

func PostDestroyServer(c Client, APIKey string,
	serverId int) (err error) {

	values := url.Values{
		"SUBID": {strconv.Itoa(serverId)},
	}

	resp, err := c.PostForm(fmt.Sprintf("/server/destroy?api_key=%s",
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
