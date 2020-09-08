package service

import (
	"errors"
	logger "github.com/sirupsen/logrus"
)

var count int64

func init() {
	count = 100
}

type mask struct{}

//Mask 实例
var Mask mask

//Buy 买一个
func (*mask) Buy(uid int64) error {
	if uid < 0 {
		return errors.New("该用户无权限")
	}
	if count < 1 {
		return errors.New("商品已售完")
	}
	if User.isBuyed(uid) {
		return errors.New("抢购机会已用完")
	}
	logger.Infof("购买用户是【%v】", uid)
	count--
	logger.Infof("口罩剩余【%v】", count)
	User.addBuyerRecord(uid)
	return nil
}
