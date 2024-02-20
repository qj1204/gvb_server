package qq

import (
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type QQInfo struct {
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Avatar   string `json:"figureurl_qq"`
	OpenID   string `json:"open_id"`
}

type QQLogin struct {
	appID       string
	appKey      string
	redirect    string
	code        string
	accessToken string
	openID      string
}

func NewQQLogin(code string) (qqInfo QQInfo, err error) {
	qqLogin := &QQLogin{
		appID:    global.Config.QQ.AppID,
		appKey:   global.Config.QQ.Key,
		redirect: global.Config.QQ.Redirect,
		code:     code,
	}
	err = qqLogin.GetAccessToken()
	if err != nil {
		return qqInfo, err
	}
	err = qqLogin.GetOpenID()
	if err != nil {
		return qqInfo, err
	}
	qqInfo, err = qqLogin.GetUserInfo()
	if err != nil {
		return qqInfo, err
	}
	qqInfo.OpenID = qqLogin.openID
	return qqInfo, nil
}

// GetAccessToken 获取access_token
func (this *QQLogin) GetAccessToken() error {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", this.appID)
	params.Add("client_secret", this.appKey)
	params.Add("code", this.code)
	params.Add("redirect_uri", this.redirect)
	// 获取access_token
	u, err := url.Parse("https://graph.qq.com/oauth2.0/token")
	if err != nil {
		return err
	}
	u.RawQuery = params.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	qs, err := parseQS(res.Body)
	if err != nil {
		return err
	}
	this.accessToken = qs[`access_token`][0]
	return nil
}

// GetOpenID 获取open_id
func (this *QQLogin) GetOpenID() error {
	u, err := url.Parse(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", this.accessToken))
	if err != nil {
		return err
	}
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	openID, err := getOpenID(res.Body)
	if err != nil {
		return err
	}
	this.openID = openID
	return nil
}

// GetUserInfo 获取用户信息
func (this *QQLogin) GetUserInfo() (qqInfo QQInfo, err error) {
	params := url.Values{}
	params.Add("access_token", this.accessToken)
	params.Add("oauth_consumer_key", this.appID)
	params.Add("openid", this.openID)
	// 获取用户信息
	u, err := url.Parse("https://graph.qq.com/user/get_user_info")
	if err != nil {
		return qqInfo, err
	}
	u.RawQuery = params.Encode()
	res, err := http.Get(u.String())
	data, err := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &qqInfo)
	if err != nil {
		return qqInfo, err
	}
	return qqInfo, nil
}

// getOpenID 从http响应的body中获取openid
func getOpenID(r io.Reader) (string, error) {
	body := readAll(r)
	start := strings.Index(body, `"openid":"`) + len(`"openid":"`)
	if start == -1 {
		return "", fmt.Errorf("openid not found")
	}
	end := strings.Index(body[start:], `"`)
	if end == -1 {
		return "", fmt.Errorf("openid not found")
	}
	return body[start : start+end], nil
}

// parseQS 将http响应的body转换为map
func parseQS(r io.Reader) (val map[string][]string, err error) {
	val, err = url.ParseQuery(readAll(r))
	if err != nil {
		return nil, err
	}
	return val, nil
}

// readAll 读取http响应的body
func readAll(r io.Reader) string {
	data, err := io.ReadAll(r)
	if err != nil {
		global.Log.Fatal(err)
	}
	return string(data)
}
