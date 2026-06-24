package shumei

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestEndpointsUseHTTPS(t *testing.T) {
	client, err := NewClient("app", "key")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	for action, endpoints := range actionRegionEndpoints {
		for region := range endpoints {
			client.SetRegion(region)
			endpoint, err := client.getEndpoint(action)
			if err != nil {
				t.Fatalf("getEndpoint(%s, %s) returned error: %v", action, region, err)
			}
			if !strings.HasPrefix(endpoint, "https://") {
				t.Fatalf("endpoint for action %s region %s must use https, got %s", action, region, endpoint)
			}
		}
	}
}

func TestEventRegionEndpoints(t *testing.T) {
	tests := map[string]string{
		RegionDefault:   "https://api-skynet-bj.fengkongcloud.com/v4/event",
		RegionBeijing:   "https://api-skynet-bj.fengkongcloud.com/v4/event",
		RegionVirginia:  "https://api-skynet-fjny.fengkongcloud.com/v4/event",
		RegionSingapore: "https://api-skynet-xjp.fengkongcloud.com/v4/event",
		RegionFrankfurt: "https://api-skynet-eur.fengkongcloud.com/v4/event",
	}

	client, err := NewClient("app", "key")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	for region, want := range tests {
		client.SetRegion(region)
		got, err := client.getEndpoint(actionTypeEvent)
		if err != nil {
			t.Fatalf("getEndpoint(event, %s) returned error: %v", region, err)
		}
		if got != want {
			t.Fatalf("getEndpoint(event, %s) = %s, want %s", region, got, want)
		}
	}
}

func TestEventUnsupportedRegionReturnsError(t *testing.T) {
	client, err := NewClient("app", "key", WithRegion(RegionShanghai))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	_, err = client.getEndpoint(actionTypeEvent)
	if err == nil {
		t.Fatal("getEndpoint(event, shanghai) returned nil error")
	}
}

func TestDefaultRegionFallbackUsesDefaultEndpoint(t *testing.T) {
	client, err := NewClient("app", "key", WithRegion(RegionShanghai), WithDefaultRegionFallback(true))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	got, err := client.getEndpoint(actionTypeEvent)
	if err != nil {
		t.Fatalf("getEndpoint(event, shanghai) returned error: %v", err)
	}

	want := "https://api-skynet-bj.fengkongcloud.com/v4/event"
	if got != want {
		t.Fatalf("getEndpoint(event, shanghai) = %s, want default endpoint %s", got, want)
	}
}

