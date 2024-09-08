package middleware

import (
	"database/sql"
	"fmt"
	"server/infrastructure"
	"server/internal/model"
	"server/util"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// LoggerMiddleware middleware for logging incoming requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := requestid.Get(c)
		start := time.Now()

		fmt.Println(start.Format("02-01-2006 15:04:05 MST"), "|", rid)
		defer func() {
			user := util.CurrentUser(c)
			err := util.GetRequestError(c)
			stop := time.Now()

			statusCode := c.Writer.Status()
			path := c.Request.URL.Path
			method := c.Request.Method

			log_message := fmt.Sprintf("%v | %v | %v | %v | %v", start.Format("15:04 MST"), statusCode, stop.Sub(start).String(), method, path)
			hlog.Info(log_message)

			modelLog := model.Log{
				Id:         rid,
				Path:       path,
				Method:     method,
				StatusCode: statusCode,
				Interval:   time.Since(start).Seconds(),
				UserId:     sql.NullString{String: user.UserId, Valid: util.NewIsValid().String(&user.UserId)},
				Body:       sql.NullString{String: util.NewConvert().AnyToStr(err), Valid: util.NewIsValid().Any(err)},
			}
			go func(m model.Log) {
				infrastructure.SqlxDB.NamedExec(`
				insert into logs (uuid, path, method, status_code, interval, user_uuid, body) values (:id, :path, :method, :status_code, :interval, :user_id, :body)
				`, m)
			}(modelLog)
		}()
		c.Next()
	}
}
