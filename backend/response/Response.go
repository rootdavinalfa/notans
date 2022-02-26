/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package response

type Response struct {
	Error   bool
	Message string
	Data    interface{}
}
