/**
 * @Author: lifameng@changba.com
 * @Description:
 * @File:  response
 * @Date: 2023/4/3 15:17
 */

package shumei

// CommonResp 是数美接口通用响应字段。
type CommonResp struct {
	// Code 是数美业务返回码，1100 表示成功。
	Code int `json:"code"`
	// RequestID 是数美请求唯一标识。
	RequestID string `json:"requestId"`
	// Message 是与 Code 对应的返回描述。
	Message string `json:"message"`
}

// TextDetectResp 文本检测响应
// See：https://help.ishumei.com/docs/tj/text/versionV4/sync/developDoc
type TextDetectResp struct {
	CommonResp
	// AllLabels 是全部风险标签列表。
	AllLabels []struct {
		// Probability 是置信度。
		Probability interface{} `json:"probability"`
		// RiskDescription 是风险原因描述。
		RiskDescription string `json:"riskDescription"`
		// RiskDetail 是风险详情。
		RiskDetail struct {
		} `json:"riskDetail"`
		// RiskLabel1 是一级风险标签。
		RiskLabel1 string `json:"riskLabel1"`
		// RiskLabel2 是二级风险标签。
		RiskLabel2 string `json:"riskLabel2"`
		// RiskLabel3 是三级风险标签。
		RiskLabel3 string `json:"riskLabel3"`
		// RiskLevel 是处置建议。
		RiskLevel string `json:"riskLevel"`
	} `json:"allLabels"`
	// AuxInfo 是文本检测辅助信息。
	AuxInfo struct {
		// ContactResult 是联系方式识别结果。
		ContactResult []struct {
			// ContactString 是命中的联系方式文本。
			ContactString string `json:"contactString"`
			// ContactType 是联系方式类型。
			ContactType int `json:"contactType"`
		} `json:"contactResult"`
		// FilteredText 是命中敏感内容后的过滤文本。
		FilteredText string `json:"filteredText"`
	} `json:"auxInfo"`
	// BusinessLabels 是业务标签列表。
	BusinessLabels []interface{} `json:"businessLabels"`
	// TokenProfileLabels 是账号属性标签，开启对应服务时返回。
	TokenProfileLabels []struct {
		// Label1 是一级账号属性标签。
		Label1 string `json:"label1"`
		// Label2 是二级账号属性标签。
		Label2 string `json:"label2"`
		// Label3 是三级账号属性标签。
		Label3 string `json:"label3"`
		// Description 是账号属性标签描述。
		Description string `json:"description"`
		// Timestamp 是打标签时间戳，单位毫秒。
		Timestamp int64 `json:"timestamp"`
	} `json:"tokenProfileLabels"`
	// TokenRiskLabels 是账号风险标签，开启对应服务时返回。
	TokenRiskLabels []struct {
		// Label1 是一级账号风险标签。
		Label1 string `json:"label1"`
		// Label2 是二级账号风险标签。
		Label2 string `json:"label2"`
		// Label3 是三级账号风险标签。
		Label3 string `json:"label3"`
		// Description 是账号风险标签描述。
		Description string `json:"description"`
		// Timestamp 是打标签时间戳，单位毫秒。
		Timestamp int64 `json:"timestamp"`
	} `json:"tokenRiskLabels"`
	// LangResult 是语种检测和翻译结果。
	LangResult struct {
		// DetectedLang 是自动语种识别结果。
		DetectedLang string `json:"detectedLang"`
		// TranslatedText 是传入 translationTargetLang 时返回的翻译文本。
		TranslatedText string `json:"translatedText"`
	} `json:"langResult"`
	// FinalResult 表示是否为最终审核结果，0 非最终、1 最终。
	FinalResult int `json:"finalResult"`
	// ResultType 表示结果来源，0 机审、1 人审。
	ResultType int `json:"resultType"`
	// RiskDescription 是主风险原因描述。
	RiskDescription string `json:"riskDescription"`
	// RiskDetail 是主风险详情。
	RiskDetail struct {
	} `json:"riskDetail"`
	// RiskLabel1 是主一级风险标签。
	RiskLabel1 string `json:"riskLabel1"`
	// RiskLabel2 是主二级风险标签。
	RiskLabel2 string `json:"riskLabel2"`
	// RiskLabel3 是主三级风险标签。
	RiskLabel3 string `json:"riskLabel3"`
	// RiskLevel 是主处置建议。
	RiskLevel string `json:"riskLevel"`
}

