/*
 * Copyright (c) 2019 - 2020. dvnlabs.xyz
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package model

type ScheduleStudentPeriode struct {
	CourseName, Class, Room, Lecturer, Days string
	Semester                                int
	Times                                   ScheduleTime
}
type ScheduleTime struct {
	FromTime, ToTime string
}
type ScheduleFull struct {
	Year, Quart string
	Data        []interface{}
}
type ScheduleList struct {
	SemesterAttended int
	List             []ScheduleListDetail
}
type ScheduleListDetail struct {
	Name, Year, Quart string
}
