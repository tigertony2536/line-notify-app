package controller

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

var (
	cfg config.Config
	db  *model.DB
)

func init() {
	cfg = config.GetConfig()
	db = model.GetDB(cfg.DB)
}

func getWeekDay() (time.Time, time.Time) {
	start := time.Now()

	for start.Weekday() != time.Monday {
		start = start.Add(-time.Hour * 24)
	}

	end := start.Add(time.Hour * 144)
	return start, end
}

func SendNotification(noti model.Notification) (string, error) {
	v := url.Values{}
	v.Set("message", "test")
	client := &http.Client{}
	req, err := http.NewRequest("POST", cfg.Url, strings.NewReader(v.Encode()))

	token := "Bearer " + cfg.LineToken

	if err != nil {
		return "", err
	}
	// req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	s := string(respText)
	return s, nil
}

func GetWeeklyNoti() ([]model.Notification, error) {
	start, end := getWeekDay()

	noti, err := db.GetByDate(start.Format(time.DateOnly), end.Format(time.DateOnly))
	if err != nil {
		return []model.Notification{}, err
	}

	return noti, nil
}