func TestRegionGatewayMapMatchesLatestDocs(t *testing.T) {
	want := map[string]map[string]string{
		actionTypeText: {
			RegionDefault:   "https://api-text-bj.fengkongcloud.com/text/v4",
			RegionBeijing:   "https://api-text-bj.fengkongcloud.com/text/v4",
			RegionShanghai:  "https://api-text-sh.fengkongcloud.com/text/v4",
			RegionGuangzhou: "https://api-text-gz.fengkongcloud.com/text/v4",
			RegionVirginia:  "https://api-text-fjny.fengkongcloud.com/text/v4",
			RegionSingapore: "https://api-text-xjp.fengkongcloud.com/text/v4",
		},
		actionTypeImage: {
			RegionDefault:       "https://api-img-bj.fengkongcloud.com/image/v4",
			RegionBeijing:       "https://api-img-bj.fengkongcloud.com/image/v4",
			RegionShanghai:      "https://api-img-sh.fengkongcloud.com/image/v4",
			RegionSiliconValley: "https://api-img-gg.fengkongcloud.com/image/v4",
			RegionSingapore:     "https://api-img-xjp.fengkongcloud.com/image/v4",
		},
		actionTypeImageQuery: {
			RegionDefault: "https://api-img-active-query.fengkongcloud.com/v4/image/query",
			RegionBeijing: "https://api-img-active-query.fengkongcloud.com/v4/image/query",
		},
		actionTypeAudioSync: {
			RegionDefault:   "https://api-audio-sh.fengkongcloud.com/audiomessage/v4",
			RegionShanghai:  "https://api-audio-sh.fengkongcloud.com/audiomessage/v4",
			RegionSingapore: "https://api-audio-xjp.fengkongcloud.com/audiomessage/v4",
		},
		actionTypeAudioAsync: {
			RegionDefault:       "https://api-audio-sh.fengkongcloud.com/audio/v4",
			RegionShanghai:      "https://api-audio-sh.fengkongcloud.com/audio/v4",
			RegionSiliconValley: "https://api-audio-gg.fengkongcloud.com/audio/v4",
			RegionSingapore:     "https://api-audio-xjp.fengkongcloud.com/audio/v4",
		},
		actionTypeAudioQuery: {
			RegionDefault:       "https://api-audio-sh.fengkongcloud.com/query_audio/v4",
			RegionShanghai:      "https://api-audio-sh.fengkongcloud.com/query_audio/v4",
			RegionSiliconValley: "https://api-audio-gg.fengkongcloud.com/query_audio/v4",
			RegionSingapore:     "https://api-audio-xjp.fengkongcloud.com/query_audio/v4",
		},
		actionTypeVideoAsync: {
			RegionDefault:   "https://api-video-bj.fengkongcloud.com/video/v4",
			RegionBeijing:   "https://api-video-bj.fengkongcloud.com/video/v4",
			RegionShanghai:  "https://api-video-sh.fengkongcloud.com/video/v4",
			RegionSingapore: "https://api-video-xjp.fengkongcloud.com/video/v4",
		},
		actionTypeVideoQuery: {
			RegionDefault:   "https://api-video-bj.fengkongcloud.com/video/query/v4",
			RegionBeijing:   "https://api-video-bj.fengkongcloud.com/video/query/v4",
			RegionShanghai:  "https://api-video-sh.fengkongcloud.com/video/query/v4",
			RegionSingapore: "https://api-video-xjp.fengkongcloud.com/video/query/v4",
		},
		actionTypeEvent: {
			RegionDefault:   "https://api-skynet-bj.fengkongcloud.com/v4/event",
			RegionBeijing:   "https://api-skynet-bj.fengkongcloud.com/v4/event",
			RegionVirginia:  "https://api-skynet-fjny.fengkongcloud.com/v4/event",
			RegionSingapore: "https://api-skynet-xjp.fengkongcloud.com/v4/event",
			RegionFrankfurt: "https://api-skynet-eur.fengkongcloud.com/v4/event",
		},
	}

	if len(actionRegionEndpoints) != len(want) {
		t.Fatalf("actionRegionEndpoints has %d actions, want %d", len(actionRegionEndpoints), len(want))
	}

	for action, wantEndpoints := range want {
		gotEndpoints, ok := actionRegionEndpoints[action]
		if !ok {
			t.Fatalf("actionRegionEndpoints missing action %s", action)
		}
		if len(gotEndpoints) != len(wantEndpoints) {
			t.Fatalf("action %s has %d endpoints, want %d: %#v", action, len(gotEndpoints), len(wantEndpoints), gotEndpoints)
		}
		for region, wantEndpoint := range wantEndpoints {
			if got := gotEndpoints[region]; got != wantEndpoint {
				t.Fatalf("endpoint for action %s region %s = %s, want %s", action, region, got, wantEndpoint)
			}
		}
	}
}

