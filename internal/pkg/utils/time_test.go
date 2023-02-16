/*时间*/
package utils

import (
	"testing"
	"time"
)

func TestTimeLayout(t *testing.T) {
	logTmFmt := "2006-01-02 15:04:05.000"
	ti, err := time.ParseInLocation(logTmFmt, "2022-10-10 12:12:12.789", time.Local)
	if err != nil {
		t.Errorf("TimeLayout = %v, err %v", logTmFmt, err)
		return
	}
	t.Log(ti.Format(logTmFmt))
}
