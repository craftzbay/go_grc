package data

import (
	"database/sql/driver"
	"time"
)

type LocalTimeZone time.Time

const localTimeZoneFormat = "2006-01-02 15:04:05"

func (t *LocalTimeZone) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+localTimeZoneFormat+`"`, string(data), time.Local)
	*t = LocalTimeZone(now)
	return
}

func (t LocalTimeZone) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localTimeZoneFormat)+2)
	b = append(b, '"')
	b = append(b, []byte(t.String())...)
	b = append(b, '"')

	return b, nil
}

func (t LocalTimeZone) String() string {
	if time.Time(t).IsZero() {
		return "0000-00-00 00:00:00"
	}

	return time.Time(t).Format(localTimeZoneFormat)
}

func (t LocalTimeZone) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return time.Now().Local(), nil
	}
	return time.Time(t).Local(), nil
}

func (t *LocalTimeZone) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = LocalTimeZone(vt)
	case string:
		tTime, _ := time.Parse("2006-01-02 15:04:05", vt)
		*t = LocalTimeZone(tTime)
	default:
		return nil
	}
	return nil
}
