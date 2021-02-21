package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Student struct
type Student struct {
	ID          string `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Rollno      string `json:"rollno"`
	Dob         string `json:"dob"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Marks       *Marks `json:"marks"`
}

//Marks struct
type Marks struct {
	Maths         int `json:"maths"`
	Science       int `json:"science"`
	Socialscience int `json:"socialscience"`
	English       int `json:"english"`
	Hindi         int `json:"hindi"`
}

var students []Student

func getstudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Student{})
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Student Student
	_ = json.NewDecoder(r.Body).Decode(&Student)
	students = append(students, Student)
	json.NewEncoder(w).Encode(Student)
}

//Unassigned values will remain as empty strings(default for a string)
func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var Student Student
			_ = json.NewDecoder(r.Body).Decode(&Student)
			Student.ID = params["id"]
			students = append(students, Student)
			json.NewEncoder(w).Encode(Student)
			return
		}
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func main() {
	r := mux.NewRouter()

	//Random data
	students = append(students, Student{ID: "1", Firstname: "Aman", Lastname: "Tirumala", Rollno: "22", Dob: "23-2-1997", Email: "amantirumala@gmail.com", Phonenumber: "6303989898", Marks: &Marks{Maths: 98, Science: 65, Socialscience: 76, English: 84, Hindi: 98}})
	students = append(students, Student{ID: "2", Firstname: "Lokseh", Lastname: "Gopalli", Rollno: "19", Dob: "23-2-1997", Email: "lokeshgopalli@gmail.com", Phonenumber: "6303999999", Marks: &Marks{Maths: 96, Science: 75, Socialscience: 66, English: 77, Hindi: 95}})

	r.HandleFunc("/students", getstudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
