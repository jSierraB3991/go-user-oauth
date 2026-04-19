package gooautherror

import "time"

type InvalidSessionTimeError struct {
	TimeMinutes time.Duration
}

func (i InvalidSessionTimeError) Error() string {
	return "INVALID_SESSION_TIME_ERROR"
}

func (i InvalidSessionTimeError) GetData() map[string]any {
	return map[string]any{
		"timeMinutes": i.TimeMinutes,
	}
}