func TestNewClientReturnsOptionError(t *testing.T) {
	wantErr := errors.New("invalid option")

	client, err := NewClient("app", "key", func(*Client) error {
		return wantErr
	})

	if client != nil {
		t.Fatalf("client = %#v, want nil", client)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
}

func TestRequestReturnsHTTPStatusAndBody(t *testing.T) {
	oldPoster := postJSON
	defer func() { postJSON = oldPoster }()

	postJSON = func(_ string, _ *resty.Request) (*resty.Response, error) {
		return restyResponse(http.StatusUnauthorized, []byte(`{"code":1901,"message":"bad access key","requestId":"req-123"}`)), nil
	}

	client, err := NewClient("app", "key")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	_, err = client.request(actionTypeText, TextDetectReq{})
	if err == nil {
		t.Fatal("request returned nil error")
	}

	msg := err.Error()
	for _, want := range []string{"http status[401]", "bad access key", "requestID[req-123]"} {
		if !strings.Contains(msg, want) {
			t.Fatalf("error %q does not contain %q", msg, want)
		}
	}
}

func TestRequestPayloadIncludesDefaultTypeAndCommonFields(t *testing.T) {
	oldPoster := postJSON
	defer func() { postJSON = oldPoster }()

	var body interface{}
	postJSON = func(_ string, req *resty.Request) (*resty.Response, error) {
		body = req.Body
		return restyResponse(http.StatusOK, []byte(`{"code":1100,"message":"ok","requestId":"req-123"}`)), nil
	}

	client, err := NewClient("app", "key")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	_, err = client.TextDetect("event", TextDetectReq{
		Data: map[string]interface{}{
			"text": "hello",
		},
	})
	if err != nil {
		t.Fatalf("TextDetect returned error: %v", err)
	}

	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal payload: %v", err)
	}

	if payload["type"] != defaultTextDetectType {
		t.Fatalf("type = %v, want %s", payload["type"], defaultTextDetectType)
	}
	if payload["accessKey"] != "key" || payload["appId"] != "app" || payload["eventId"] != "event" {
		t.Fatalf("common fields not populated: %#v", payload)
	}
}

func restyResponse(statusCode int, body []byte) *resty.Response {
	return (&resty.Response{
		RawResponse: &http.Response{
			StatusCode: statusCode,
			Status:     http.StatusText(statusCode),
		},
	}).SetBody(body)
}

func TestSkyNetEventReqJSONTag(t *testing.T) {
	raw, err := json.Marshal(SkyNetEventReq{
		Data: map[string]interface{}{"tokenId": "token"},
	})
	if err != nil {
		t.Fatalf("marshal SkyNetEventReq: %v", err)
	}

	if !strings.Contains(string(raw), `"data"`) {
		t.Fatalf("SkyNetEventReq JSON = %s, want data field", raw)
	}
	if strings.Contains(string(raw), `"Data"`) {
		t.Fatalf("SkyNetEventReq JSON = %s, must not contain Data field", raw)
	}
}

func TestImageRiskSourceUnmarshalsFromCorrectField(t *testing.T) {
	raw := []byte(`{
		"code": 1100,
		"allLabels": [{
			"riskDetail": {
				"riskSource": 7
			}
		}]
	}`)

	var resp ImageDetectResp
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal ImageDetectResp: %v", err)
	}

	if got := resp.AllLabels[0].RiskDetail.RiskSource; got != 7 {
		t.Fatalf("RiskSource = %d, want 7", got)
	}
}

func TestBtIDUnmarshalsFromBtIDField(t *testing.T) {
	raw := []byte(`{"code":1100,"requestId":"req-123","btId":"bt-456"}`)

	var audioResp AudioSyncDetectResp
	if err := json.Unmarshal(raw, &audioResp); err != nil {
		t.Fatalf("unmarshal AudioSyncDetectResp: %v", err)
	}
	if audioResp.BtID != "bt-456" {
		t.Fatalf("AudioSyncDetectResp.BtID = %s, want bt-456", audioResp.BtID)
	}

	var videoResp VideoAsyncDetectResp
	if err := json.Unmarshal(raw, &videoResp); err != nil {
		t.Fatalf("unmarshal VideoAsyncDetectResp: %v", err)
	}
	if videoResp.BtID != "bt-456" {
		t.Fatalf("VideoAsyncDetectResp.BtID = %s, want bt-456", videoResp.BtID)
	}
}

