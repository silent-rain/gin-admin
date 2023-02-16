/*时间
 */
package utils

import (
	"fmt"
	"time"

	"database/sql/driver"
)

// LocalTime grom 数据库查询时间格式设置
type LocalTime time.Time

// MarshalJSON 重写 MarshalJSON ⽅法来实现数据解析
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Value⽅法即在存储时调⽤，将该⽅法的返回值进⾏存储，该⽅法可以实现数据存储前对数据进⾏相关操作。
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// Scan⽅法可以实现在数据查询出来之前对数据进⾏相关操作。
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
