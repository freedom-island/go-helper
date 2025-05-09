package tests

import (
	"github.com/freedom-island/go-helper"
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
		result := go_helper.RemoveNonDigits(k)
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
	result := go_helper.RemoveSingleChinese(k, []string{"买", "卖"})
	var expected = v
	compareIntArr(t, expected, result)

	k = []string{"1", "处", "开", "寒假", "汉", "5.95"}
	v = []string{"1", "寒假", "5.95"}
	result = go_helper.RemoveSingleChinese(k, []string{"买", "卖"})
	expected = v
	compareIntArr(t, expected, result)

	k = []string{"1", "处", "买", "mai", "寒假", "汉", "5.95"}
	v = []string{"1", "买", "mai", "寒假", "5.95"}
	result = go_helper.RemoveSingleChinese(k, []string{"买", "卖"})
	expected = v
	compareIntArr(t, expected, result)
}
