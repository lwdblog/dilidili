// Code generated by liverpcgen, DO NOT EDIT.
// source: *.proto files under this directory
// If you want to change this file, Please see README in go-common/app/tool/liverpc/protoc-gen-liverpc/
package liverpc

import (
	"go-common/app/service/live/live_user/api/liverpc/v1"
	"go-common/library/net/rpc/liverpc"
)

// Client that represents a liverpc live_user service api
type Client struct {
	cli *liverpc.Client
	// V1Guard presents the controller in liverpc
	V1Guard v1.Guard
	// V1Note presents the controller in liverpc
	V1Note v1.Note
	// V1RoomAdmin presents the controller in liverpc
	V1RoomAdmin v1.RoomAdmin
	// V1UserSetting presents the controller in liverpc
	V1UserSetting v1.UserSetting
}

// DiscoveryAppId the discovery id is not the tree name
var DiscoveryAppId = "live.liveuser"

// New a Client that represents a liverpc live.liveuser service api
// conf can be empty, and it will use discovery to find service by default
// conf.AppID will be overwrite by a fixed value DiscoveryAppId
// therefore is no need to set
func New(conf *liverpc.ClientConfig) *Client {
	if conf == nil {
		conf = &liverpc.ClientConfig{}
	}
	conf.AppID = DiscoveryAppId
	var realCli = liverpc.NewClient(conf)
	cli := &Client{cli: realCli}
	cli.clientInit(realCli)
	return cli
}

func (cli *Client) GetRawCli() *liverpc.Client {
	return cli.cli
}

func (cli *Client) clientInit(realCli *liverpc.Client) {
	cli.V1Guard = v1.NewGuardRpcClient(realCli)
	cli.V1Note = v1.NewNoteRpcClient(realCli)
	cli.V1RoomAdmin = v1.NewRoomAdminRpcClient(realCli)
	cli.V1UserSetting = v1.NewUserSettingRpcClient(realCli)
}
