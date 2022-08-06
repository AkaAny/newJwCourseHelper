package core

import (
	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
	"time"
)

func (u *User) getCookie() string {
	if u.auth == nil {
		return ""
	}
	return u.auth.GetCookie()
}

func (u *User) getRequestRate() *rate.Limiter {
	return u.rateBucket
}

func (u *User) init() {
	u.formParam = make(map[string]string, 10)
	u.info = new(baseInfo)
	u.rateBucket = rate.NewLimiter(rate.Every(time.Duration(u.config.rate)*time.Second), u.config.bucketFull)
	u.config = new(missionConfig)
	u.cache = cache.New(1*time.Hour, 24*time.Hour)
	u.client = u.newRequestClient()
}

func (u *User) getCache() *cache.Cache {
	return u.cache
}

func (u *User) getField() map[string]string {
	return u.info.field
}

func (u *User) getBaseQuery() string {
	return u.info.baseQuery
}

func (u *User) SetTarget(r []struct {
	Name string `json:"name"`
	Type string `json:"type"`
}) *User {
	for _, target := range r {
		u.config.target = append(u.config.target, target)
	}
	u.e = nil
	return u
}

func (u *User) getTarget() []Target {
	return u.config.target
}
