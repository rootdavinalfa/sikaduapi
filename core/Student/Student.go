/*
 * Copyright (c) 2019. dvnlabs.ml
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package Student

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"unbajaUAPI/model"
)

func GetStudentInfo(cookieVal string) interface{} {
	url := "http://sikadu.unbaja.ac.id/mahasiswa/"
	document := MakeRequest(url, cookieVal)

	religion := ""
	address := ""
	profilePict := ""
	var infose []string = nil
	// Find all links and process them with the function
	// defined earlier
	//Data except religion and address
	document.Find(".form-control").Each(func(index int, element *goquery.Selection) {
		// See if the href attribute exists on the element
		divForm, exists := element.Attr("value")
		if exists {
			infose = append(infose, divForm)
		}
	})
	//religion
	document.Find("select").Each(func(i int, s *goquery.Selection) {
		s.ChildrenFiltered("option").Each(func(i int, s *goquery.Selection) {
			_, ok := s.Attr("selected")
			if ok {
				religion = s.Text()
			}
		})
		if religion == "" {
			s.ChildrenFiltered("option").EachWithBreak(func(i int, s *goquery.Selection) bool {
				religion = s.Text()
				return false
			})
		}
	})
	//Address
	document.Find("textarea").Each(func(i int, selection *goquery.Selection) {
		address = selection.Text()
	})
	//Profilepict
	document.Find(".img-responsive").Each(func(i int, selection *goquery.Selection) {
		profilePict, _ = selection.Attr("src")
		if profilePict != "" {
			profilePict = "http://sikadu.unbaja.ac.id" + profilePict
		}
	})
	if (infose != nil && address != "" && religion != "") || profilePict != "" {
		data := model.BasicStudentInfo{
			NPM:         infose[0],
			Name:        infose[1],
			PlaceBorn:   infose[2],
			BornOn:      infose[3],
			Gender:      infose[4],
			Religion:    religion,
			Phone:       infose[5],
			Email:       infose[6],
			Address:     address,
			ProfilePict: profilePict,
			College: model.StudentInfoOnCollege{
				Faculty: infose[7],
				Branch:  infose[8],
				Degree:  infose[9],
				Class:   infose[10],
				Group:   infose[11],
				Status:  infose[12],
			},
		}
		return data
	}
	return nil
}
func GetStudentSchedule(cookieVal string, year string, quart string) interface{} {
	url := "http://sikadu.unbaja.ac.id/mahasiswa/akademik/jadwal?periode=" + year + quart
	document := MakeRequest(url, cookieVal)
	//Schedule parsing
	scheduleM := map[int][]string{}
	num := 0
	isErr := true
	document.Find(".col-md-12").Each(func(i int, selection *goquery.Selection) {
		selection.ChildrenFiltered("h3").Each(func(i int, selection *goquery.Selection) {
			if selection.Text() != "" {
				isErr = false
			}
		})
	})

	//println("SMT: "+strconv.Itoa(numOfSemester))
	document.Find("tbody").Each(func(i int, s *goquery.Selection) {
		s.ChildrenFiltered("tr").Each(func(i int, tr *goquery.Selection) {
			var scheduler []string = nil
			tr.ChildrenFiltered("td").Each(func(i int, td *goquery.Selection) {
				scheduler = append(scheduler, td.Text())
			})
			scheduleM[num] = scheduler
			num++
		})
	})
	if !isErr {
		var datas model.ScheduleFull
		for i := 0; i < len(scheduleM); i++ {
			smt, _ := strconv.Atoi(scheduleM[i][1])
			datas.Data = append(datas.Data, model.ScheduleStudentPeriode{
				CourseName: scheduleM[i][0],
				Class:      scheduleM[i][2],
				Room:       scheduleM[i][4],
				Lecturer:   scheduleM[i][3],
				Days:       scheduleM[i][5],
				Semester:   smt,
			})
		}

		data := model.ScheduleFull{
			Year:  year,
			Quart: quart,
			Data:  datas.Data,
		}
		return data
	}
	return nil
}

//Get list academic period
func GetStudentScheduleList(cookieVal string) interface{} {
	url := "http://sikadu.unbaja.ac.id/mahasiswa/akademik/jadwal"
	document := MakeRequest(url, cookieVal)
	numOfSemester := 0

	isErr := true
	document.Find(".col-md-12").Each(func(i int, selection *goquery.Selection) {
		selection.ChildrenFiltered("h3").Each(func(i int, selection *goquery.Selection) {
			if selection.Text() != "" {
				isErr = false
			}
		})
	})
	var datas model.ScheduleList
	document.Find("select").Each(func(i int, selection *goquery.Selection) {
		selection.ChildrenFiltered("option").Each(func(i int, selection *goquery.Selection) {
			numOfSemester++
			val, _ := selection.Attr("value")
			baseUrl, _ := url2.Parse(val)
			query := baseUrl.Query()
			period := query.Get("periode")
			year := period[:4]
			quart := period[len(period)-1:]
			datas.List = append(datas.List, model.ScheduleListDetail{
				Name:  selection.Text(),
				Year:  year,
				Quart: quart,
			})
		})
	})
	if !isErr {
		data := model.ScheduleList{
			SemesterAttended: numOfSemester,
			List:             datas.List,
		}
		return data
	}
	return nil
}

/*Section for grade getter
GetStudentGradeSummary for getting your average grade on all academic period you attented
GetStudentGradeDetail for getting grade for individual course grade*/

