// Package time 时间工具
package time

import (
	"testing"
	"time"
)

func TestRFC3339ToCSTLayout(t *testing.T) {
	t.Log(RFC3339ToCSTLayout("2020-11-08T08:18:46+08:00"))
}

func TestCSTLayoutString(t *testing.T) {
	t.Log(CSTLayoutString())
}

func TestCSTLayoutStringToUnix(t *testing.T) {
	t.Log(CSTLayoutStringToUnix("2020-01-24 21:11:11"))
}

func TestGMTLayoutString(t *testing.T) {
	t.Log(GMTLayoutString())
}

func TestTimeLayout(t *testing.T) {
	logTmFmt := "2006-01-02 15:04:05.000"
	ti, err := time.ParseInLocation(logTmFmt, "2022-10-10 12:12:12.789", time.Local)
	if err != nil {
		t.Errorf("TimeLayout = %v, err %v", logTmFmt, err)
		return
	}
	t.Log(ti.Format(logTmFmt))
}
