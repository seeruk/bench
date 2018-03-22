package reader

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func BenchmarkReading4MFile(b *testing.B) {
	var bs []byte

	for i := 0; i < b.N; i++ {
		// Generate this file with: `$ dd if=/dev/urandom of=/tmp/4m bs=1K count=4096`
		file, err := os.Open("/tmp/4m")
		if err != nil {
			b.Error(err)
		}

        //bs, err = ioutil.ReadAll(io.LimitReader(file, 4194000))
        bs, err = ioutil.ReadAll(io.LimitReader(file, 128))
		if err != nil {
            file.Close()
			b.Error(err)
		}

        file.Close()
	}

	b.Log(string(bs))
}

func BenchmarkReadingString(b *testing.B) {
	var bs []byte
	var err error

	input := "Hello, 世界"
	for i := 0; i < b.N; i++ {
		rdr := strings.NewReader(input)

		bs, err = ioutil.ReadAll(rdr)
		if err != nil {
			b.Error(err)
		}
	}

	_ = bs
}
