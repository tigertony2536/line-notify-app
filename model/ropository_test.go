package model_test

import (
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

func TestInsertNotification(t *testing.T) {
	cfg := config.GetConfig()
	db := model.GetDB(cfg.DB)
	tm := time.Now().Format(time.TimeOnly)

	noti := model.Notification{
		Message: "ส่งคลิปจิตอาสา",
		Date:    "2023-12-30",
		Time:    tm,
	}

	id, err := db.InsertNotification(noti)

	expectedNoti, _ := db.GetByID(id)

	assert.Equalf(t, expectedNoti.Message, noti.Message, "Expect %q got %q", expectedNoti.Message, noti.Message)
	assert.Equalf(t, expectedNoti.Date, noti.Date, "Expect %q got %q", expectedNoti.Date, noti.Date)
	assert.Equalf(t, expectedNoti.Time, noti.Time, "Expect %q got %q", expectedNoti.Time, noti.Time)
	assert.NoError(t, err, "Insert notification to database successfully")
}

func TestGetByID(t *testing.T) {
	tc := struct {
		Name    string
		ID      int
		Message string
		Date    string
		Time    string
	}{
		Name:    "Get by ID Success",
		ID:      19,
		Message: "ส่งคลิปจิตอาสา",
		Date:    "2023-12-30",
		Time:    "09:00:00",
	}

	t.Run(tc.Name, func(t *testing.T) {

		cfg := config.GetConfig()
		db := model.GetDB(cfg.DB)

		noti, err := db.GetByID(19)

		defer db.Close()

		assert.Equalf(t, tc.ID, noti.ID, "Expected %q got %q", tc.ID, noti.ID)
		assert.Equalf(t, tc.Message, noti.Message, "Expected %q got %q", tc.Message, noti.Message)
		assert.Equalf(t, tc.Date, noti.Date, "Expected %q got %q", tc.Date, noti.Date)
		assert.Equalf(t, tc.Time, noti.Time, "Expected %q got %q", tc.Time, noti.Time)
		assert.NoError(t, err, "No error")
	})

}

func TestGetDB(t *testing.T) {
	cfg := config.GetConfig()
	db := model.GetDB(cfg.DB)
	db.Ping()
	expectType := model.DB{}

	t.Run("get db success", func(t *testing.T) {
		assert.IsTypef(t, &expectType, db, "Expected %T  got %T", &expectType, db)
		assert.NoErrorf(t, db.Ping(), "Connect db successfully")
	})

}
