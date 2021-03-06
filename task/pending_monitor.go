package task

import (
	"encoding/json"
	"fmt"
	"github.com/ismdeep/justoj-monitor/config"
	"github.com/ismdeep/justoj-monitor/notification"
	"github.com/ismdeep/log"
	"io/ioutil"
	"net/http"
	"time"
)

type responseData struct {
	Code uint `json:"code"`
	Data struct {
		PendingCnt   uint `json:"pending_cnt"`
		RejudgingCnt uint `json:"rejudging_cnt"`
		CompilingCnt uint `json:"compiling_cnt"`
		RunningCnt   uint `json:"running_cnt"`
	} `json:"data"`
}

var recovered = true
var duration time.Duration
var errDuration time.Duration

func check() error {
	// https://oj.ismdeep.com/api/solution/pending_cnt
	url := fmt.Sprintf("%v/api/solution/pending_cnt", config.JustOJ.Site)
	client := http.Client{}
	get, err := client.Get(url)
	if err != nil {
		return err
	}

	dataByte, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return err
	}

	data := &responseData{}
	if err := json.Unmarshal(dataByte, data); err != nil {
		return err
	}

	log.Info("check()", "url", url, "data", data)

	if data.Data.PendingCnt < config.PendingMonitor.PendingThreshold {
		// 一切正常
		recovered = true
		return nil
	}

	// 服务异常，需要发送通知邮件。
	log.Warn("check", "do", "notification.Send", "to", config.Notification.To)
	if recovered {
		err := notification.Send(&notification.Pack{
			SenderName: "JustOJ Monitor",
			Subject:    "JustOJ Notification",
			Type:       "text/plain",
			Content:    fmt.Sprintf("JustOJ服务器判题服务器出现大量未判题提交。判题服务器堵塞。请尽快处理！ PendingCnt: %v", data.Data.PendingCnt),
			ToMailList: config.Notification.To,
		})
		if err != nil {
			log.Warn("check", "err", err)
		}
		recovered = false
	}

	return nil
}

func daemon() {
	for {
		err := check()
		if err != nil {
			time.Sleep(errDuration)
			continue
		}
		time.Sleep(duration)
	}
}

func StartPendingMonitor() {
	var err error

	duration, err = time.ParseDuration(config.PendingMonitor.Duration)
	if err != nil {
		duration, _ = time.ParseDuration("5m")
	}

	errDuration, err = time.ParseDuration(config.PendingMonitor.ErrDuration)
	if err != nil {
		errDuration, _ = time.ParseDuration("1m")
	}

	go func() {
		daemon()
	}()
}
