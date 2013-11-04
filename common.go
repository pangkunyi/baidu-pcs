package pcs

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

const (
	CFG_FILENAME       = `.baidu_pcs.cfg.json`
	AUTH_DATA_FILENAME = `.baidu_pcs.auth_data.json`
	PCS_URL            = `https://pcs.baidu.com/rest/2.0/pcs/file`
	D_PCS_URL          = `https://d.pcs.baidu.com/rest/2.0/pcs/file`
	DEVICE_CODE_URL    = `https://openapi.baidu.com/oauth/2.0/device/code`
	ACCESS_TOKEN_URL   = `https://openapi.baidu.com/oauth/2.0/token`
)

var (
	CLIENT_ID    string
	SECRET_ID    string
	OPEN_DIR     string
	ACCESS_TOKEN string
)

/**
 * 配置信息
 */
type Config struct {
	ClientId string `json:"client_id"`
	SecretId string `json:"secret_id"`
	OpenDir  string `json:"open_dir"`
}

/**
 * 获取Oauth授权后的授权信息
 */
type AuthData struct {
	AccessToken string `json:"access_token"`
}

type BasicFileReq struct {
	Method      string `pcs:"method"`       //required
	AccessToken string `pcs:"access_token"` //required, 开发者准入标识，HTTPS调用时必须使用。
	Path        string `pcs:"path"`         //required, 需要访问的文件，以/开头的绝对路径。
}

/**
 *初始化, 加载配置文件
 */
func init() {
	//加载配置文件
	body, err := ioutil.ReadFile(os.Getenv("HOME") + "/" + CFG_FILENAME)
	if err != nil {
		panic(err)
	}
	cfg := new(Config)
	err = json.Unmarshal(body, cfg)
	if err != nil {
		panic(err)
	}
	CLIENT_ID = cfg.ClientId
	SECRET_ID = cfg.SecretId
	OPEN_DIR = cfg.OpenDir

	//加载授权数据文件
	body, err = ioutil.ReadFile(os.Getenv("HOME") + "/" + AUTH_DATA_FILENAME)
	if err != nil {
		panic(err)
	}
	authData := new(AuthData)
	err = json.Unmarshal(body, authData)
	if err != nil {
		panic(err)
	}
	ACCESS_TOKEN = authData.AccessToken
}

/*
 * request url with params in get method, and write binary data to out writer
 */
func getData(_url string, params url.Values, w io.Writer) (err error) {
	if resp, err := http.Get(_url + "?" + params.Encode()); err == nil {
		defer resp.Body.Close()
		_, err = io.Copy(w, resp.Body)
	}
	return
}

/*
 * request url with params in get method, and reponse would be a json string
 */
func getJson(_url string, params url.Values, v interface{}) (err error) {
	if resp, err := http.Get(_url + "?" + params.Encode()); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			err = json.Unmarshal(body, v)
		}
	}
	return
}

/*
 * request url with params in post method, and reponse would be a json string
 */
func postJson(_url string, params url.Values, v interface{}) (err error) {
	if resp, err := http.PostForm(_url, params); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			err = json.Unmarshal(body, v)
		}
	}
	return
}

/**
 * parse a request object to url.Values, the request object's feild must be all string type and has the pcs tag
 */
func parseUrlValues(i interface{}) (values url.Values) {
	values = make(url.Values)
	_parseUrlValues(i, &values)
	return
}
func _parseUrlValues(i interface{}, values *url.Values) {
	v := reflect.ValueOf(i).Elem()
	t := v.Type()
	idx := 0
	for ; idx < v.NumField(); idx++ {
		f := v.Field(idx)
		tf := t.Field(idx)
		if tf.Anonymous {
			_parseUrlValues(f.Addr().Interface(), values)
		} else {
			tag := tf.Tag.Get("pcs")
			if tag != "" {
				val := f.Interface().(string)
				if val != "" {
					values.Add(tag, val)
				}
			}
		}
	}
	return
}
