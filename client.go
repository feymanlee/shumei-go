/**
 * @Author: lifameng@changba.com
 * @Description:
 * @File:  client
 * @Date: 2023/4/3 15:13
 */

package shumei

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-zoox/fetch"
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

var regionGatewayMap = map[string]map[string]string{
	// 文本
	actionTypeText: {
		RegionDefault:   "http://api-text-bj.fengkongcloud.com/text/v4",   // 北京默认
		RegionBeijing:   "http://api-text-bj.fengkongcloud.com/text/v4",   // 北京
		RegionShanghai:  "http://api-text-sh.fengkongcloud.com/text/v4",   // 上海
		RegionGuangzhou: "http://api-text-gz.fengkongcloud.com/text/v4",   // 广州
		RegionVirginia:  "http://api-text-fjny.fengkongcloud.com/text/v4", // 美国（弗吉尼亚）
		RegionSingapore: "http://api-text-xjp.fengkongcloud.com/text/v4",  // 新加坡
	},
	actionTypeImage: {
		RegionDefault:       "http://api-img-bj.fengkongcloud.com/image/v4",  // 北京默认
		RegionBeijing:       "http://api-img-bj.fengkongcloud.com/image/v4",  // 北京
		RegionShanghai:      "http://api-img-sh.fengkongcloud.com/image/v4",  // 上海
		RegionSiliconValley: "http://api-img-gg.fengkongcloud.com/image/v4",  // 硅谷
		RegionIndia:         "http://api-img-yd.fengkongcloud.com/image/v4",  // 印度
		RegionSingapore:     "http://api-img-xjp.fengkongcloud.com/image/v4", // 新加坡
	},
	actionTypeImageQuery: {
		RegionDefault: "http://api-img-active-query.fengkongcloud.com/v4/image/query", // 北京默认
		RegionBeijing: "http://api-img-active-query.fengkongcloud.com/v4/image/query", // 北京
	},
	actionTypeAudioSync: {
		RegionDefault:  "http://api-audio-sh.fengkongcloud.com/audiomessage/v4", // 上海默认
		RegionShanghai: "http://api-audio-sh.fengkongcloud.com/audiomessage/v4", // 上海
	},
	actionTypeAudioAsync: {
		RegionDefault:       "http://api-audio-sh.fengkongcloud.com/audio/v4",  // 上海默认
		RegionShanghai:      "http://api-audio-sh.fengkongcloud.com/audio/v4",  // 上海
		RegionSiliconValley: "http://api-audio-gg.fengkongcloud.com/audio/v4",  // 硅谷
		RegionSingapore:     "http://api-audio-xjp.fengkongcloud.com/audio/v4", // 新加坡
	},
	actionTypeAudioQuery: {
		RegionDefault:       "http://api-audio-sh.fengkongcloud.com/query_audio/v4", // 上海默认
		RegionShanghai:      "http://api-audio-sh.fengkongcloud.com/query_audio/v4", // 上海
		RegionSiliconValley: "http://api-audio-gg.fengkongcloud.com/query_audio/v4", // 硅谷
	},
	actionTypeVideoAsync: {
		RegionDefault:       "http://api-video-bj.fengkongcloud.com/video/v4",  // 北京默认
		RegionBeijing:       "http://api-video-bj.fengkongcloud.com/video/v4",  // 北京
		RegionShanghai:      "http://api-video-sh.fengkongcloud.com/video/v4",  // 上海
		RegionSiliconValley: "http://api-video-gg.fengkongcloud.com/video/v4",  // 硅谷
		RegionIndia:         "http://api-video-yd.fengkongcloud.com/video/v4",  // 印度
		RegionSingapore:     "http://api-video-xjp.fengkongcloud.com/video/v4", // 新加坡
	},
	actionTypeVideoQuery: {
		RegionDefault:       "http://api-video-bj.fengkongcloud.com/video/query/v4",  // 北京默认
		RegionBeijing:       "http://api-video-bj.fengkongcloud.com/video/query/v4",  // 北京
		RegionShanghai:      "http://api-video-sh.fengkongcloud.com/video/query/v4",  // 上海
		RegionSiliconValley: "http://api-video-gg.fengkongcloud.com/video/query/v4",  // 硅谷
		RegionIndia:         "http://api-video-yd.fengkongcloud.com/video/query/v4",  // 印度
		RegionSingapore:     "http://api-video-xjp.fengkongcloud.com/video/query/v4", // 新加坡
	},
	actionTypeEvent: {
		RegionDefault:       "http://api-skynet-bj.fengkongcloud.com/v4/event",   // 北京默认
		RegionBeijing:       "http://api-skynet-bj.fengkongcloud.com/v4/event",   // 北京
		RegionShanghai:      "http://api-skynet-fjny.fengkongcloud.com/v4/event", // 美国（弗吉尼亚）
		RegionSiliconValley: "http://api-skynet-xjp.fengkongcloud.com/v4/event",  // 新加坡
		RegionIndia:         "http://api-skynet-eur.fengkongcloud.com/v4/event",  // 欧洲（法兰克福
	},
}

