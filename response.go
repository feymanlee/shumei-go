/**
 * @Author: lifameng@changba.com
 * @Description:
 * @File:  response
 * @Date: 2023/4/3 15:17
 */

package shumei

type CommonResp struct {
	Code      int    `json:"code"`
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
}

// TextDetectResp 文本检测响应
// See：
type TextDetectResp struct {
	CommonResp
	AllLabels []struct {
		Probability     interface{} `json:"probability"`
		RiskDescription string      `json:"riskDescription"`
		RiskDetail      struct {
		} `json:"riskDetail"`
		RiskLabel1 string `json:"riskLabel1"`
		RiskLabel2 string `json:"riskLabel2"`
		RiskLabel3 string `json:"riskLabel3"`
		RiskLevel  string `json:"riskLevel"`
	} `json:"allLabels"`
	AuxInfo struct {
		ContactResult []struct {
			ContactString string `json:"contactString"`
			ContactType   int    `json:"contactType"`
		} `json:"contactResult"`
		FilteredText string `json:"filteredText"`
	} `json:"auxInfo"`
	BusinessLabels  []interface{} `json:"businessLabels"`
	RiskDescription string        `json:"riskDescription"`
	RiskDetail      struct {
	} `json:"riskDetail"`
	RiskLabel1 string `json:"riskLabel1"`
	RiskLabel2 string `json:"riskLabel2"`
	RiskLabel3 string `json:"riskLabel3"`
	RiskLevel  string `json:"riskLevel"`
}

// ImageDetectResp 图片检测同步响应
// See：https://help.ishumei.com/docs/tj/image/newest/developDoc#%E5%90%8C%E6%AD%A5%E8%BF%94%E5%9B%9E%E7%A4%BA%E4%BE%8B%EF%BC%9A
type ImageDetectResp struct {
	RequestID       string `json:"requestId"`
	Code            int    `json:"code"`
	Message         string `json:"message"`
	RiskLevel       string `json:"riskLevel"`
	RiskLabel1      string `json:"riskLabel1"`
	RiskLabel2      string `json:"riskLabel2"`
	RiskLabel3      string `json:"riskLabel3"`
	RiskDescription string `json:"riskDescription"`
	RiskDetail      struct {
		Faces []struct {
			FaceRatio   float64 `json:"face_ratio"`
			ID          string  `json:"id"`
			Location    []int   `json:"location"`
			Name        string  `json:"name"`
			Probability float64 `json:"probability"`
		} `json:"faces"`
		RiskSource int `json:"riskSource"`
	} `json:"riskDetail"`
	AuxInfo struct {
		Segments    int `json:"segments"`
		TypeVersion struct {
			POLITICS string `json:"POLITICS"`
			VIOLENCE string `json:"VIOLENCE"`
			BAN      string `json:"BAN"`
			PORN     string `json:"PORN"`
		} `json:"typeVersion"`
	} `json:"auxInfo"`
	AllLabels []struct {
		RiskLabel1      string  `json:"riskLabel1"`
		RiskLabel2      string  `json:"riskLabel2"`
		RiskLabel3      string  `json:"riskLabel3"`
		RiskLevel       string  `json:"riskLevel"`
		Probability     float64 `json:"probability,omitempty"`
		RiskDescription string  `json:"riskDescription,omitempty"`
		RiskDetail      struct {
			Faces []struct {
				FaceRatio   float64 `json:"face_ratio"`
				ID          string  `json:"id"`
				Location    []int   `json:"location"`
				Name        string  `json:"name"`
				Probability float64 `json:"probability"`
			} `json:"faces"`
			RiskSocrce int `json:"riskSocrce"`
		} `json:"riskDetail,omitempty"`
	} `json:"allLabels"`
	BusinessLabels []struct {
		BusinessDescription string `json:"businessDescription"`
		BusinessDetail      struct {
		} `json:"businessDetail"`
		BusinessLabel1  string  `json:"businessLabel1"`
		BusinessLabel2  string  `json:"businessLabel2"`
		BusinessLabel3  string  `json:"businessLabel3"`
		ConfidenceLevel int     `json:"confidenceLevel"`
		Probability     float64 `json:"probability"`
	} `json:"businessLabels"`
	TokenLabels struct {
		UGCAccountRisk struct {
		} `json:"UGC_account_risk"`
	} `json:"tokenLabels"`
}

// SkyNetEventResp 天网响应
// See doc：https://help.ishumei.com/docs/tw/marketing/newest/developDoc#%E5%90%8C%E6%AD%A5%E8%BF%94%E5%9B%9E%E7%A4%BA%E4%BE%8B%EF%BC%9A
type SkyNetEventResp struct {
	CommonResp
	RiskLevel       string `json:"riskLevel"`
	TokenRiskLabels []struct {
		Timestamp   int64  `json:"timestamp"`
		Label1      string `json:"label1"`
		Label2      string `json:"label2"`
		Label3      string `json:"label3"`
		Description string `json:"description"`
	} `json:"tokenRiskLabels"`
	Detail struct {
		IPProvince string `json:"ip_province"`
		Hits       []struct {
			Description string `json:"description"`
			Model       string `json:"model"`
			RiskLevel   string `json:"riskLevel"`
		} `json:"hits"`
		Model              string `json:"model"`
		MachineAccountRisk struct {
			TokenSampleDesc   string `json:"tokenSampleDesc"`
			TokenSampleLastTs int64  `json:"tokenSampleLastTs"`
		} `json:"machineAccountRisk"`
		Description string `json:"description"`
		IPCountry   string `json:"ip_country"`
		IPCity      string `json:"ip_city"`
	} `json:"detail"`
}
