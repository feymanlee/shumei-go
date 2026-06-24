/**
 * @Author: lifameng@changba.com
 * @Description:
 * @File:  client
 * @Date: 2023/4/3 15:13
 */

package shumei

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	successCode                 = 1100
	defaultTextDetectType       = "TEXTRISK"
	defaultImageDetectType      = "POLITY_EROTIC_VIOLENT_QRCODE_ADVERT_IMGTEXTRISK"
	defaultAudioDetectType      = "POLITY_EROTIC_MOAN_ADVERT"
	defaultVideoImgDetectType   = "POLITY_EROTIC_VIOLENT_QRCODE_ADVERT_IMGTEXTRISK"
	defaultVideoAudioDetectType = "POLITICS_PORN_AD_MOAN_ABUSE_ANTHEN_AUDIOPOLITICAL"
)

// 定义地区常量集合
const (
	RegionDefault       = "default" // 默认地区
	RegionBeijing       = "bj"      // 北京地区
	RegionShanghai      = "sh"      // 上海地区
	RegionGuangzhou     = "gz"      // 广州地区
	RegionVirginia      = "fjny"    // 弗吉尼亚地区
	RegionSingapore     = "xjp"     // 新加坡地区
	RegionSiliconValley = "gg"      // 硅谷地区
	RegionIndia         = "yd"      // 印度地区
	RegionFrankfurt     = "eur"     // 欧洲（法兰克福）地区
)

const (
	actionTypeText       = "text"
	actionTypeImage      = "image"
	actionTypeImageQuery = "image_query"
	actionTypeAudioSync  = "audio_sync"
	actionTypeAudioAsync = "audio_async"
	actionTypeAudioQuery = "audio_query"
	actionTypeVideoAsync = "video_async"
	actionTypeVideoQuery = "video_query"
	actionTypeEvent      = "event"
)

var postJSON = func(endpoint string, request *resty.Request) (*resty.Response, error) {
	return request.Post(endpoint)
}

