package gentwolf

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	rnd "math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var Util *util

type util struct {
}

func init() {
	Util = &util{}
}

func (this *util) ToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func (this *util) ToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (this *util) ToFloat32(str string) float32 {
	i, _ := strconv.ParseFloat(str, 64)
	return float32(i)
}

func (this *util) ToFloat64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
	return i
}

func (this *util) ToStr(n interface{}, args ...int) string {
	if len(args) == 0 {
		return fmt.Sprintf("%d", n)
	} else {
		format := "%." + fmt.Sprintf("%d", args[0]) + "f"
		return fmt.Sprintf(format, n)
	}
}

func (this *util) Int64ToStr(i int64) string {
	return this.ToStr(i)
}

func (this *util) Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (this *util) RndStr(length int) string {
	r := rnd.New(rnd.NewSource(time.Now().UnixNano()))
	rs := make([]string, length)
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		index := r.Intn(62)

		rs[i] = string(str[index])
	}
	return strings.Join(rs, "")
}

func (this *util) SubString(str string, start, length int) string {
	rs := []rune(str)
	strLen := len(rs)

	if start >= strLen {
		return ""
	}

	if start < 0 {
		start += strLen
	}

	if length < 0 {
		length = strLen + length - 1
	}

	end := start + length
	if end > strLen {
		end = strLen
	}

	s := rs[start:end]
	return string(s)
}

func (this *util) Base64DecodeToByte(str string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(str)
}

func (this *util) Base64EncodeByte(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (this *util) Base64Decode(str string) string {
	b, _ := base64.RawStdEncoding.DecodeString(str)
	return string(b)
}

func (this *util) Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (this *util) Ceil(size, count int) int {
	return int(math.Ceil(float64(count) / float64(size)))
}

func (this *util) Uuid() string {
	u := make([]byte, 16)
	if _, err := rand.Read(u); err != nil {
		panic(err)
	}
	u[6] = (u[6] & 0x0f) | (4 << 4)
	u[8] = (u[8] & 0xbf) | 0x80

	return hex.EncodeToString(u)
}

func (this *util) UrlValuesToMap(urlValues url.Values) map[string]string {
	params := make(map[string]string, len(urlValues))
	for k, v := range urlValues {
		params[k] = v[0]
	}
	return params
}

func (this *util) MapToUrlValues(data map[string]string) url.Values {
	params := url.Values{}

	for k, v := range data {
		params.Add(k, v)
	}
	return params
}
