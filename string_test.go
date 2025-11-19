package helper

import (
	"testing"
)

func TestSplitNumberAndChinese(t *testing.T) {
	val := "402472.289测试"
	number, chinese, ok := SplitNumberAndChinese(val)
	if !ok {
		t.Fatal("测试失败: ", val)
	}
	if number != "402472.289" || chinese != "测试" {
		t.Fatal("测试失败: ", number, chinese)
	}

	val = "42测试"
	_, _, ok = SplitNumberAndChinese(val)
	if ok {
		t.Fatal("测试失败: ", val)
	}
}
