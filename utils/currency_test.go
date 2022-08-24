package utils

import (
	"fmt"
	"testing"
)

func TestCurrency_Int2Float(t *testing.T) {
	ret, str := Currency().Int2Float(280, 2)
	fmt.Println(ret) // 2.8
	fmt.Println(str) // "2.80"
}

func TestCurrency_Float2IntRoundBank(t *testing.T) {
	// 四舍
	ret := Currency().Float2IntRoundBank(2.804, 2)
	fmt.Println(ret) // 280

	// 六入
	ret = Currency().Float2IntRoundBank(2.806, 2)
	fmt.Println(ret) // 281

	// 偶不进
	ret = Currency().Float2IntRoundBank(2.805, 2)
	fmt.Println(ret) // 280

	// 奇进
	ret = Currency().Float2IntRoundBank(2.815, 2)
	fmt.Println(ret) // 282

}

func TestCurrency_Float2IntRound(t *testing.T) {
	ret := Currency().Float2IntRound(2.804, 2)
	fmt.Println(ret) // 280
	ret = Currency().Float2IntRound(2.805, 2)
	fmt.Println(ret) // 281
}

func TestCurrency_Float2IntCeil(t *testing.T) {
	ret := Currency().Float2IntCeil(2.801, 2)
	fmt.Println(ret) // 281
}

func TestCurrency_Float2IntTruncate(t *testing.T) {
	ret := Currency().Float2IntTruncate(2.809, 2)
	fmt.Println(ret) // 280
}
