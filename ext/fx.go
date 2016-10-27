package ext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/reechou/real-wx/config"
)

const (
	FX_RSP_RESULT_OK = iota
)

const (
	CREATE_ACCOUNT_URI = "/fx/create_fx_account"
)

type FxSystem struct {
	cfg *config.Config

	client *http.Client
}

func NewFxSystem(cfg *config.Config) *FxSystem {
	return &FxSystem{
		cfg:    cfg,
		client: &http.Client{},
	}
}

func (fs *FxSystem) AddUser(info *CreateFxAccountReq) error {
	u := fs.cfg.FxSystemInfo.HostURL + CREATE_ACCOUNT_URI
	body, err := json.Marshal(info)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest("POST", u, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	rsp, err := fs.client.Do(httpReq)
	defer func() {
		if rsp != nil {
			rsp.Body.Close()
		}
	}()
	if err != nil {
		return err
	}
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	var response FxResponse
	err = json.Unmarshal(rspBody, &response)
	if err != nil {
		return err
	}
	if response.Code != FX_RSP_RESULT_OK {
		logrus.Errorf("add user error: %s", response.Msg)
		return fmt.Errorf("add user error: %s", response.Msg)
	}

	return nil
}
