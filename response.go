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

type TextDetectResp struct {
	CommonResp
	AllLabels []struct {
		Probability     int    `json:"probability"`
		RiskDescription string `json:"riskDescription"`
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
