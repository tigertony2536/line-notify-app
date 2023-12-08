package service

import (
	"time"

	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

func getWeekDay() (time.Time, time.Time) {
	start := time.Now()

	for start.Weekday() != time.Monday {
		start = start.Add(-time.Hour * 24)
	}

	end := start.Add(time.Hour * 144)
	return start, end
}

func SendNotification() bool {

	// client := &http.Client{}
	// res, err := http.Post("https://notify-api.line.me/api/notify", "multipart/form-data", noti)
	// panic(err)
	// re
	return true
}

func GetWeeklyNoti() ([]model.Notification, error) {
	start, end := getWeekDay()

	cfg := config.GetConfig()
	db := model.GetDB(cfg.DB)

	defer db.Close()
	noti, err := db.GetByDate(start.Format(time.DateOnly), end.Format(time.DateOnly))
	if err != nil {
		return []model.Notification{}, err
	}

	return noti, nil
}
