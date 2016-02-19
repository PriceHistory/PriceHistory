package util

import (
	"testing"
)

func TestConvertPrice(t *testing.T) {
	convertedPrice := ConvertPrice("1 178 грн")
	if convertedPrice != 1178 {
		t.Error("ConvertPrice(\"1 178 грн\") != 1178")
	}
}

func TestConvertPrice2(t *testing.T) {
	convertedPrice := ConvertPrice("1 грн")
	if convertedPrice != 1 {
		t.Error("ConvertPrice(\"1 грн\") != 1")
	}
}

func TestConvertPrice3(t *testing.T) {
	convertedPrice := ConvertPrice("6 099 грн")
	if convertedPrice != 6099 {
		t.Error("ConvertPrice(\"6099 грн\") != 6099")
	}
}

func TestConvertPriceWithThinSpace(t *testing.T) {
	convertedPrice := ConvertPrice("7\u2009077\u2009грн")
	if convertedPrice != 7077 {
		t.Error("ConvertPrice(\"7\u2009077\u2009грн\") != 7077")
	}
}
