package consts

import "time"

var (
	UserAccountStatus = struct {
		INACTIVE int8
		ACTIVE   int8
	}{
		INACTIVE: 0,
		ACTIVE:   1,
	}

	REDIS_OTP_EXPIRATION     = 60 * time.Second // 1 minute
	REDIS_DEFAULT_EXPIRATION = 60 * time.Minute // 1 hour
)
