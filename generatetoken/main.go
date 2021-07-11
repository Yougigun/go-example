package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	token := GenerateToken("gary")
	fmt.Println("token",token,"len:",len(token))
}

var incrCounter uint64
func GenerateToken(str string) string {
	var p []byte
	p = append(p, str...)
	p = strconv.AppendInt(p, time.Now().UnixNano(),10)
	p = strconv.AppendUint(p,atomic.AddUint64(&incrCounter, 1), 10)
	md5hash := md5.Sum(p)
	md5hashStr := hex.EncodeToString(md5hash[:])
	return md5hashStr
}