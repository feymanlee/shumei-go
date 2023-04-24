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
	CommonReq
	Type string `json:"type"`
	Data map[string]any
}

// SkyNetEventReq 数美天网请求
// See doc：https://help.ishumei.com/docs/tw/marketing/newest/developDoc#data
type SkyNetEventReq struct {
	CommonReq
	Data map[string]any
}