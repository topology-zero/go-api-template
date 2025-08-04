package util

import (
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

// CopyValue 数据复制
func CopyValue(toValue, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					if src == nil {
						return "", nil
					}
					s, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("CopyValue 时间转换错误")
					}
					return s.Format(time.DateTime), nil
				},
			},
		},
	})
}

type empty struct{}

var (
	pkgNameOnce sync.Once
	pkgName     string
)

func GetPkgName() string {
	pkgNameOnce.Do(func() {
		pkgNames := reflect.TypeOf(empty{}).PkgPath()
		split := strings.Split(pkgNames, "/")
		pkgName = split[0]
	})
	return pkgName
}

const charset = "0123456789"

func RandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

// RFC3339toDateTime 将 RFC3339 格式转换为 DateTime 格式
func RFC3339toDateTime(s string) string {
	if len(s) == 0 {
		return s
	}

	t, err := time.ParseInLocation(time.RFC3339, s, time.Local)
	if err != nil {
		return s
	}
	return t.Format(time.DateTime)
}
