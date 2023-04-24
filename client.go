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
	successCode           = 1100
	defaultTextDetectType = "DEFAULT_FRUAD"
)

var regionGatewayMap = map[string]map[string]string{
	// 文本
	"text": {
		"default": "http://api-text-bj.fengkongcloud.com/text/v4",   // 北京默认
		"bj":      "http://api-text-bj.fengkongcloud.com/text/v4",   // 北京
		"sh":      "http://api-text-sh.fengkongcloud.com/text/v4",   // 上海
		"gz":      "http://api-text-gz.fengkongcloud.com/text/v4",   // 广州
		"fjny":    "http://api-text-fjny.fengkongcloud.com/text/v4", // 美国（弗吉尼亚）
		"xjp":     "http://api-text-xjp.fengkongcloud.com/text/v4",  // 新加坡
	},
	"image": {
		"default": "http://api-img-bj.fengkongcloud.com/image/v4",  // 北京默认
		"bj":      "http://api-img-bj.fengkongcloud.com/image/v4",  // 北京
		"sh":      "http://api-img-sh.fengkongcloud.com/image/v4",  // 上海
		"gg":      "http://api-img-gg.fengkongcloud.com/image/v4",  // 硅谷
		"yd":      "http://api-img-yd.fengkongcloud.com/image/v4",  // 印度
		"xjp":     "http://api-img-xjp.fengkongcloud.com/image/v4", // 新加坡
	},
	"image_query": {
		"default": "http://api-img-active-query.fengkongcloud.com/v4/image/query", // 北京默认
		"bj":      "http://api-img-active-query.fengkongcloud.com/v4/image/query", // 北京
	},
	"audio": {
		"default": "http://api-audio-sh.fengkongcloud.com/audio/v4",  // 上海默认
		"sh":      "http://api-audio-sh.fengkongcloud.com/audio/v4",  // 上海
		"gg":      "http://api-audio-gg.fengkongcloud.com/audio/v4",  // 硅谷
		"xjp":     "http://api-audio-xjp.fengkongcloud.com/audio/v4", // 新加坡
	},
	"audio_query": {
		"default": "http://api-audio-sh.fengkongcloud.com/query_audio/v4", // 上海默认
		"sh":      "http://api-audio-sh.fengkongcloud.com/query_audio/v4", // 上海
		"gg":      "http://api-audio-gg.fengkongcloud.com/query_audio/v4", // 硅谷
	},
	"video": {
		"default": "http://api-video-bj.fengkongcloud.com/video/v4",  // 北京默认
		"bj":      "http://api-video-bj.fengkongcloud.com/video/v4",  // 北京
		"sh":      "http://api-video-sh.fengkongcloud.com/video/v4",  // 上海
		"gg":      "http://api-video-gg.fengkongcloud.com/video/v4",  // 硅谷
		"yd":      "http://api-video-yd.fengkongcloud.com/video/v4",  // 印度
		"xjp":     "http://api-video-xjp.fengkongcloud.com/video/v4", // 新加坡
	},
	"video_query": {
		"default": "http://api-video-bj.fengkongcloud.com/video/query/v4",  // 北京默认
		"bj":      "http://api-video-bj.fengkongcloud.com/video/query/v4",  // 北京
		"sh":      "http://api-video-sh.fengkongcloud.com/video/query/v4",  // 上海
		"gg":      "http://api-video-gg.fengkongcloud.com/video/query/v4",  // 硅谷
		"yd":      "http://api-video-yd.fengkongcloud.com/video/query/v4",  // 印度
		"xjp":     "http://api-video-xjp.fengkongcloud.com/video/query/v4", // 新加坡
	},
	"event": {
		"default": "http://api-skynet-bj.fengkongcloud.com/v4/event",   // 北京默认
		"bj":      "http://api-skynet-bj.fengkongcloud.com/v4/event",   // 北京
		"sh":      "http://api-skynet-fjny.fengkongcloud.com/v4/event", // 美国（弗吉尼亚）
		"gg":      "http://api-skynet-xjp.fengkongcloud.com/v4/event",  // 新加坡
		"yd":      "http://api-skynet-eur.fengkongcloud.com/v4/event",  // 欧洲（法兰克福
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
		region:    "default",
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
func (c *Client) SkyNetEvent(eventId string, data map[string]any) (*SkyNetEventResp, error) {
	// 创建一个 SkyNetEventReq 结构体，用于存储请求所需的公共数据和传入的 data 参数
	reqData := SkyNetEventReq{
		CommonReq: c.getCommonReq(eventId),
		Data:      data,
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request("event", reqData)
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

func (c *Client) TextDetect(eventId, detectType string, data map[string]any) (*TextDetectResp, error) {
	if detectType == "" {
		detectType = defaultTextDetectType
	}
	// 向指定的接口发送请求，并接收响应
	response, err := c.request("text", TextDetectReq{
		CommonReq: c.getCommonReq(eventId),
		Type:      detectType,
		Data:      data,
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

func (c *Client) ImageDetect() {

}
func (c *Client) AudioDetect() {

}
func (c *Client) VideoDetect() {

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
