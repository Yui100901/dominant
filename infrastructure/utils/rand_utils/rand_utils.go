package rand_utils

import (
	"math/rand"
	"time"
)

//
// @Author yfy2001
// @Date 2024/9/13 12 47
//

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

func RandomFromSlice[T any](slice []T) T {
	return slice[r.Intn(len(slice))]
}
