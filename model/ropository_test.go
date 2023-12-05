package model_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

func TestInsertMessage(t *testing.T) {
	// tsk := model.Task{Name: "task1", Description: "Do Something", Time: }
}

func TestGetByID(t *testing.T) {
	dt := "2023-12-05"
	tt := "09:00:00"

	tc := []struct {
		Name   string
		ID     int
		Rows   int
		Result model.Nofify
	}{
		{
			Name: "Get 1 task Success",
			ID:   1,
			Rows: 1,
			Result: model.Nofify{
				ID:      1,
				Message: "ส่งงาน",
				Date:    dt,
				Time:    tt,
			},
		},
	}

	t.Run(tc[0].Name, func(t *testing.T) {
		cfg := config.GetConfig()
		db := model.GetDB(cfg.DB)

		mes := db.GetByID(tc[0].ID)

		defer db.Close()

		assert.Equalf(t, tc[0].Rows, len(mes), "Expected %d got %d", tc[0].Rows, len(mes))
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
