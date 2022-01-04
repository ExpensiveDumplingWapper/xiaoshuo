package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func Str2Int64(value string, defaultV int64) int64 {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultV
	}
	return v
}

func Str2Float64(value string, defaultV float64) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultV
	}
	return v
}

func Str2Int(value string, defaultV int) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		return defaultV
	}
	return v
}

func ReverseString(s string) string {
	r := []rune(s)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}

func Json2Struct(target interface{}, str string) interface{} {
	err := json.Unmarshal([]byte(str), &target)
	if err != nil {
		return nil
	}
	return target
}

func MyRand(n int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(n)
}
