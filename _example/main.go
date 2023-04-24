package main

import (
	"log"
	"time"

	"github.com/feymanlee/shumei-go"
)

const (
	appId = "CalfVR"
	Ac    = "J6Zx8A5UFj9IDkuwZmC5"
)

func main() {
	// Region 可选项：
	// RegionDefault 默认地区
	// RegionBeijing 北京地区
	// RegionShanghai 上海地区
	// RegionGuangzhou 广州地区
	// RegionVirginia 弗吉尼亚地区
	// RegionSingapore 新加坡地区
	// RegionSiliconValley 硅谷地区
	// RegionIndia 印度地区
	// client := shumei.NewClient("Your App ID", "Your Access Key", shumei.WithRegion(shumei.RegionDefault), shumei.WithTimeout(time.Second*5))
	client := shumei.NewClient(appId, Ac, shumei.WithRegion(shumei.RegionDefault), shumei.WithTimeout(time.Second*5))
	// 文本检测
	res, err := client.TextDetect("CalfVR_video_chat", shumei.TextDetectReq{
		Data: map[string]interface{}{
			"text":     "加个好友吧 qq12345",
			"tokenId":  "4567898765jhgfdsa",
			"ip":       "118.89.214.89",
			"deviceId": "*************",
			"nickname": "***********",
			"extra": map[string]interface{}{
				"topic":          "12345",
				"atId":           "username1",
				"room":           "ceshi123",
				"receiveTokenId": "username2",
				"role":           "USER",
			},
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
