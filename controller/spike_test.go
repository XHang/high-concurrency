package controller

import (
	logger "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func test_Buyer(t *testing.T, uid int64) {
	bytes := GetTest(t, "/grab/mask", map[string]string{"userID": strconv.FormatInt(uid, 10)})
	logger.Infof("响应数据是【%v】", string(bytes))
}
func Test_splke_GrabMask(t *testing.T) {
	for i := 1; i < 200; i++ {
		uid := Random(0, int64(i))
		go test_Buyer(t, uid)
	}
	time.Sleep(time.Second * 15)
}

//Random 随机数
func Random(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min+1)
}