type Client struct {
	endpoint  string
	timeout   time.Duration
	region    string
	accessKey string
	appId     string
}

func NewClient(appId, accessKey string, options ...ClientOption) *Client {
	c := &Client{
		accessKey: accessKey,
		appId:     appId,
		region:    RegionDefault,
		timeout:   time.Second * 3, // 默认超时时间为 3s
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil
		}
	}

	return c
}

// SkyNetEvent
// Client 是一个自定义结构体，表示客户端对象。
// SkyNetEvent 是 Client 的一个方法，它接收三个参数：
// appId (string) - 应用的ID
// eventId (string) - 事件的ID
// data (ReqSkyNetEventData) - 一个结构体，包含请求中的相关数据
// 返回两个值：
// *SkyNetEventResp - 结构体指针，包含处理请求后的响应数据
// error - 如果有错误发生，将返回一个错误对象
func (c *Client) SkyNetEvent(eventId string, req SkyNetEventReq) (*SkyNetEventResp, error) {
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeEvent, struct {
		CommonReq
		SkyNetEventReq
	}{
		c.getCommonReq(eventId),
		req,
	})
	// 如果请求发生错误，返回错误对象
	if err != nil {
		return nil, err
	}
	// 创建一个 SkyNetEventResp 结构体，用于存储响应数据
	result := SkyNetEventResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, errors.New(fmt.Sprintf("sky net event request failed, code[%v],message[%s]", result.Code, result.Message))
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) TextDetect(eventId string, req TextDetectReq) (*TextDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultTextDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeText, struct {
		CommonReq
		TextDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := TextDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, errors.New(fmt.Sprintf("text detect request failed, code[%v],message[%s]", result.Code, result.Message))
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) ImageDetect(eventId string, req ImageDetectReq) (*ImageDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultImageDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeImage, struct {
		CommonReq
		ImageDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := ImageDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, errors.New(fmt.Sprintf("image detect request failed, code[%v],message[%s]", result.Code, result.Message))
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) AudioSyncDetect(eventId string, req AudioSyncDetectReq) (*AudioSyncDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultAudioDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeAudioSync, struct {
		CommonReq
		AudioSyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := AudioSyncDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, errors.New(fmt.Sprintf("audio sync detect request failed, code[%v],message[%s]", result.Code, result.Message))
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) AudioAsyncDetect(eventId string, req AudioAsyncDetectReq) (*AudioAsyncDetectResp, error) {
	if req.Type == "" {
		req.Type = defaultAudioDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request(actionTypeAudioAsync, struct {
		CommonReq
		AudioAsyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := AudioAsyncDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, errors.New(fmt.Sprintf("audio async detect request failed, code[%v],message[%s]", result.Code, result.Message))
	}
	// 如果一切正常，返回 result 结构体指针和空的错误对象
	return &result, nil
}

func (c *Client) VideoAsyncDetect(eventId string, req VideoAsyncDetectReq) (*VideoAsyncDetectResp, error) {
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
		c.getCommonReq(eventId),
		req,
	})
	if err != nil {
		return nil, err
	}
	result := VideoAsyncDetectResp{}
	// 将收到的 JSON 响应解析到 result 结构体中
	if err = response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	// 检查响应代码，如果不是预期的 successCode，返回错误信息
	if result.Code != successCode {
		return nil, errors.New(fmt.Sprintf("video async detect request failed, code[%v],message[%s]", result.Code, result.Message))
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
func (c *Client) request(shumeiEvent string, payload fetch.Body) (*fetch.Response, error) {
	endpoint, err := c.getEndpoint(shumeiEvent)
	if err != nil {
		return nil, err
	}
	return fetch.Post(endpoint, &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:    payload,
		Timeout: c.timeout,
	})
}

// getEndpoint 根据事件类型和地区获取对应的 API 网关
func (c *Client) getEndpoint(shumeiEvent string) (string, error) {
	var endpoint string
	endpoints, ok := regionGatewayMap[shumeiEvent]
	if !ok {
		return endpoint, fmt.Errorf("shumei event[%v] not found", shumeiEvent)
	}
	endpoint, ok = endpoints[c.region]
	if !ok {
		return endpoint, fmt.Errorf("shumei event[%v] region[%v] endpoint not found", shumeiEvent, c.region)
	}
	return endpoint, nil
}

// getCommonReq 获取公共请求参数
func (c *Client) getCommonReq(eventId string) CommonReq {
	return CommonReq{
		AccessKey: c.accessKey,
		AppID:     c.appId,
		EventID:   eventId,
	}
}
