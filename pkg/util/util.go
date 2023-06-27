package util

import (
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
