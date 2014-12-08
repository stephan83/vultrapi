package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	SSHKeyId             int
	EnableAutoBackups    bool
}

func PostCreateServer(APIKey string, regionId, plandId, osId int,
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
		values["iso_id"] = []string{strconv.Itoa(options.ISOId)}
	}
	if options.ScriptId > 0 {
		values["script_id"] = []string{strconv.Itoa(options.ScriptId)}
	}
	if options.SnapshotId > 0 {
		values["snapshot_id"] = []string{strconv.Itoa(options.SnapshotId)}
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
	if options.SSHKeyId > 0 {
		values["ssh_key_id"] = []string{strconv.Itoa(options.SSHKeyId)}
	}
	if options.EnableAutoBackups {
		values["auto_backups"] = []string{"yes"}
	}

	resp, err := http.PostForm(fmt.Sprintf(
		"https://api.vultr.com/v1/server/create?api_key=%s",
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
