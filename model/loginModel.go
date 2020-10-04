/*
 * Copyright (c) 2019 - 2020. dvnlabs.xyz
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package model

type BasicStudentInfo struct {
	NIK, NPM, Name, PlaceBorn, BornOn, Gender, Religion, Phone, Email, Address, ProfilePict string
	College                                                                                 StudentInfoOnCollege
}
type StudentInfoOnCollege struct {
	Faculty, Branch, Degree, Class, Group, Status string
}
type LoginAuth struct {
	User   string
	Cookie string
}