var actionRegionEndpoints = map[string]map[string]string{
	// 文本
	actionTypeText: {
		RegionDefault:   "https://api-text-bj.fengkongcloud.com/text/v4",   // 北京默认
		RegionBeijing:   "https://api-text-bj.fengkongcloud.com/text/v4",   // 北京
		RegionShanghai:  "https://api-text-sh.fengkongcloud.com/text/v4",   // 上海
		RegionGuangzhou: "https://api-text-gz.fengkongcloud.com/text/v4",   // 广州
		RegionVirginia:  "https://api-text-fjny.fengkongcloud.com/text/v4", // 美国（弗吉尼亚）
		RegionSingapore: "https://api-text-xjp.fengkongcloud.com/text/v4",  // 新加坡
	},
	actionTypeImage: {
		RegionDefault:       "https://api-img-bj.fengkongcloud.com/image/v4",  // 北京默认
		RegionBeijing:       "https://api-img-bj.fengkongcloud.com/image/v4",  // 北京
		RegionShanghai:      "https://api-img-sh.fengkongcloud.com/image/v4",  // 上海
		RegionSiliconValley: "https://api-img-gg.fengkongcloud.com/image/v4",  // 硅谷
		RegionSingapore:     "https://api-img-xjp.fengkongcloud.com/image/v4", // 新加坡
	},
	actionTypeImageQuery: {
		RegionDefault: "https://api-img-active-query.fengkongcloud.com/v4/image/query", // 北京默认
		RegionBeijing: "https://api-img-active-query.fengkongcloud.com/v4/image/query", // 北京
	},
	actionTypeAudioSync: {
		RegionDefault:   "https://api-audio-sh.fengkongcloud.com/audiomessage/v4",  // 上海默认
		RegionShanghai:  "https://api-audio-sh.fengkongcloud.com/audiomessage/v4",  // 上海
		RegionSingapore: "https://api-audio-xjp.fengkongcloud.com/audiomessage/v4", // 新加坡
	},
	actionTypeAudioAsync: {
		RegionDefault:       "https://api-audio-sh.fengkongcloud.com/audio/v4",  // 上海默认
		RegionShanghai:      "https://api-audio-sh.fengkongcloud.com/audio/v4",  // 上海
		RegionSiliconValley: "https://api-audio-gg.fengkongcloud.com/audio/v4",  // 硅谷
		RegionSingapore:     "https://api-audio-xjp.fengkongcloud.com/audio/v4", // 新加坡
	},
	actionTypeAudioQuery: {
		RegionDefault:       "https://api-audio-sh.fengkongcloud.com/query_audio/v4",  // 上海默认
		RegionShanghai:      "https://api-audio-sh.fengkongcloud.com/query_audio/v4",  // 上海
		RegionSiliconValley: "https://api-audio-gg.fengkongcloud.com/query_audio/v4",  // 硅谷
		RegionSingapore:     "https://api-audio-xjp.fengkongcloud.com/query_audio/v4", // 新加坡
	},
	actionTypeVideoAsync: {
		RegionDefault:   "https://api-video-bj.fengkongcloud.com/video/v4",  // 北京默认
		RegionBeijing:   "https://api-video-bj.fengkongcloud.com/video/v4",  // 北京
		RegionShanghai:  "https://api-video-sh.fengkongcloud.com/video/v4",  // 上海
		RegionSingapore: "https://api-video-xjp.fengkongcloud.com/video/v4", // 新加坡
	},
	actionTypeVideoQuery: {
		RegionDefault:   "https://api-video-bj.fengkongcloud.com/video/query/v4",  // 北京默认
		RegionBeijing:   "https://api-video-bj.fengkongcloud.com/video/query/v4",  // 北京
		RegionShanghai:  "https://api-video-sh.fengkongcloud.com/video/query/v4",  // 上海
		RegionSingapore: "https://api-video-xjp.fengkongcloud.com/video/query/v4", // 新加坡
	},
	actionTypeEvent: {
		RegionDefault:   "https://api-skynet-bj.fengkongcloud.com/v4/event",   // 北京默认
		RegionBeijing:   "https://api-skynet-bj.fengkongcloud.com/v4/event",   // 北京
		RegionVirginia:  "https://api-skynet-fjny.fengkongcloud.com/v4/event", // 美国（弗吉尼亚）
		RegionSingapore: "https://api-skynet-xjp.fengkongcloud.com/v4/event",  // 新加坡
		RegionFrankfurt: "https://api-skynet-eur.fengkongcloud.com/v4/event",  // 欧洲（法兰克福）
	},
}

type Client struct {
	endpoint                 string
	timeout                  time.Duration
	region                   string
	accessKey                string
	appID                    string
	useDefaultRegionFallback bool
}

