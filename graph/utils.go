package graph

import (
	"math/rand"
	"strconv"
	"time"
)

func generateRandomCode(length int) (code string) {
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < length; i++ {
		code += strconv.FormatInt(int64(rand.Int31n(10)), 10)
	}
	return code
}
