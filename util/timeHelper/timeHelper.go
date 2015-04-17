package timeHelper

import (
	"time"

	"github.com/dorajistyle/goyangi/util/log"
)

func FewDaysLater(day int) time.Time {
	return FewDurationLater(time.Duration(day) * 24 * time.Hour)
}

func TwentyFourHoursLater() time.Time {
	return FewDurationLater(time.Duration(24) * time.Hour)
}

func SixHoursLater() time.Time {
	return FewDurationLater(time.Duration(6) * time.Hour)
}

func FewDurationLater(duration time.Duration) time.Time {
	// When Save time should considering UTC
	baseTime := time.Now()
	log.Debugf("basetime : %s", baseTime)
	fewDurationLater := baseTime.Add(duration)
	log.Debugf("time : %s", fewDurationLater)
	return fewDurationLater
}

func IsExpired(expirationTime time.Time) bool {
	baseTime := time.Now()
	log.Debugf("basetime : %s", baseTime)
	log.Debugf("expirationTime : %s", expirationTime)
	elapsed := time.Since(expirationTime)
	log.Debugf("elapsed : %s", elapsed)
	after := time.Now().After(expirationTime)
	log.Debugf("after : %t", after)
	return after
}
