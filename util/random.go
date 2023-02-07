package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//Return random full_name
func RandomFullname() string {
	return RandomString(8)
}

//Return random email
func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

//Return random mobile number
func RandomMobile() string {
	return fmt.Sprintf("%d", RandomInt(1000000000, 9999999999))
}

//Return random  password
func RandomPassword() string {
	return RandomString(8)
}

//
func RandomBookingDates() (time.Time, time.Time) {
	startDate := time.Now().UTC().Add(time.Duration(RandomInt(0, 5)) * time.Hour)
	endDate := startDate.UTC().Add(time.Duration(RandomInt(1, 5)) * time.Hour)
	return startDate, endDate
}
