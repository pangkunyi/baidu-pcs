package pcs

import (
	"net/url"
)

type OauthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (this *OauthError) HasError() bool {
	return this.Error != ""
}

/*
 *设备授权码回复
 */
type DeviceCodeResp struct {
	OauthError
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUrl string `json:"verification_url"`
	QrCodeUrl       string `json:"qrcode_url"`
	ExpiresIn       int64  `json:"expires_in"`
	Interval        int64  `json:"interval"`
}

type AccessTokenResp struct {
	OauthError
	AccessToken   string `json:"access_token"`   //要获取的Access Token
	ExpiresIn     int64  `json:"expires_in"`     //Access Token的有效期，以秒为单位；请参考“Access Token生命周期方案”
	RefreshToken  string `json:"refresh_token"`  //用于刷新Access Token 的 Refresh Token,所有应用都会返回该参数；（10年的有效期）
	Scope         string `json:"scope"`          //Access Token最终的访问范围，即用户实际授予的权限列表（用户在授权页面时，有可能会取消掉某些请求的权限），关于权限的具体信息参考“权限列表”一节；
	SessionKey    string `json:"session_key"`    //session_key：基于http调用Open API时所需要的Session Key，其有效期与Access Token一致；
	SessionSecret string `json:"session_secret"` //session_secret：基于http调用Open API时计算参数签名用的签名密钥。
}

/*
 *request device code
 */
func ReqDeviceCode() (dcr *DeviceCodeResp, err error) {
	data := make(url.Values)
	data.Add("client_id", CLIENT_ID)
	data.Add("response_type", "device_code")
	data.Add("scope", "basic,netdisk")
	dcr = new(DeviceCodeResp)
	err = postJson(DEVICE_CODE_URL, data, dcr)
	return
}

/*
 *request access token
 */
func ReqAccessToken(deviceCode string) (atr *AccessTokenResp, err error) {
	data := make(url.Values)
	data.Add("grant_type", "device_token")
	data.Add("code", deviceCode)
	data.Add("client_id", CLIENT_ID)
	data.Add("client_secret", SECRET_ID)
	atr = new(AccessTokenResp)
	err = postJson(ACCESS_TOKEN_URL, data, atr)
	return
}
