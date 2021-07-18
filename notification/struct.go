package notification

import (
	"encoding/json"
	"fmt"
	"github.com/ismdeep/justoj-monitor/config"
	"io/ioutil"
	"net/http"
	"strings"
)

type Pack struct {
	SenderName string   `json:"sender_name"`
	Subject    string   `json:"subject"`
	Type       string   `json:"type"`
	Content    string   `json:"content"`
	ToMailList []string `json:"to_mail_list"`
}

// Send 发送
func Send(pack *Pack) error {
	jsonData, err := json.Marshal(pack)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/mails", config.Notification.Host), strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", config.Notification.Token)

	got, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}

	if _, err := ioutil.ReadAll(got.Body); err != nil {
		return err
	}

	return nil
}
