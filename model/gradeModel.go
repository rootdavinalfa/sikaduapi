/*
 * Copyright (c) 2019 - 2020. dvnlabs.xyz
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package model

type GradeModelSummary struct {
	Year, Quart, Semester, Periodic, Detail string
	NumCourse, Credit                       int
	Cumulative                              float64
}
type GradeModel struct {
	StudentID string
	GPA       float64
	Data      []interface{}
}
type GradeModelFull struct {
	Year       string
	Quart      string
	Cumulative float64
	Data       []GradeModelDetail
}
type GradeModelDetail struct {
	CourseName, GradeLetter                                       string
	Num, Credit                                                   int
	Availability, Quiz, Assignment, MidTerm, LastTerm, GradePoint float64
}
