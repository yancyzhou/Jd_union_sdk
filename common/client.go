package common

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	Host string = "https://api.jd.com/routerjson?"
)

type JDUnionConfig struct {
}

type SysParams struct {
	Method       string `json:"method"`
	App_key      string `json:"app_key"`
	Access_token string `json:"access_token"`
	Timestamp    string `json:"timestamp"`
	Format       string `json:"format"`
	V            string `json:"v"`
	Sign_method  string `json:"sign_method"`
	Param_json   string `json:"360buy_param_json"`
}

type JdSdk struct {
	SysParams  SysParams
	SignAndUri string
	RequestRul string
	JdAppKey   string
	AppSecret  string
}

func (J *JdSdk) NewContext() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	J.SysParams.V = "1.0"
	J.SysParams.Format = "json"
	J.SysParams.App_key = J.JdAppKey
	J.SysParams.Sign_method = "md5"
	J.SysParams.Timestamp = time.Now().In(location).Format("2006-01-02 15:04:05")
}

func StructToString(query interface{}) (res string, err error) {
	marshal, err := json.Marshal(query)
	if err != nil {
		return res, err
	}
	return string(marshal), nil
}

func (J *JdSdk) BodyBytes(query interface{}) []byte {
	J.NewContext()
	J.SysParams.Param_json, _ = StructToString(query)
	J.SetSignJointUrlParam()
	urls := strings.Builder{}
	urls.WriteString(Host)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	return body
}
