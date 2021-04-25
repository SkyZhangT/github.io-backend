package handler

import (
	"app/config"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)


  
func Limiter(conf config.LimiterConfiguration) *limiter.Limiter{
	lmt := tollbooth.NewLimiter(float64(conf.Max), nil)

	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}).
	SetMethods(conf.Methods).
	SetMessage(conf.Message).
	SetMessageContentType(conf.ContentType).
	SetTokenBucketExpirationTTL(time.Duration(conf.TokenBucketTTL))

	return lmt
}