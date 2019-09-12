/*
 * Copyright (c) 2019. dvnlabs.ml
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package model

type ScheduleStudentPeriode struct {
	CourseName, Class, Room, Lecturer, Days string
	Semester                                int
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
