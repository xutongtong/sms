package utils

import (
	"time"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func InArray(key string, params []string) bool {
	for _, v := range params {
		if key == v {
			return true
		}
	}

	return false
}

func GenerateRandom(len int) string{
	i, _ := strconv.ParseInt("1" + strings.Repeat("0", len), 10, 32)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := fmt.Sprintf("%0" + (string(len)) + "v", random.Int63n(i))

	return number
}