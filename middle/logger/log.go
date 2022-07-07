package logger

import (
	"github.com/SherrillJoyceGit/go-bass-scaffold/config"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
	"time"
)

var currentLog *logrus.Logger

// 返回当前的日志记录器，默认为LogrusLogger
func LoggerCurrent() *logrus.Logger {
	if currentLog == nil {
		NewLogrusLogger()
	}

	return currentLog
}

func NewLogrusLogger() fiber.Handler {

	// Set variables
	var (
		errHandler fiber.ErrorHandler
	)

	log := logrus.New()

	// 设置等级，可配置
	log.SetLevel(logrus.DebugLevel)

	// 设置时间格式，可配置
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	// 连接，从配置文件获取
	conn, err := net.Dial(config.LogStashConfig.Net, config.LogStashConfig.Host)
	if err != nil {
		log.Fatal(err)
	}

	// 从配置文件获取服务名 serviceId
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"type":      "logrus",
		"serviceId": config.LogStashConfig.ServiceId,
	}))

	// 记录连接情况日志
	log.Hooks.Add(hook)
	log.WithFields(logrus.Fields{
		"method": "log-logstash-connect",
	}).Infof("connect to " + config.LogStashConfig.Host + " for logstash is ok")

	currentLog = log
	// 默认记录出入参，并限定业务接口
	return func(c *fiber.Ctx) error {
		// 请求开始时间
		requestTime := time.Now()

		// 处理请求
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		// 结束时间
		responseTime := time.Now()

		// 执行时间
		latencyTime := responseTime.Sub(requestTime).Milliseconds()

		// 请求方式
		reqMethod := c.Method()

		// 请求路由
		reqUrl := c.Request().URI().String()

		// 状态码
		statusCode := c.Response().StatusCode()

		//处理请求参数
		requestBody := string(c.Request().Body())

		//返回参数
		responseBody := string(c.Response().Body())

		if (!strings.Contains(reqUrl, "/swagger/")) && (!strings.Contains(reqUrl, "/health")) {
			if strings.Contains(responseBody, "error") {
				currentLog.WithFields(logrus.Fields{
					"method":      reqUrl,
					"latencyTime": latencyTime,
					"reqMethod":   reqMethod,
					"statusCode":  statusCode,
					"param":       requestBody,
					"result":      responseBody,
				}).Error(chainErr)
			} else {
				currentLog.WithFields(logrus.Fields{
					"method":      reqUrl,
					"latencyTime": latencyTime,
					"reqMethod":   reqMethod,
					"statusCode":  statusCode,
					"param":       requestBody,
					"result":      responseBody,
				}).Info(chainErr)
			}
		}
		return chainErr
	}
}
