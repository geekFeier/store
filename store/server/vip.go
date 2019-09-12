package serve

import "time"

func overTime(t int64, years, months, days int) bool {
	now := time.Now().Unix()
	after := time.Unix(t, 0).AddDate(years, months, days).Unix()

	if after > now {
		return true
	}
	return false
}

func isVip(login string) bool {
	vip := &VIP{Login: login}
	ok, err := vip.Get(login)
	if err != nil || !ok {
		return false
	}
	if overTime(vip.Date, 1, 0, 0) {
		return false
	}

	return true
}
