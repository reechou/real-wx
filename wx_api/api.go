package wx_api

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/user"
	"github.com/reechou/real-wx/config"
)

type WXApi struct {
	cfg               *config.Config
	accessTokenServer core.AccessTokenServer
	wxClient          *core.Client
}

func NewWXApi(cfg *config.Config) *WXApi {
	api := &WXApi{
		cfg: cfg,
	}
	api.accessTokenServer = core.NewDefaultAccessTokenServer(config.WxAppId, config.WxAppSecret, nil)
	api.wxClient = core.NewClient(api.accessTokenServer, nil)

	return api
}

func (api *WXApi) GetUserInfo(openId string) (*user.UserInfo, error) {
	return user.Get(api.wxClient, openId, "")
}
