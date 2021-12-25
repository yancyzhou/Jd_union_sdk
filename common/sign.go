package common

import (
	"github.com/yancyzhou/Jd_union_sdk/pkg"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

func (J *JdSdk) SetSignJointUrlParam() *JdSdk {
	J.SysParams.Timestamp = time.Now().In(pkg.Location).Format(pkg.TimeLayout)
	var values = reflect.ValueOf(J.SysParams)
	var keys = reflect.TypeOf(J.SysParams)
	var count = values.NumField()
	var slices = pkg.Persons{}
	for i := 0; i < count; i++ {
		value := values.Field(i)
		key := keys.Field(i)
		if key.Name == "Param_json" {
			key.Name = "360buy_param_json"
		}
		switch value.Kind() {
		case reflect.String:
			slices = append(slices, pkg.Person{strings.ToLower(key.Name), value.String()})
		case reflect.Int:
			slices = append(slices, pkg.Person{strings.ToLower(key.Name), value.String()})
		}
	}
	sort.Sort(slices)
	var builder strings.Builder
	u := url.Values{}
	builder.WriteString(J.AppSecret)
	for _, person := range slices {
		if person.Value == "" {
			continue
		}
		builder.WriteString(strings.ToLower(person.Key) + person.Value)
		u.Add(strings.ToLower(person.Key), person.Value)
	}
	builder.WriteString(J.AppSecret)
	//生成签名
	u.Add("sign", strings.ToUpper(pkg.Md5(builder.String())))
	//拼接参数
	J.SignAndUri = u.Encode()
	return J
}
