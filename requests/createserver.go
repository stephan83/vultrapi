package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"io/ioutil"
	"net/url"
	"strconv"
)

type CreateServerOptions struct {
	IPXEChainURL         string
	ISOId                int
	ScriptId             int
	SnapshotId           int
	EnableIPV6           bool
	EnablePrivateNetwork bool
	Label                string
	SSHKeyId             string
	EnableAutoBackups    bool
}

func PostCreateServer(c Client, APIKey string, regionId, plandId, osId int,
	options CreateServerOptions) (serverId int, err error) {

	values := url.Values{
		"DCID":      {strconv.Itoa(regionId)},
		"VPSPLANID": {strconv.Itoa(plandId)},
		"OSID":      {strconv.Itoa(osId)},
	}

	if options.IPXEChainURL != "" {
		values["ipxe_chain_url"] = []string{options.IPXEChainURL}
	}
	if options.ISOId > 0 {
		values["ISOID"] = []string{strconv.Itoa(options.ISOId)}
	}
	if options.ScriptId > 0 {
		values["SCRIPTID"] = []string{strconv.Itoa(options.ScriptId)}
	}
	if options.SnapshotId > 0 {
		values["SNAPSHOTID"] = []string{strconv.Itoa(options.SnapshotId)}
	}
	if options.EnableIPV6 {
		values["enable_ipv6"] = []string{"yes"}
	}
	if options.EnablePrivateNetwork {
		values["enable_private_network"] = []string{"yes"}
	}
	if options.Label != "" {
		values["label"] = []string{options.Label}
	}
	if options.SSHKeyId != "" {
		values["SSHKEYID"] = []string{options.SSHKeyId}
	}
	if options.EnableAutoBackups {
		values["auto_backups"] = []string{"yes"}
	}

	resp, err := c.PostForm(fmt.Sprintf("/server/create?api_key=%s",
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

	server := struct {
		Id int `json:"SUBID,string"`
	}{}
	err = json.Unmarshal(body, &server)
	serverId = server.Id

	return
}
