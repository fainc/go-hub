// 货币存储转换工具

package utils

import (
	"github.com/shopspring/decimal"
)

type currency struct{}

var currencyVar = currency{}

func Currency() *currency {
	return &currencyVar
}

// Int2Float int转float
func (rec *currency) Int2Float(x int64, p int32) (float64, string) {
	u := decimal.New(10, p-1) // u = 10 ^ (p - 1)
	ret := decimal.NewFromInt(x).Div(u).Truncate(p)
	f, _ := ret.Float64()
	return f, ret.StringFixed(p)
}

// Float2IntTruncate float转int 去尾法
func (rec *currency) Float2IntTruncate(x float64, p int32) (ret int64) {
	u := decimal.New(10, p-1) // u = 10 ^ (p - 1)
	ret = decimal.NewFromFloat(x).Mul(u).IntPart()
	return
}

// Float2IntRound float转int 四舍五入法
func (rec *currency) Float2IntRound(x float64, p int32) (ret int64) {
	u := decimal.New(10, p-1) // u = 10 ^ (p - 1)
	ret = decimal.NewFromFloat(x).Mul(u).Round(0).IntPart()
	return
}

// Float2IntCeil float转int 进一法
func (rec *currency) Float2IntCeil(x float64, p int32) (ret int64) {
	u := decimal.New(10, p-1) // u = 10 ^ (p - 1)
	ret = decimal.NewFromFloat(x).Mul(u).Ceil().IntPart()
	return
}

// Float2IntRoundBank float转int 银行家法 四舍六入五成双 奇进偶不进
func (rec *currency) Float2IntRoundBank(x float64, p int32) (ret int64) {
	u := decimal.New(10, p-1) // u = 10 ^ (p - 1)
	ret = decimal.NewFromFloat(x).Mul(u).RoundBank(0).IntPart()
	return
}
