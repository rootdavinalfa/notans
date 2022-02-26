/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package common

func FindfirstDigit(number int) int {
	for number >= 10 {
		number = number / 10
	}
	return number
}
