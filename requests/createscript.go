package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	. "github.com/stephan83/vultrapi/types"
	"io/ioutil"
	"net/url"
)

func PostCreateScript(c Client, APIKey string, scriptType ScriptType, name, script string) (ScriptId string, err error) {

	values := url.Values{
		"type":   {scriptType.String()},
		"name":   {name},
		"script": {script},
	}

	resp, err := c.PostForm(fmt.Sprintf("/startupscript/create?api_key=%s",
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
		Id string `json:"SCRIPTID"`
	}{}
	err = json.Unmarshal(body, &js)
	ScriptId = js.Id

	return
}
