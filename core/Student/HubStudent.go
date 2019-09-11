/*
 * Copyright (c) 2019. dvnlabs.ml
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
func MustMarshal(data interface{}) []byte {
	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return out
}
