# shumei-go

数美 Go SDK，封装文本、图片、音频、视频和天网事件检测接口。

## 安装

```bash
go get github.com/feymanlee/shumei-go
```

## 快速开始

建议通过环境变量传入凭据，避免把 `appID` 和 `accessKey` 写进代码或提交到仓库。

```bash
export SHUMEI_APP_ID="your-app-id"
export SHUMEI_ACCESS_KEY="your-access-key"
```

```go
package main

import (
	"log"
	"os"
	"time"

	shumei "github.com/feymanlee/shumei-go"
)

func main() {
	appID := os.Getenv("SHUMEI_APP_ID")
	accessKey := os.Getenv("SHUMEI_ACCESS_KEY")
	if appID == "" || accessKey == "" {
		log.Fatal("SHUMEI_APP_ID and SHUMEI_ACCESS_KEY are required")
	}

	client, err := shumei.NewClient(
		appID,
		accessKey,
		shumei.WithRegion(shumei.RegionDefault),
		shumei.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.TextDetect("your-event-id", shumei.TextDetectReq{
		Data: map[string]interface{}{
			"text":    "加个好友吧 qq12345",
			"tokenId": "user-123",
			"ip":      "127.0.0.1",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("riskLevel=%s requestID=%s", resp.RiskLevel, resp.RequestID)
}
```

## 客户端配置

```go
client, err := shumei.NewClient(
	appID,
	accessKey,
	shumei.WithRegion(shumei.RegionBeijing),
	shumei.WithTimeout(3*time.Second),
	shumei.WithDefaultRegionFallback(true),
)
```

- `NewClient(appID, accessKey, ...options)` 返回 `(*Client, error)`。
- `WithRegion(region)` 指定服务集群；默认是 `RegionDefault`。
- `WithTimeout(timeout)` 指定 HTTP 请求超时时间；默认是 `3s`。
- `WithDefaultRegionFallback(true)` 开启默认区域补位：当指定 Region 不支持某个能力时，自动使用该能力的 `RegionDefault` endpoint。
- SDK 使用 HTTPS endpoint 访问数美服务。

## Region

不同能力在数美侧支持的集群不同。选择了某个能力不支持的 Region 时，请求前会返回 endpoint not found 错误；如果开启了 `WithDefaultRegionFallback(true)`，则会回退使用该能力的 `RegionDefault` endpoint。

| 常量 | 含义 |
| --- | --- |
| `RegionDefault` | 默认集群 |
| `RegionBeijing` | 北京 |
| `RegionShanghai` | 上海 |
| `RegionGuangzhou` | 广州 |
| `RegionVirginia` | 美国弗吉尼亚 |
| `RegionSingapore` | 新加坡 |
| `RegionSiliconValley` | 美国硅谷 |
| `RegionIndia` | 印度 |
| `RegionFrankfurt` | 欧洲法兰克福 |

### 能力支持的 Region

| 能力 | 支持的 Region |
| --- | --- |
| 文本检测 | `RegionDefault`, `RegionBeijing`, `RegionShanghai`, `RegionGuangzhou`, `RegionVirginia`, `RegionSingapore` |
| 图片检测 | `RegionDefault`, `RegionBeijing`, `RegionShanghai`, `RegionSiliconValley`, `RegionSingapore` |
| 图片主动查询 | `RegionDefault`, `RegionBeijing` |
| 音频同步检测 | `RegionDefault`, `RegionShanghai`, `RegionSingapore` |
| 音频异步检测 | `RegionDefault`, `RegionShanghai`, `RegionSiliconValley`, `RegionSingapore` |
| 音频主动查询 | `RegionDefault`, `RegionShanghai`, `RegionSiliconValley`, `RegionSingapore` |
| 视频异步检测 | `RegionDefault`, `RegionBeijing`, `RegionShanghai`, `RegionSingapore` |
| 视频主动查询 | `RegionDefault`, `RegionBeijing`, `RegionShanghai`, `RegionSingapore` |
| 天网事件 | `RegionDefault`, `RegionBeijing`, `RegionVirginia`, `RegionSingapore`, `RegionFrankfurt` |

## 接口示例

### 文本检测

```go
resp, err := client.TextDetect("your-event-id", shumei.TextDetectReq{
	Data: map[string]interface{}{
		"text":    "hello",
		"tokenId": "user-123",
	},
})
```

### 图片检测

```go
resp, err := client.ImageDetect("your-event-id", shumei.ImageDetectReq{
	Data: map[string]interface{}{
		"img":     "https://example.com/image.jpg",
		"tokenId": "user-123",
	},
})
```

### 音频同步检测

```go
resp, err := client.AudioSyncDetect("your-event-id", shumei.AudioSyncDetectReq{
	ContentType: "URL",
	Content:     "https://example.com/audio.mp3",
	Data: map[string]interface{}{
		"tokenId": "user-123",
	},
})
```

### 音频异步检测

```go
resp, err := client.AudioAsyncDetect("your-event-id", shumei.AudioAsyncDetectReq{
	ContentType: "URL",
	Content:     "https://example.com/audio.mp3",
	Callback:   "https://example.com/shumei/audio-callback",
	Data: map[string]interface{}{
		"tokenId": "user-123",
	},
})
```

### 视频异步检测

```go
resp, err := client.VideoAsyncDetect("your-event-id", shumei.VideoAsyncDetectReq{
	Callback: "https://example.com/shumei/video-callback",
	Data: map[string]interface{}{
		"btId":    "video-123",
		"url":     "https://example.com/video.mp4",
		"tokenId": "user-123",
	},
})
```

### 天网事件

```go
resp, err := client.SkyNetEvent("your-event-id", shumei.SkyNetEventReq{
	Data: map[string]interface{}{
		"tokenId": "user-123",
		"ip":      "127.0.0.1",
	},
})
```

## 错误处理

所有检测方法都会返回 `(*Resp, error)`：

- HTTP 非 2xx 响应会返回包含 `http status`、`code`、`message`、`requestID` 的错误。
- 数美业务返回码不是 `1100` 时，会返回包含业务 `code`、`message`、`requestID` 的错误。
- JSON 解析失败会直接返回解析错误。

```go
resp, err := client.TextDetect("your-event-id", req)
if err != nil {
	log.Println("shumei request failed:", err)
	return
}
log.Println(resp.RequestID)
```

## 注意事项

- `eventID` 需要先在数美后台配置，否则接口会返回类似 `eventId value not supported` 的业务错误。
- `Type` 字段不传时，SDK 会为文本、图片、音频、视频填充默认检测类型。
- `Data` 使用 `map[string]interface{}`，字段名和内容请以数美对应产品文档为准。
- 不要把真实 `accessKey` 提交到代码仓库；若已经泄露，应及时在数美后台轮换密钥。

## 开发

```bash
go test ./...
go vet ./...
```
