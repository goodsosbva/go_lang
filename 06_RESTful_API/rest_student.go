package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var (
	students   map[int]Student
	lastId     int
	studentsMu sync.RWMutex
)

func MakeWebHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/students", GetStudentListHandler).Methods(http.MethodGet)
	router.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods(http.MethodGet)
	router.HandleFunc("/students", PostStudentHandler).Methods(http.MethodPost)
	router.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods(http.MethodDelete)

	studentsMu.Lock()
	students = map[int]Student{
		1: {Id: 1, Name: "aaa", Age: 16, Score: 87},
		2: {Id: 2, Name: "bbb", Age: 18, Score: 98},
	}
	lastId = 2
	studentsMu.Unlock()

	return router
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Students) Less(i, j int) bool { return s[i].Id < s[j].Id }

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	studentsMu.RLock()
	list := make(Students, 0, len(students))
	for _, student := range students {
		list = append(list, student)
	}
	studentsMu.RUnlock()

	sort.Sort(list)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid student id", http.StatusBadRequest)
		return
	}

	studentsMu.RLock()
	student, ok := students[id]
	studentsMu.RUnlock()
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&student); err != nil {
		http.Error(w, "invalid student JSON", http.StatusBadRequest)
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF || strings.TrimSpace(student.Name) == "" {
		http.Error(w, "invalid student JSON", http.StatusBadRequest)
		return
	}

	studentsMu.Lock()
	lastId++
	student.Id = lastId
	students[lastId] = student
	studentsMu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid student id", http.StatusBadRequest)
		return
	}

	studentsMu.Lock()
	_, ok := students[id]
	if ok {
		delete(students, id)
	}
	studentsMu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", MakeWebHandler()))
}