// ImageDetectResp 图片检测同步响应
// See：https://help.ishumei.com/docs/tj/image/versionV4/syncSingle/developDoc
type ImageDetectResp struct {
	// RequestID 是数美请求唯一标识。
	RequestID string `json:"requestId"`
	// Code 是数美业务返回码，1100 表示成功。
	Code int `json:"code"`
	// Message 是与 Code 对应的返回描述。
	Message string `json:"message"`
	// RiskLevel 是主处置建议。
	RiskLevel string `json:"riskLevel"`
	// RiskLabel1 是主一级风险标签。
	RiskLabel1 string `json:"riskLabel1"`
	// RiskLabel2 是主二级风险标签。
	RiskLabel2 string `json:"riskLabel2"`
	// RiskLabel3 是主三级风险标签。
	RiskLabel3 string `json:"riskLabel3"`
	// RiskDescription 是主风险原因描述。
	RiskDescription string `json:"riskDescription"`
	// ResultType 表示结果来源，0 机审、1 人审。
	ResultType int `json:"resultType"`
	// FinalResult 表示是否为最终审核结果，0 非最终、1 最终。
	FinalResult int `json:"finalResult"`
	// RiskDetail 是主风险详情。
	RiskDetail struct {
		// Faces 是涉政人物等人脸位置信息。
		Faces []struct {
			// FaceRatio 是人脸占比。
			FaceRatio float64 `json:"face_ratio"`
			// ID 是同一位置人物在不同标签下的编号。
			ID string `json:"id"`
			// Location 是左上角和右下角坐标。
			Location []int `json:"location"`
			// Name 是人物名称。
			Name string `json:"name"`
			// Probability 是置信度。
			Probability float64 `json:"probability"`
		} `json:"faces"`
		// RiskSource 是风险来源，1001 文本风险、1002 视觉图片风险。
		RiskSource int `json:"riskSource"`
	} `json:"riskDetail"`
	// AuxInfo 是图片检测辅助信息。
	AuxInfo struct {
		// Segments 是实际处理的片段数量。
		Segments int `json:"segments"`
		// TypeVersion 是各风险类型的模型版本。
		TypeVersion struct {
			POLITICS string `json:"POLITICS"`
			VIOLENCE string `json:"VIOLENCE"`
			BAN      string `json:"BAN"`
			PORN     string `json:"PORN"`
		} `json:"typeVersion"`
	} `json:"auxInfo"`
	// AllLabels 是全部风险标签列表。
	AllLabels []struct {
		// RiskLabel1 是一级风险标签。
		RiskLabel1 string `json:"riskLabel1"`
		// RiskLabel2 是二级风险标签。
		RiskLabel2 string `json:"riskLabel2"`
		// RiskLabel3 是三级风险标签。
		RiskLabel3 string `json:"riskLabel3"`
		// RiskLevel 是处置建议。
		RiskLevel string `json:"riskLevel"`
		// Probability 是置信度。
		Probability float64 `json:"probability,omitempty"`
		// RiskDescription 是风险原因描述。
		RiskDescription string `json:"riskDescription,omitempty"`
		// RiskDetail 是风险详情。
		RiskDetail struct {
			// Faces 是涉政人物等人脸位置信息。
			Faces []struct {
				FaceRatio   float64 `json:"face_ratio"`
				ID          string  `json:"id"`
				Location    []int   `json:"location"`
				Name        string  `json:"name"`
				Probability float64 `json:"probability"`
			} `json:"faces"`
			// RiskSource 是风险来源，1001 文本风险、1002 视觉图片风险。
			RiskSource int `json:"riskSource"`
		} `json:"riskDetail,omitempty"`
	} `json:"allLabels"`
	// BusinessLabels 是业务标签列表。
	BusinessLabels []struct {
		// BusinessDescription 是业务标签中文描述。
		BusinessDescription string `json:"businessDescription"`
		// BusinessDetail 是业务标签详情。
		BusinessDetail struct {
		} `json:"businessDetail"`
		// BusinessLabel1 是一级业务标签。
		BusinessLabel1 string `json:"businessLabel1"`
		// BusinessLabel2 是二级业务标签。
		BusinessLabel2 string `json:"businessLabel2"`
		// BusinessLabel3 是三级业务标签。
		BusinessLabel3 string `json:"businessLabel3"`
		// ConfidenceLevel 是置信等级。
		ConfidenceLevel int `json:"confidenceLevel"`
		// Probability 是置信度。
		Probability float64 `json:"probability"`
	} `json:"businessLabels"`
	// TokenLabels 是账号相关标签，开启对应服务时返回。
	TokenLabels struct {
		UGCAccountRisk struct {
		} `json:"UGC_account_risk"`
	} `json:"tokenLabels"`
}

