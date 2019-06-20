package log

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/log"
)

const logEventUrl = api.BASE_URL + "/log"

type LogEvent domain.LogEvent

func GetLogEvents() ([]LogEvent, error) {
	entites := []LogEvent{}
	resp, _ := api.Rest().Get(logEventUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}
