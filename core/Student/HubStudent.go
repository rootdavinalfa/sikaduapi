/*
 * Copyright (c) 2019 - 2020. dvnlabs.xyz
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package Student

import (
	"encoding/json"
	"net/http"
	"unbajaUAPI/libs"
	"unbajaUAPI/model"
)

func GetStudentInfoHub(w http.ResponseWriter, token string) interface{} {
	type InfoInterface struct {
		Error   bool
		Message string
		Info    interface{}
	}
	isError := true
	Messag := ""
	var infos interface{} = nil
	if token != "" {
		isOk, data, message := libs.VerifyToken(token)
		if isOk {
			jsosString := MustMarshal(data)
			var auth model.LoginAuth
			_ = json.Unmarshal(jsosString, &auth)
			isError = false
			Messag = "Successfully"
			infos = GetStudentInfo(auth.Cookie)
			if infos == nil {
				w.WriteHeader(http.StatusForbidden)
				isError = true
				Messag = "Not authorized"
			}
		} else {
			Messag = message
		}
	} else {
		Messag = "Token not found"
	}

	Inf := InfoInterface{
		Error:   isError,
		Message: Messag,
		Info:    infos,
	}
	return Inf
}

func GetStudentScheduleListHub(w http.ResponseWriter, token string) interface{} {
	type InfoInterface struct {
		Error   bool
		Message string
		Info    interface{}
	}
	isError := true
	Messag := ""
	var infos interface{} = nil
	if token != "" {
		isOk, data, message := libs.VerifyToken(token)
		if isOk {
			jsosString := MustMarshal(data)
			var auth model.LoginAuth
			_ = json.Unmarshal(jsosString, &auth)
			isError = false
			Messag = "Successfully"
			infos = GetStudentScheduleList(auth.Cookie)
			if infos == nil {
				w.WriteHeader(http.StatusForbidden)
				isError = true
				Messag = "Not authorized"
			}
		} else {
			Messag = message
		}
	} else {
		Messag = "Token not found"
	}

	Inf := InfoInterface{
		Error:   isError,
		Message: Messag,
		Info:    infos,
	}
	return Inf
}

func GetStudentScheduleHub(w http.ResponseWriter, token string, year string, quart string) interface{} {
	type InfoInterface struct {
		Error     bool
		Message   string
		StudentID string
		Info      interface{}
	}
	isError := true
	Messag := ""
	studentID := ""
	var infos interface{} = nil
	if token != "" {
		isOk, data, message := libs.VerifyToken(token)
		if isOk {
			if year == "" {
				w.WriteHeader(http.StatusBadRequest)
				isError = true
				Messag = "Year empty,please insert"
			} else if quart == "" {
				w.WriteHeader(http.StatusBadRequest)
				isError = true
				Messag = "Quart empty,please insert"
			} else if year != "" && quart != "" {
				jsosString := MustMarshal(data)
				var auth model.LoginAuth
				_ = json.Unmarshal(jsosString, &auth)
				isError = false
				Messag = "Successfully"
				infos = GetStudentSchedule(auth.Cookie, year, quart)
				if infos == nil {
					w.WriteHeader(http.StatusForbidden)
					isError = true
					Messag = "Not authorized"
				} else {
					studentID = auth.User
				}
			}
		} else {
			Messag = message
		}
	} else {
		Messag = "Token not found"
	}

	Inf := InfoInterface{
		Error:     isError,
		Message:   Messag,
		StudentID: studentID,
		Info:      infos,
	}
	return Inf
}

func GetStudentGradeDetailHub(w http.ResponseWriter, token string, year string, quart string) interface{} {
	type InfoInterface struct {
		Error     bool
		Message   string
		StudentID string
		Grade     interface{}
	}
	isError := true
	Messag := ""
	studentID := ""
	var infos interface{} = nil
	if token != "" {
		isOk, data, message := libs.VerifyToken(token)
		if isOk {
			if year == "" {
				w.WriteHeader(http.StatusBadRequest)
				isError = true
				Messag = "Year empty,please insert"
			} else if quart == "" {
				w.WriteHeader(http.StatusBadRequest)
				isError = true
				Messag = "Quart empty,please insert"
			} else if year != "" && quart != "" {
				jsosString := MustMarshal(data)
				var auth model.LoginAuth
				_ = json.Unmarshal(jsosString, &auth)
				isError = false
				Messag = "Successfully"
				infos = GetStudentGradeDetail(auth.Cookie, year, quart)
				if infos == nil {
					w.WriteHeader(http.StatusForbidden)
					isError = true
					Messag = "Not authorized"
				} else {
					studentID = auth.User
				}
			}
		} else {
			Messag = message
		}
	} else {
		Messag = "Token not found"
	}

	Inf := InfoInterface{
		Error:     isError,
		Message:   Messag,
		StudentID: studentID,
		Grade:     infos,
	}
	return Inf
}

func GetStudentGradeSummaryHub(w http.ResponseWriter, token string) interface{} {
	type InfoInterface struct {
		Error   bool
		Message string
		Grade   interface{}
	}
	isError := true
	Messag := ""
	var grades interface{} = nil
	if token != "" {
		isOk, data, message := libs.VerifyToken(token)
		if isOk {
			jsosString := MustMarshal(data)
			var auth model.LoginAuth
			_ = json.Unmarshal(jsosString, &auth)
			isError = false
			Messag = "Successfully"
			grades = GetStudentGradeSummary(auth.Cookie, auth.User)
			if grades == nil {
				w.WriteHeader(http.StatusForbidden)
				isError = true
				Messag = "Not authorized"
			}
		} else {
			Messag = message
		}
	} else {
		Messag = "Token not found"
	}

	Inf := InfoInterface{
		Error:   isError,
		Message: Messag,
		Grade:   grades,
	}
	return Inf
}

func GetStudentFinanceStatus(w http.ResponseWriter, token string) interface{} {
	type InfoInterface struct {
		Error   bool
		Message string
		Finance interface{}
	}
	isError := true
	Messag := ""
	var finance interface{} = nil
	if token != "" {
		isOk, data, message := libs.VerifyToken(token)
		if isOk {
			jsosString := MustMarshal(data)
			var auth model.LoginAuth
			_ = json.Unmarshal(jsosString, &auth)
			isError = false
			Messag = "Successfully"
			finance = GetFinanceStatus(auth.Cookie)
			if finance == nil {
				w.WriteHeader(http.StatusForbidden)
				isError = true
				Messag = "Not authorized"
			}
		} else {
			Messag = message
		}
	} else {
		Messag = "Token not found"
	}

	Inf := InfoInterface{
		Error:   isError,
		Message: Messag,
		Finance: finance,
	}
	return Inf
}

func MustMarshal(data interface{}) []byte {
	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return out
}
