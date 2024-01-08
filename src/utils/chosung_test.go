package utils

import (
	"testing"
)

func TestExtractChosung(t *testing.T) {
	input := "안녕하세요 te 선생님"
	result := ExtractChosung(input)
	t.Logf("결과: %v", result)
}