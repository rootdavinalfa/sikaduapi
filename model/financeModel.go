/*
 * Copyright (c) 2019 - 2020. dvnlabs.xyz
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
	Quart                        int
	Period                       string
	Charged                      int
	Paid                         int
	First, Second, Third, Remain int
	Percentage                   float64
	Status                       string
}
