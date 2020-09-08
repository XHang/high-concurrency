package service

type user struct{}

//User 实例
var User user

var record = map[int64]struct{}{}

func (*user) addBuyerRecord(uid int64) {
	record[uid] = struct{}{}
}

func (*user) isBuyed(uid int64) bool {
	_, has := record[uid]
	return has
}
