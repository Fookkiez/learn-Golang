package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	ID         int     `json: "id"`
	Name       string  `json: "name"`
	Price      float64 `json: "price"`
	Instructor string  `json: "instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
		{
			"id" : 1,
			"name" : "python",
			"price" : 2300,
			"instructor" : "FOokkie541"
		},
		{
			"id" : 2,
			"name" : "C#",
			"price" : 13333,
			"instructor" : "FOokkie133"
		},
		{
			"id" : 3,
			"name" : "GO",
			"price" : 313,
			"instructor" : "FOokkie231"
		}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	CourseJSON, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(CourseJSON)
	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(" r.Body", r.Body)
		err = json.Unmarshal(Bodybyte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.ID = getnextID()

		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func getnextID() int {
	highestID := -1
	for _, course := range CourseList {
		if highestID < course.ID {
			highestID = course.ID
		}
	}
	return highestID + 1
}

func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.ID == ID {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("no coures id : %d", ID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		courserJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courserJSON)
	case http.MethodPut:
		var updateCourse Course
		fmt.Println(updateCourse)
		byteBody, err := ioutil.ReadAll(r.Body)
		fmt.Println("Bodybyte", byteBody, "\n")
		fmt.Println("err", r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(byteBody, &updateCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updateCourse.ID != ID {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(jso`[{"Data"}]`)))
			return
		}
		course = &updateCourse
		CourseList[listItemIndex] = *course
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/course/", courseHandler)
	http.HandleFunc("/course", coursesHandler)
	http.ListenAndServe(":5000", nil)
}
