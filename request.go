/**
 * @Author: lifameng@changba.com
 * @Description:
 * @File:  request.go
 * @Date: 2023/4/3 15:18
 */

package shumei

// CommonReq 是所有数美接口共用的请求参数。
type CommonReq struct {
	// AccessKey 是数美分配的接口认证密钥。
	AccessKey string `json:"accessKey"`
	// EventID 是数美侧约定的事件或场景标识。
	EventID string `json:"eventId"`
	// AppID 是数美侧约定的应用标识。
	AppID string `json:"appId"`
}

// SkyNetEventDataProductReq 是天网事件中商品信息的辅助结构。
type SkyNetEventDataProductReq struct {
	// MerchantID 是商户标识。
	MerchantID string `json:"merchantId"`
	// ProductCount 是商品数量。
	ProductCount int `json:"productCount"`
	// ProductID 是商品标识。
	ProductID string `json:"productId"`
}

// TextDetectReq 是文本同步检测请求的业务参数。
type TextDetectReq struct {
	// Type 是文本检测风险类型，不传时 SDK 使用默认值。
	Type string `json:"type"`
	// AgentType 是文本 Agent 大模型检测的风险类型。
	AgentType string `json:"agentType,omitempty"`
	// TranslationTargetLang 是翻译目标语种，可选 zh、en，需联系数美开通。
	TranslationTargetLang string `json:"translationTargetLang,omitempty"`
	// AcceptLang 是返回标签语种，可选 zh、en。
	AcceptLang string `json:"acceptLang,omitempty"`
	// Callback 是回调地址。
	Callback string `json:"callback,omitempty"`
	// Data 是请求数据内容，常见字段包括 text、tokenId、lang 等。
	Data map[string]interface{} `json:"data"`
}

// ImageDetectReq 是图片同步检测请求的业务参数。
type ImageDetectReq struct {
	// Type 是图片检测风险类型，不传时 SDK 使用默认值。
	Type string `json:"type"`
	// BusinessType 是图片业务标签类型，和 Type 至少传一个。
	BusinessType string `json:"businessType"`
	// AgentType 是图片 Agent 大模型检测的风险类型。
	AgentType string `json:"agentType,omitempty"`
	// Callback 是异步场景下的回调地址；同步检测通常不需要传。
	Callback string `json:"callback"`
	// AcceptLang 是返回标签语种，可选 zh、en。
	AcceptLang string `json:"acceptLang,omitempty"`
	// Data 是请求数据内容，常见字段包括 img、tokenId、backupUrl 等。
	Data map[string]interface{} `json:"data"`
}

// AudioSyncDetectReq 是音频同步检测请求的业务参数。
type AudioSyncDetectReq struct {
	// Type 是音频检测风险类型，不传时 SDK 使用默认值。
	Type string `json:"type"`
	// BusinessType 是音频业务标签类型，和 Type 至少传一个。
	BusinessType string `json:"businessType"`
	// ContentType 是待检测音频内容格式，可选 URL、RAW。
	ContentType string `json:"contentType"`
	// Content 是音频 URL 或 base64/RAW 内容。
	Content string `json:"content"`
	// Callback 是回调地址；同步检测通常不需要传。
	Callback string `json:"callback"`
	// BtID 是客户侧音频文件唯一标识。
	BtID string `json:"btId"`
	// Data 是请求数据内容，常见字段包括 tokenId、returnAllText、lang 等。
	Data map[string]interface{} `json:"data"`
}

// AudioAsyncDetectReq 是音频异步检测请求的业务参数。
type AudioAsyncDetectReq struct {
	// Type 是音频检测风险类型，不传时 SDK 使用默认值。
	Type string `json:"type"`
	// BusinessType 是音频业务标签类型，和 Type 至少传一个。
	BusinessType string `json:"businessType"`
	// TranslationTargetLang 是翻译目标语种，可选 zh、en，需联系数美开通。
	TranslationTargetLang string `json:"translationTargetLang,omitempty"`
	// ContentType 是待检测音频内容格式，可选 URL、RAW。
	ContentType string `json:"contentType"`
	// Content 是音频 URL 或 base64/RAW 内容。
	Content string `json:"content"`
	// Callback 是回调地址。
	Callback string `json:"callback"`
	// BtID 是客户侧音频文件唯一标识。
	BtID string `json:"btId"`
	// AcceptLang 是返回标签语种，可选 zh、en。
	AcceptLang string `json:"acceptLang,omitempty"`
	// Data 是请求数据内容，常见字段包括 tokenId、returnAllText、lang 等。
	Data map[string]interface{} `json:"data"`
	// RetryURL 是 content 中音频 URL 下载失败时使用的重试地址。
	RetryURL string `json:"retryUrl,omitempty"`
}

// VideoAsyncDetectReq 是视频文件异步检测请求的业务参数。
type VideoAsyncDetectReq struct {
	// ImgType 是视频画面检测风险类型，不传时 SDK 使用默认值。
	ImgType string `json:"imgType"`
	// ImgBusinessType 是视频画面业务标签类型，和 ImgType 至少传一个。
	ImgBusinessType string `json:"imgBusinessType"`
	// AudioType 是视频音频检测风险类型，不传时 SDK 使用默认值。
	AudioType string `json:"audioType"`
	// AudioBusinessType 是视频音频业务标签类型，和 AudioType 至少传一个。
	AudioBusinessType string `json:"audioBusinessType"`
	// Callback 是视频检测结果回调地址。
	Callback string `json:"callback"`
	// AcceptLang 是返回标签语种，可选 zh、en。
	AcceptLang string `json:"acceptLang,omitempty"`
	// Data 是请求数据内容，常见字段包括 btId、url、tokenId、screenshotType 等。
	Data map[string]interface{} `json:"data"`
}

// SkyNetEventReq 数美天网请求
// See doc：https://help.ishumei.com/docs/tw/diversion/newest/developDoc#data
type SkyNetEventReq struct {
	// Data 是事件特有参数，常见字段包括 tokenId、ip、timestamp、deviceId 等。
	Data map[string]interface{} `json:"data"`
}
