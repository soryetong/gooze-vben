package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"gooze-vben-api/internal/admin/logic"
	"gooze-vben-api/models"

	"github.com/gin-gonic/gin"
	"github.com/soryetong/gooze-starter/gooze"
	"github.com/soryetong/gooze-starter/pkg/gzhttp"
	"github.com/soryetong/gooze-starter/pkg/gzutil"
	"github.com/soryetong/gooze-starter/services/gzauth"
	"go.uber.org/zap"
)

const (
	maxBodySize = 2 << 20 // 2MB
)

var notRecord = map[string]struct{}{}

func Record() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		path, id := gzutil.GetRequestPath(ctx.Request.URL.Path, "/api")
		if _, ok := notRecord[path]; ok || ctx.Request.Method == gzhttp.Method_GET {
			ctx.Next()
			return
		}

		// 请求参数采集处理
		reqData := make(map[string]interface{})
		if id != 0 {
			reqData["id"] = id
		}
		for k, v := range ctx.Request.URL.Query() {
			if len(v) == 1 {
				reqData[k] = v[0]
			} else {
				reqData[k] = v
			}
		}

		if ctx.Request.Body != nil && ctx.Request.ContentLength > 0 && ctx.Request.ContentLength <= maxBodySize {
			ct := ctx.GetHeader("Content-Type")
			if strings.Contains(ct, "application/json") ||
				strings.Contains(ct, "application/x-www-form-urlencoded") {
				bodyBytes, err := io.ReadAll(ctx.Request.Body)
				if err == nil {
					ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
					var bodyJSON map[string]interface{}
					if json.Unmarshal(bodyBytes, &bodyJSON) == nil {
						maskSensitive(bodyJSON)
						reqData["body"] = bodyJSON
					} else {
						reqData["body"] = string(bodyBytes)
					}
				}
			}
		}
		reqJson, _ := json.Marshal(reqData)

		writer := &responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBuffer(nil),
		}
		ctx.Writer = writer
		ctx.Next()

		userAgent := ctx.GetHeader("User-Agent")
		newPath := strings.TrimPrefix(ctx.Request.URL.Path, "/api/v1")
		record := models.SysRecords{
			Method:   ctx.Request.Method,
			Path:     path,
			Request:  string(reqJson),
			UserId:   gzauth.GetTokenValue[int64](ctx, "id"),
			Username: gzauth.GetTokenValue[string](ctx, "username"),
			Platform: gzutil.GetPlatform(userAgent) + " " + gzutil.GetBrowser(userAgent),
			Description: new(logic.SystemLogic).GetRecordDescription(
				gzutil.ConvertToRestfulURL(newPath),
				strings.ToUpper(ctx.Request.Method),
			),
			Ip:         gzutil.GetClientRealIP(ctx),
			Elapsed:    fmt.Sprintf("%.2f", time.Since(startTime).Seconds()*1000),
			StatusCode: int64(ctx.Writer.Status()),
		}

		var resp gooze.Response
		if err := json.Unmarshal(writer.body.Bytes(), &resp); err == nil {
			record.Msg = resp.Msg
			if data, err := json.Marshal(resp.Data); err == nil {
				record.Response = string(data)
			}
		} else {
			record.Response = writer.body.String()
		}

		asyncSaveRecord(record)
	}
}

func asyncSaveRecord(record models.SysRecords) {
	gzutil.SafeGo(func() {
		if err := gooze.Gorm().Create(&record).Error; err != nil {
			gooze.Log.Error(
				"记录操作日志失败",
				zap.Error(err),
				zap.Any("record", record),
			)
		}
	})
}

func maskSensitive(data map[string]interface{}) {
	sensitiveKeys := []string{
		"password",
		"pwd",
		"token",
		"access_token",
		"refresh_token",
	}

	for _, key := range sensitiveKeys {
		if _, ok := data[key]; ok {
			data[key] = "***"
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseBodyWriter) WriteHeader(statusCode int) {
	if !w.Written() {
		w.ResponseWriter.WriteHeader(statusCode)
	}
}
