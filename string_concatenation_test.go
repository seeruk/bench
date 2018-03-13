package main

import (
	"bytes"
    "fmt"
	"strings"
	"testing"
)

var strLen int = 1000

func BenchmarkConcatFmt(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var str string
        for i := 0; i < strLen; i++ {
            str = fmt.Sprintf("%s%s", str, "x")
        }
    }
}

func BenchmarkConcatString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var str string
		for i := 0; i < strLen; i++ {
			str += "x"
		}
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		for i := 0; i < strLen; i++ {
			buffer.WriteString("x")
		}
		buffer.String()
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var builder strings.Builder
		for i := 0; i < strLen; i++ {
			builder.WriteString("x")
		}
		builder.String()
	}
}