func TestRequestStructsMarshalLatestDocumentedTopLevelFields(t *testing.T) {
	tests := []struct {
		name string
		req  interface{}
		want map[string]interface{}
	}{
		{
			name: "text agent type translation accept lang callback",
			req: TextDetectReq{
				Type:                  "TEXTRISK",
				AgentType:             "POLITY",
				TranslationTargetLang: "en",
				AcceptLang:            "zh",
				Callback:              "https://example.com/callback",
				Data:                  map[string]interface{}{"text": "hello"},
			},
			want: map[string]interface{}{
				"type":                  "TEXTRISK",
				"agentType":             "POLITY",
				"translationTargetLang": "en",
				"acceptLang":            "zh",
				"callback":              "https://example.com/callback",
				"data":                  map[string]interface{}{"text": "hello"},
			},
		},
		{
			name: "image agent type and accept lang",
			req: ImageDetectReq{
				Type:       "POLITY",
				AgentType:  "QRCODE",
				AcceptLang: "en",
				Data:       map[string]interface{}{"img": "https://example.com/image.jpg"},
			},
			want: map[string]interface{}{
				"type":       "POLITY",
				"agentType":  "QRCODE",
				"acceptLang": "en",
				"data":       map[string]interface{}{"img": "https://example.com/image.jpg"},
			},
		},
		{
			name: "audio async translation accept lang retry url",
			req: AudioAsyncDetectReq{
				Type:                  "POLITY",
				TranslationTargetLang: "en",
				ContentType:           "URL",
				Content:               "https://example.com/audio.mp3",
				BtID:                  "audio-1",
				AcceptLang:            "zh",
				RetryURL:              "https://example.com/retry.mp3",
				Data:                  map[string]interface{}{"tokenId": "token-1"},
			},
			want: map[string]interface{}{
				"type":                  "POLITY",
				"translationTargetLang": "en",
				"contentType":           "URL",
				"content":               "https://example.com/audio.mp3",
				"btId":                  "audio-1",
				"acceptLang":            "zh",
				"retryUrl":              "https://example.com/retry.mp3",
				"data":                  map[string]interface{}{"tokenId": "token-1"},
			},
		},
		{
			name: "video data bt id and accept lang",
			req: VideoAsyncDetectReq{
				ImgType:    "POLITY",
				AudioType:  "POLITY",
				Callback:   "https://example.com/callback",
				AcceptLang: "en",
				Data: map[string]interface{}{
					"btId": "video-1",
					"url":  "https://example.com/video.mp4",
				},
			},
			want: map[string]interface{}{
				"imgType":    "POLITY",
				"audioType":  "POLITY",
				"callback":   "https://example.com/callback",
				"acceptLang": "en",
				"data": map[string]interface{}{
					"btId": "video-1",
					"url":  "https://example.com/video.mp4",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw, err := json.Marshal(tt.req)
			if err != nil {
				t.Fatalf("marshal request: %v", err)
			}

			got := map[string]interface{}{}
			if err := json.Unmarshal(raw, &got); err != nil {
				t.Fatalf("unmarshal request JSON: %v", err)
			}
			if tt.name == "video data bt id and accept lang" {
				if _, ok := got["btId"]; ok {
					t.Fatalf("video request must not marshal top-level btId: %s", raw)
				}
			}

			for key, want := range tt.want {
				gotRaw, ok := got[key]
				if !ok {
					t.Fatalf("JSON %s missing key %s: %s", tt.name, key, raw)
				}
				if !jsonEqual(gotRaw, want) {
					t.Fatalf("JSON %s key %s = %#v, want %#v; raw=%s", tt.name, key, gotRaw, want, raw)
				}
			}
		})
	}
}

func TestVideoAsyncDetectReqDoesNotExposeTopLevelBtID(t *testing.T) {
	if _, ok := reflect.TypeOf(VideoAsyncDetectReq{}).FieldByName("BtID"); ok {
		t.Fatal("VideoAsyncDetectReq must not expose top-level BtID; video btId belongs in Data")
	}
}

func TestAudioSyncDetailTimesUnmarshalAsFloat(t *testing.T) {
	raw := []byte(`{
		"code": 1100,
		"detail": {
			"audioDetail": [{
				"audioStarttime": 1.5,
				"audioEndtime": 2.75
			}]
		}
	}`)

	var resp AudioSyncDetectResp
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal AudioSyncDetectResp: %v", err)
	}

	got := resp.Detail.AudioDetail[0]
	if got.AudioStarttime != 1.5 {
		t.Fatalf("AudioStarttime = %v, want 1.5", got.AudioStarttime)
	}
	if got.AudioEndtime != 2.75 {
		t.Fatalf("AudioEndtime = %v, want 2.75", got.AudioEndtime)
	}
}

