package eyes

import (
	"strings"
	"time"
)

func (l *Logger)GetNowTime() string {
	nowTime := time.Unix(time.Now().Unix(), 0).In(l.TimeZone)
	timeStr := nowTime.Format(DateFormat)
	return timeStr
}

// HMS means Hour, Minute and Second, without these.
func (l *Logger)GetNowTimeWithoutHMS() string {
	nowTimeWithoutHMS := strings.Fields(l.GetNowTime())[0]
	return nowTimeWithoutHMS
}