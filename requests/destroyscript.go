package requests

import (
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"io/ioutil"
	"net/url"
	"strconv"
)

func PostDestroyScript(c Client, APIKey string, scriptId int) (err error) {

	values := url.Values{
		"SCRIPTID": {strconv.Itoa(scriptId)},
	}

	resp, err := c.PostForm(fmt.Sprintf("/startupscript/destroy?api_key=%s",
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
