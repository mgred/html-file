package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/mazen160/go-random"
)

func GenerateHash() string {
	if h := os.Getenv("HTML_FILE_HASH"); h != "" {
		return h
	}
	hash, err := random.String(8)
	if err != nil {
		return fmt.Sprintf("%d", time.Now().Unix())
	}
	return hash
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
