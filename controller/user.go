package controller

import (
	"github.com/Sirupsen/logrus"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/reechou/real-wx/ext"
)

const (
	SCAN_PREFIX = "qrscene_"
)

func (wxl *WXLogic) addUser(ctx *core.Context) error {
	userInfo, err := wxl.api.GetUserInfo(ctx.MixedMsg.MsgHeader.FromUserName)
	if err != nil {
		logrus.Errorf("get user[%s] info error: %v", ctx.MixedMsg.MsgHeader.FromUserName, err)
		return err
	}
	req := &ext.CreateFxAccountReq{
		UnionId:   userInfo.UnionId,
		WXAccount: ctx.MixedMsg.MsgHeader.ToUserName,
		OpenId:    ctx.MixedMsg.MsgHeader.FromUserName,
		Name:      userInfo.Nickname,
	}
	if ctx.MixedMsg.EventKey != "" {
		req.Superior = ctx.MixedMsg.EventKey[len(SCAN_PREFIX):]
	}
	err = wxl.fxExt.AddUser(req)
	if err != nil {
		logrus.Errorf("fx ext add user error: %v", err)
		return err
	}
	return nil
}