// AudioSyncDetectResp 是音频同步检测响应。
// See：https://help.ishumei.com/docs/tj/audio/versionV4/sync/developDoc
type AudioSyncDetectResp struct {
	// Code 是数美业务返回码，1100 表示成功。
	Code int `json:"code"`
	// Message 是与 Code 对应的返回描述。
	Message string `json:"message"`
	// RequestID 是数美请求唯一标识。
	RequestID string `json:"requestId"`
	// BtID 是客户侧音频文件唯一标识。
	BtID string `json:"btId"`
	// Detail 是 code 为 1100 时返回的检测明细。
	Detail struct {
		// AudioDetail 是音频片段检测结果。
		AudioDetail []struct {
			// RequestID 是音频片段请求唯一标识。
			RequestID string `json:"requestId"`
			// AudioStarttime 是片段起始时间，单位秒。
			AudioStarttime float64 `json:"audioStarttime"`
			// AudioEndtime 是片段结束时间，单位秒。
			AudioEndtime float64 `json:"audioEndtime"`
			// AudioURL 是音频片段链接。
			AudioURL string `json:"audioUrl"`
			// RiskLevel 是片段处置建议。
			RiskLevel string `json:"riskLevel"`
			// RiskLabel1 是片段一级风险标签。
			RiskLabel1 string `json:"riskLabel1"`
			// RiskLabel2 是片段二级风险标签。
			RiskLabel2 string `json:"riskLabel2"`
			// RiskLabel3 是片段三级风险标签。
			RiskLabel3 string `json:"riskLabel3"`
			// RiskDescription 是片段风险原因描述。
			RiskDescription string `json:"riskDescription"`
			// RiskDetail 是片段风险详情。
			RiskDetail struct {
				// AudioText 是该片段识别出的文字内容。
				AudioText string `json:"audioText"`
			} `json:"riskDetail"`
		} `json:"audioDetail"`
		// AudioTags 是音频标签，历史兼容字段，建议优先使用业务标签。
		AudioTags struct {
		} `json:"audioTags"`
		// AudioText 是整段音频转译文本。
		AudioText string `json:"audioText"`
		// AudioTime 是整段音频时长，单位秒。
		AudioTime int `json:"audioTime"`
		// Code 是明细状态码。
		Code int `json:"code"`
		// RequestParams 是 data 下请求参数的透传结果。
		RequestParams struct {
			Channel       string `json:"channel"`
			Lang          string `json:"lang"`
			ReturnAllText int    `json:"returnAllText"`
			TokenID       string `json:"tokenId"`
		} `json:"requestParams"`
		// RiskLevel 是整段音频处置建议。
		RiskLevel string `json:"riskLevel"`
	} `json:"detail"`
}

// AudioAsyncDetectResp 是音频异步检测提交响应。
// See：https://help.ishumei.com/docs/tj/audio/versionV4/async/developDoc
type AudioAsyncDetectResp struct {
	// Code 是数美业务返回码，1100 表示成功。
	Code int `json:"code"`
	// Message 是与 Code 对应的返回描述。
	Message string `json:"message"`
	// RequestID 是数美请求唯一标识。
	RequestID string `json:"requestId"`
	// BtID 是客户侧音频文件唯一标识，code 为 1100 时返回。
	BtID string `json:"btId"`
}

// VideoAsyncDetectResp 是视频检测提交响应。
// See：https://help.ishumei.com/docs/tj/video/versionV4/requestInterface/developDoc
type VideoAsyncDetectResp struct {
	// Code 是数美业务返回码，1100 表示成功。
	Code int `json:"code"`
	// Message 是与 Code 对应的返回描述。
	Message string `json:"message"`
	// RequestID 是数美请求唯一标识。
	RequestID string `json:"requestId"`
	// BtID 是客户侧视频请求唯一标识。
	BtID string `json:"btId"`
}

// SkyNetEventResp 天网响应
// See doc：https://help.ishumei.com/docs/tw/diversion/newest/developDoc#%E5%90%8C%E6%AD%A5%E8%BF%94%E5%9B%9E%E7%A4%BA%E4%BE%8B%EF%BC%9A
type SkyNetEventResp struct {
	CommonResp
	// RiskLevel 是天网事件处置建议。
	RiskLevel string `json:"riskLevel"`
	// TokenRiskLabels 是账号风险标签，开启对应服务时返回。
	TokenRiskLabels []struct {
		// Timestamp 是打标签时间戳，单位毫秒。
		Timestamp int64 `json:"timestamp"`
		// Label1 是一级标签。
		Label1 string `json:"label1"`
		// Label2 是二级标签。
		Label2 string `json:"label2"`
		// Label3 是三级标签。
		Label3 string `json:"label3"`
		// Description 是标签描述。
		Description string `json:"description"`
	} `json:"tokenRiskLabels"`
	// Detail 是天网事件命中详情。
	Detail struct {
		// IPProvince 是 IP 所在省份。
		IPProvince string `json:"ip_province"`
		// Hits 是命中的风险模型列表。
		Hits []struct {
			Description string `json:"description"`
			Model       string `json:"model"`
			RiskLevel   string `json:"riskLevel"`
		} `json:"hits"`
		// Model 是主命中模型。
		Model string `json:"model"`
		// MachineAccountRisk 是机器账号风险相关详情。
		MachineAccountRisk struct {
			TokenSampleDesc   string `json:"tokenSampleDesc"`
			TokenSampleLastTs int64  `json:"tokenSampleLastTs"`
		} `json:"machineAccountRisk"`
		// Description 是详情描述。
		Description string `json:"description"`
		// IPCountry 是 IP 所在国家。
		IPCountry string `json:"ip_country"`
		// IPCity 是 IP 所在城市。
		IPCity string `json:"ip_city"`
	} `json:"detail"`
}