func GetStudentGradeSummary(cookieVal string, studentID string) interface{} {
	url := "http://sikadu.unbaja.ac.id/mahasiswa/akademik/khs"
	document := MakeRequest(url, cookieVal)

	isErr := true

	//Check is login or not by checking h3 tag
	document.Find(".col-md-12").Each(func(i int, selection *goquery.Selection) {
		selection.ChildrenFiltered("h3").Each(func(i int, selection *goquery.Selection) {
			if selection.Text() != "" {
				isErr = false
			}
		})
	})
	gradesM := map[int][]string{}
	num := 0
	//Select element
	document.Find("tbody").Each(func(i int, s *goquery.Selection) {
		s.ChildrenFiltered("tr").Each(func(i int, tr *goquery.Selection) {
			var grader []string = nil
			tr.ChildrenFiltered("td").Each(func(i int, td *goquery.Selection) {
				grader = append(grader, td.Text())
			})
			gradesM[num] = grader
			num++
		})
	})
	if !isErr {
		var datas model.GradeModel
		for i := 0; i < len(gradesM); i++ {
			numCourse, _ := strconv.Atoi(gradesM[i][1])
			credit, _ := strconv.Atoi(gradesM[i][2])
			cumulative, _ := strconv.ParseFloat(gradesM[i][3], 64)
			periodic := gradesM[i][0]
			year := periodic[:4]
			var quart string
			if i%2 == 0 && num%2 == 0 {
				quart = "1"
			} else {
				quart = "2"
			}
			datas.Data = append(datas.Data, model.GradeModelSummary{
				Year:       year,
				Quart:      quart,
				Semester:   strconv.Itoa(i + 1),
				Periodic:   periodic,
				Detail:     gradesM[i][4],
				NumCourse:  numCourse,
				Credit:     credit,
				Cumulative: cumulative,
			})
		}

		data := model.GradeModel{
			StudentID: studentID,
			Data:      datas.Data,
		}
		return data
	}
	return nil
}
func GetStudentGradeDetail(cookieVal string, year string, quart string) interface{} {
	hash := md5.New()

	_, _ = io.WriteString(hash, year+quart) // append into the hash

	url := "http://sikadu.unbaja.ac.id/mahasiswa/akademik/khs/view/" + hex.EncodeToString(hash.Sum(nil))
	document := MakeRequest(url, cookieVal)
	gradesM := map[int][]string{}
	num := 0
	isErr := true
	document.Find(".col-md-12").Each(func(i int, selection *goquery.Selection) {
		selection.ChildrenFiltered("h3").Each(func(i int, selection *goquery.Selection) {
			if selection.Text() != "" {
				isErr = false
			}
		})
	})
	document.Find("tbody").Each(func(i int, s *goquery.Selection) {
		s.ChildrenFiltered("tr").Each(func(i int, tr *goquery.Selection) {
			var grader []string = nil
			tr.ChildrenFiltered("td").Each(func(i int, td *goquery.Selection) {
				grader = append(grader, td.Text())
				//fmt.Print(td.Text())
			})
			gradesM[num] = grader
			num++
		})
	})
	if !isErr {
		var datas model.GradeModelFull
		lastRow := len(gradesM)
		//Read cumulative
		cumulative, _ := strconv.ParseFloat(gradesM[lastRow-1][1], 64)
		//Delete last row due not related to grade
		delete(gradesM, lastRow-1)
		//
		for i := 0; i < len(gradesM); i++ {
			num, _ := strconv.Atoi(gradesM[i][0])
			credit, _ := strconv.Atoi(gradesM[i][2])
			avail, _ := strconv.ParseFloat(gradesM[i][3], 64)
			quiz, _ := strconv.ParseFloat(gradesM[i][4], 64)
			assign, _ := strconv.ParseFloat(gradesM[i][5], 64)
			mid, _ := strconv.ParseFloat(gradesM[i][6], 64)
			last, _ := strconv.ParseFloat(gradesM[i][7], 64)
			gradef, _ := strconv.ParseFloat(gradesM[i][9], 64)
			datas.Data = append(datas.Data, model.GradeModelDetail{
				CourseName:   gradesM[i][1],
				GradeLetter:  gradesM[i][8],
				Num:          num,
				Credit:       credit,
				Availability: avail,
				Quiz:         quiz,
				Assignment:   assign,
				MidTerm:      mid,
				LastTerm:     last,
				GradePoint:   gradef,
			})
		}

		data := model.GradeModelFull{
			Year:       year,
			Quart:      quart,
			Cumulative: cumulative,
			Data:       datas.Data,
		}
		return data
	}
	return nil
}

//This function for make request to some url,add cookie and returned document which is ready to implemented
func MakeRequest(url string, cookieVal string) *goquery.Document {
	//url := "http://sikadu.unbaja.ac.id/mahasiswa/"
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a new cookie with the only required fields
	myCookie := &http.Cookie{
		Name:  "ci_session",
		Value: cookieVal,
	}
	// Add the cookie to request
	request.AddCookie(myCookie)
	resp, err := client.Do(request)
	if err != nil {
		println(err.Error())
	}
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	return document
}
