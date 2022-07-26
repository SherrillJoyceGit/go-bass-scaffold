package logger

import (
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"strings"
	"time"
)

var currentLog *logrus.Logger

func CurrentLogger() *logrus.Logger {
	if currentLog == nil {
		NewLogrusLogger()
	}

	return currentLog
}

func NewLogrusLogger() fiber.Handler {

	cfg, err := getCurrentConfig()

	if err == nil {
		log.Fatal(2, "Fail to get lgs config: %v", err)
	}

	// Set variables
	var (
		errHandler fiber.ErrorHandler
	)

	lgs := logrus.New()

	// 设置等级，可配置
	lgs.SetLevel(logrus.DebugLevel)

	// 设置时间格式，可配置
	lgs.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	// 连接，从配置文件获取
	conn, err := net.Dial(string(cfg.Network), cfg.Host)
	if err != nil {
		lgs.Fatal(err)
	}

	// 从配置文件获取服务名 serviceId
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"type":      string(cfg.LogType),
		"serviceId": cfg.ServiceId,
	}))

	// 记录连接情况日志
	lgs.Hooks.Add(hook)
	lgs.WithFields(logrus.Fields{
		"method": "lgs-logstash-connect",
	}).Infof("connect to " + cfg.Host + " for logstash is ok")

	currentLog = lgs
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
