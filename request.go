/**
 * @Author: lifameng@changba.com
 * @Description:
 * @File:  request.go
 * @Date: 2023/4/3 15:18
 */

package shumei

type CommonReq struct {
	AccessKey string `json:"accessKey"`
	EventID   string `json:"eventId"`
	AppID     string `json:"appId"`
}

type SkyNetEventDataProductReq struct {
	MerchantID   string `json:"merchantId"`
	ProductCount int    `json:"productCount"`
	ProductID    string `json:"productId"`
}

type TextDetectReq struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type ImageDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	Callback     string                 `json:"callback"`
	Data         map[string]interface{} `json:"data"`
}

type AudioSyncDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	ContentType  string                 `json:"contentType"`
	Content      string                 `json:"content"`
	Callback     string                 `json:"callback"`
	BtId         string                 `json:"btId"`
	Data         map[string]interface{} `json:"data"`
}

type AudioAsyncDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	ContentType  string                 `json:"contentType"`
	Content      string                 `json:"content"`
	Callback     string                 `json:"callback"`
	BtId         string                 `json:"btId"`
	Data         map[string]interface{} `json:"data"`
}

type VideoAsyncDetectReq struct {
	ImgType           string                 `json:"imgType"`
	ImgBusinessType   string                 `json:"imgBusinessType"`
	AudioType         string                 `json:"audioType"`
	AudioBusinessType string                 `json:"audioBusinessType"`
	Callback          string                 `json:"callback"`
	Data              map[string]interface{} `json:"data"`
}

// SkyNetEventReq 数美天网请求
// See doc：https://help.ishumei.com/docs/tw/marketing/newest/developDoc#data
type SkyNetEventReq struct {
	Data map[string]interface{}
}
