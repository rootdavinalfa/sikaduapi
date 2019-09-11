/*
 * Copyright (c) 2019. dvnlabs.ml
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package Student

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"unbajaUAPI/model"
)

func GetStudentInfo(cookieVal string) interface{} {
	url := "http://sikadu.unbaja.ac.id/mahasiswa/"
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
	//Schedule parsing
	//var schedules []interface{}
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
	document.Find("tbody").Each(func(i int, s *goquery.Selection) {
		s.ChildrenFiltered("tr").Each(func(i int, tr *goquery.Selection) {
			//println("")
			var scheduler []string = nil
			tr.ChildrenFiltered("td").Each(func(i int, td *goquery.Selection) {
				//print(td.Text())
				scheduler = append(scheduler, td.Text())
			})
			//schedules = append(schedules,scheduler)
			scheduleM[num] = scheduler
			num++
		})
	})
	if !isErr {
		var datas model.ScheduleFull
		for i := 0; i < len(scheduleM); i++ {
			/*for a := 0;a< len(scheduleM[i]);a++{
				println(scheduleM[i][a])
			}*/
			smt, _ := strconv.Atoi(scheduleM[i][1])
			datas.Data = append(datas.Data, model.ScheduleStudentPeriode{
				CourseName: scheduleM[i][0],
				Class:      scheduleM[i][2],
				Room:       scheduleM[i][3],
				Lecturer:   scheduleM[i][4],
				Days:       scheduleM[i][5],
				Semester:   smt,
			})
		}
		//test := MustMarshal(schedules)
		//print(string(test))
		/*for k, v := range scheduleM {
			fmt.Printf("key[%s] value[%s]\n", k, v)
		}*/

		data := model.ScheduleFull{
			Year:  year,
			Quart: quart,
			Data:  datas.Data,
		}
		return data
	}
	return nil
}
