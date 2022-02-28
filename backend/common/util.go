/*
 * Copyright (c) 2021-2022.
 *
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * All rights reserved
 *
 * This program contains research , trial - errors. So this program can't guarantee your system will work as intended.
 */

package common

import (
	"math/rand"
	"time"
)

func FindfirstDigit(number int) int {
	for number >= 10 {
		number = number / 10
	}
	return number
}

func RandAlphaNumeric(length int) string {
	// Based On https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func IsStringEmpty(str string) bool {
	if len(str) > 0 {
		return false
	}
	return true
}
