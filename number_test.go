package helper

import (
	"fmt"
	"slices"
	"testing"
)

func TestRemoveNonDigits(t *testing.T) {
	compareIntArr := func(t testing.TB, want, got string) {
		t.Helper()

		if want != got {
			t.Errorf("期望为 '%s'，实际获得 '%s'", want, got)
		}
	}

	assertArr := map[string]string{
		"C2": "2",
		"C8": "8",
	}
	for k, v := range assertArr {
		result := RemoveNonDigits(k)
		var expected = v

		compareIntArr(t, expected, result)
	}
}

func TestRemoveSingleChinese(t *testing.T) {
	compareIntArr := func(t testing.TB, want, got []string) {
		t.Helper()

		if slices.Compare(want, got) != 0 {
			t.Errorf("期望为 '%s'，实际获得 '%s'", want, got)
		}
	}

	k := []string{"1", "寒假", "汉", "5.95"}
	v := []string{"1", "寒假", "5.95"}
	result := RemoveSingleChinese(k, []string{"买", "卖"})
	var expected = v
	compareIntArr(t, expected, result)

	k = []string{"1", "处", "开", "寒假", "汉", "5.95"}
	v = []string{"1", "寒假", "5.95"}
	result = RemoveSingleChinese(k, []string{"买", "卖"})
	expected = v
	compareIntArr(t, expected, result)

	k = []string{"1", "处", "买", "mai", "寒假", "汉", "5.95"}
	v = []string{"1", "买", "mai", "寒假", "5.95"}
	result = RemoveSingleChinese(k, []string{"买", "卖"})
	expected = v
	compareIntArr(t, expected, result)
}

func assertEqualNumber[T Number](t testing.TB, want, got T) {
	t.Helper()

	if want != got {
		t.Errorf("期望为 '%v'，实际获得 '%v'", want, got)
	}
}

func TestDiv(t *testing.T) {
	// 小数除法
	resultFloat := Div(1.0, 10.0)
	assertEqualNumber(t, 0.1, resultFloat)

	// 整数除法
	resultInt := Div(1, 10)
	assertEqualNumber(t, 0, resultInt)

	resultInt = Div(120, 10)
	assertEqualNumber(t, 12, resultInt)

	// 多参数的快速除法
	resultInt = Div(120, 10, 2)
	assertEqualNumber(t, 6, resultInt)

	resultFloat = Div(1.0, 4.0, 2.0)
	assertEqualNumber(t, 0.125, resultFloat)
}

func ExampleDiv() {
	result := Div(100, 2, 5)
	fmt.Println(result)
	// Output: 10
}
