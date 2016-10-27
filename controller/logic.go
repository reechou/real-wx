package controller

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"github.com/reechou/real-wx/config"
	"github.com/reechou/real-wx/ext"
	"github.com/reechou/real-wx/utils"
	"github.com/reechou/real-wx/wx_api"
)

type WXLogic struct {
	cfg        *config.Config
	msgHandler core.Handler
	msgServer  *core.Server
	api        *wx_api.WXApi
	fxExt      *ext.FxSystem
}

func NewWXLogic(cfg *config.Config) *WXLogic {
	wxl := &WXLogic{
		cfg: cfg,
	}
	wxl.init()

	wxl.api = wx_api.NewWXApi(cfg)
	wxl.fxExt = ext.NewFxSystem(cfg)

	return wxl
}

func (wxl *WXLogic) init() {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(wxl.defaultMsgHandler)
	mux.DefaultEventHandleFunc(wxl.defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, wxl.textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, wxl.menuClickEventHandler)
	wxl.msgHandler = mux
	wxl.msgServer = core.NewServer(config.WxOriId, config.WxAppId, config.WxToken, config.WxEncodedAESKey, wxl.msgHandler, nil)

	http.HandleFunc("/wx_callback", wxl.wxCallbackHandler)
}

func (wxl *WXLogic) wxCallbackHandler(w http.ResponseWriter, r *http.Request) {

}

func (wxl *WXLogic) textMsgHandler(ctx *core.Context) {
	logrus.Debugf("收到文本消息:%s", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func (wxl *WXLogic) defaultMsgHandler(ctx *core.Context) {
	logrus.Debugf("收到消息:%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func (wxl *WXLogic) menuClickEventHandler(ctx *core.Context) {
	logrus.Debugf("收到菜单 click 事件:%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func (wxl *WXLogic) defaultEventHandler(ctx *core.Context) {
	logrus.Debugf("收到事件:%s\n", ctx.MsgPlaintext)

	switch ctx.MixedMsg.EventType {
	case WX_EVENT_SUBSCRIBE:
		wxl.addUser(ctx)
	case WX_EVENT_SCAN:
	}

	ctx.NoneResponse()
}

func (wxl *WXLogic) Run() {
	if wxl.cfg.Debug {
		utils.EnableDebug()
	}
	logrus.Info("wx connect starting..")
	logrus.Infoln(http.ListenAndServe(wxl.cfg.Host, nil))
}
