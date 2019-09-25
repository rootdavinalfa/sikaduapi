/*
 * Copyright (c) 2019. dvnlabs.ml
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package model

type FinanceDetail struct {
	StudentID, Name string
	Bill            []FinanceBilled
}
type FinanceBilled struct {
	No                           int
	Period                       string
	First, Second, Third, Remain int
	Percentage                   float64
	Status                       string
}