func NewClient(appID, accessKey string, options ...ClientOption) (*Client, error) {
	c := defaultClient(appID, accessKey)

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func defaultClient(appID, accessKey string) *Client {
	c := &Client{
		accessKey: accessKey,
		appID:     appID,
		region:    RegionDefault,
		timeout:   time.Second * 3, // 默认超时时间为 3s
	}

	return c
}

// SkyNetEvent
// Client 是一个自定义结构体，表示客户端对象。
// SkyNetEvent 是 Client 的一个方法，它接收三个参数：
// appID (string) - 应用的ID
// eventID (string) - 事件的ID
// data (ReqSkyNetEventData) - 一个结构体，包含请求中的相关数据
// 返回两个值：
// *SkyNetEventResp - 结构体指针，包含处理请求后的响应数据
// error - 如果有错误发生，将返回一个错误对象
func (c *Client) SkyNetEvent(eventID string, req SkyNetEventReq) (*SkyNetEventResp, error) {
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeEvent, struct {
		CommonReq
		SkyNetEventReq
	}{
		c.getCommonReq(eventID),
		req,
	})
	// 如果请求发生错误，返回错误对象
	if err != nil {
		return nil, err
	}
	// 创建一个 SkyNetEventResp 结构体，用于存储响应数据
	result := SkyNetEventResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, fmt.Errorf("sky net event request failed, code[%v], message[%s], requestID[%s]", result.Code, result.Message, result.RequestID)
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) TextDetect(eventID string, req TextDetectReq) (*TextDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultTextDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeText, struct {
		CommonReq
		TextDetectReq
	}{
		c.getCommonReq(eventID),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := TextDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, fmt.Errorf("text detect request failed, code[%v], message[%s], requestID[%s]", result.Code, result.Message, result.RequestID)
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) ImageDetect(eventID string, req ImageDetectReq) (*ImageDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultImageDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeImage, struct {
		CommonReq
		ImageDetectReq
	}{
		c.getCommonReq(eventID),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := ImageDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, fmt.Errorf("image detect request failed, code[%v], message[%s], requestID[%s]", result.Code, result.Message, result.RequestID)
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) AudioSyncDetect(eventID string, req AudioSyncDetectReq) (*AudioSyncDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultAudioDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeAudioSync, struct {
		CommonReq
		AudioSyncDetectReq
	}{
		c.getCommonReq(eventID),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := AudioSyncDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, fmt.Errorf("audio sync detect request failed, code[%v], message[%s], requestID[%s]", result.Code, result.Message, result.RequestID)
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) AudioAsyncDetect(eventID string, req AudioAsyncDetectReq) (*AudioAsyncDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultAudioDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeAudioAsync, struct {
		CommonReq
		AudioAsyncDetectReq
	}{
		c.getCommonReq(eventID),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := AudioAsyncDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, fmt.Errorf("audio async detect request failed, code[%v], message[%s], requestID[%s]", result.Code, result.Message, result.RequestID)
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) VideoAsyncDetect(eventID string, req VideoAsyncDetectReq) (*VideoAsyncDetectResp, error) {
	if req.ImgType == "" {
		req.ImgType = defaultVideoImgDetectType
	}
	if req.AudioType == "" {
		req.AudioType = defaultVideoAudioDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeVideoAsync, struct {
		CommonReq
		VideoAsyncDetectReq
	}{
		c.getCommonReq(eventID),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := VideoAsyncDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, fmt.Errorf("video async detect request failed, code[%v], message[%s], requestID[%s]", result.Code, result.Message, result.RequestID)
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) AccessKey() string {
	return c.accessKey
}

func (c *Client) SetAccessKey(accessKey string) {
	c.accessKey = accessKey
}

func (c *Client) Timeout() time.Duration {
	return c.timeout
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *Client) Region() string {
	return c.region
}

func (c *Client) SetRegion(region string) {
	c.region = region
}

// 执行请求
func (c *Client) request(actionType string, payload interface{}) (*resty.Response, error) {
	endpoint, err := c.getEndpoint(actionType)
	if err != nil {
		return nil, err
	}
	request := resty.New().
		SetTimeout(c.timeout).
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload)
	response, err := postJSON(endpoint, request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode() < http.StatusOK || response.StatusCode() >= http.StatusMultipleChoices {
		return nil, newHTTPError(response)
	}
	return response, nil
}

type httpErrorBody struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

func newHTTPError(response *resty.Response) error {
	body := httpErrorBody{}
	if err := json.Unmarshal(response.Body(), &body); err == nil && (body.Code != 0 || body.Message != "" || body.RequestID != "") {
		return fmt.Errorf("http status[%d], code[%v], message[%s], requestID[%s]", response.StatusCode(), body.Code, body.Message, body.RequestID)
	}

	return fmt.Errorf("http status[%d], body[%s]", response.StatusCode(), string(response.Body()))
}

// getEndpoint 根据事件类型和地区获取对应的 API 网关
func (c *Client) getEndpoint(actionType string) (string, error) {
	var endpoint string
	endpoints, ok := actionRegionEndpoints[actionType]
	if !ok {
		return endpoint, fmt.Errorf("action type[%v] not found", actionType)
	}
	endpoint, ok = endpoints[c.region]
	if !ok && c.useDefaultRegionFallback {
		endpoint, ok = endpoints[RegionDefault]
	}
	if !ok {
		return endpoint, fmt.Errorf("action type[%v] region[%v] endpoint not found", actionType, c.region)
	}
	return endpoint, nil
}

// getCommonReq 获取公共请求参数
func (c *Client) getCommonReq(eventID string) CommonReq {
	return CommonReq{
		AccessKey: c.accessKey,
		AppID:     c.appID,
		EventID:   eventID,
	}
}