func TestDetectResponsesUnmarshalReviewResultMetadata(t *testing.T) {
	raw := []byte(`{"code":1100,"finalResult":1,"resultType":0}`)

	var textResp TextDetectResp
	if err := json.Unmarshal(raw, &textResp); err != nil {
		t.Fatalf("unmarshal TextDetectResp: %v", err)
	}
	if textResp.FinalResult != 1 || textResp.ResultType != 0 {
		t.Fatalf("TextDetectResp finalResult/resultType = %d/%d, want 1/0", textResp.FinalResult, textResp.ResultType)
	}

	var imageResp ImageDetectResp
	if err := json.Unmarshal(raw, &imageResp); err != nil {
		t.Fatalf("unmarshal ImageDetectResp: %v", err)
	}
	if imageResp.FinalResult != 1 || imageResp.ResultType != 0 {
		t.Fatalf("ImageDetectResp finalResult/resultType = %d/%d, want 1/0", imageResp.FinalResult, imageResp.ResultType)
	}
}

func TestTextDetectResponseUnmarshalsAccountAndLanguageMetadata(t *testing.T) {
	raw := []byte(`{
		"code": 1100,
		"langResult": {
			"detectedLang": "zh",
			"translatedText": "hello"
		},
		"tokenProfileLabels": [{
			"label1": "profile",
			"label2": "age",
			"label3": "adult",
			"description": "adult user",
			"timestamp": 1710000000000
		}],
		"tokenRiskLabels": [{
			"label1": "risk",
			"label2": "spam",
			"label3": "spam_user",
			"description": "spam user",
			"timestamp": 1710000000001
		}]
	}`)

	var resp TextDetectResp
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal TextDetectResp: %v", err)
	}
	if resp.LangResult.DetectedLang != "zh" || resp.LangResult.TranslatedText != "hello" {
		t.Fatalf("LangResult = %#v, want detectedLang zh and translatedText hello", resp.LangResult)
	}
	if len(resp.TokenProfileLabels) != 1 || resp.TokenProfileLabels[0].Timestamp != 1710000000000 {
		t.Fatalf("TokenProfileLabels = %#v, want one profile label", resp.TokenProfileLabels)
	}
	if len(resp.TokenRiskLabels) != 1 || resp.TokenRiskLabels[0].Timestamp != 1710000000001 {
		t.Fatalf("TokenRiskLabels = %#v, want one risk label", resp.TokenRiskLabels)
	}
}

func jsonEqual(got, want interface{}) bool {
	gotRaw, err := json.Marshal(got)
	if err != nil {
		return false
	}
	wantRaw, err := json.Marshal(want)
	if err != nil {
		return false
	}
	return string(gotRaw) == string(wantRaw)
}

func TestDefaultTimeoutIsThreeSeconds(t *testing.T) {
	client, err := NewClient("app", "key")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	if got := client.Timeout(); got != 3*time.Second {
		t.Fatalf("timeout = %v, want %v", got, 3*time.Second)
	}
}
