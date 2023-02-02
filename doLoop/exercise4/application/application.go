package application

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type StudentDB struct {
	Router *mux.Router
	DB     *sql.DB
}

func (studentdb *StudentDB) Initialize(User, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", User, password, dbname)
	fmt.Println(connectionString)
	var err error
	studentdb.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	studentdb.Router = mux.NewRouter()
	studentdb.initializeRoutes()
}
func (studentdb *StudentDB) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, studentdb.Router))
}

func (studentdb *StudentDB) initializeRoutes() {
	studentdb.Router.HandleFunc("/students", studentdb.GetStudents).Methods("GET")
	studentdb.Router.HandleFunc("/student", studentdb.createStudent).Methods("POST")
	studentdb.Router.HandleFunc("/student/{id}", studentdb.GetStudent).Methods("GET")
	studentdb.Router.HandleFunc("/students/{id:[0-9]+}", studentdb.updateStudents).Methods("PUT")
	studentdb.Router.HandleFunc("/students/{id:[0-9]+}", studentdb.deleteStudent).Methods("DELETE")
}
func (studentdb *StudentDB) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := getstudents(studentdb.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, students)
}
func (studentdb *StudentDB) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fmt.Println(vars)
	id := vars["id"]
	idint, _ := strconv.Atoi(id)
	student := Student{Id: idint}
	if err := student.getstudent(studentdb.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "student not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, student)
}
func (studentdb *StudentDB) createStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := student.createstudent(studentdb.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, student)
}

func (studentdb *StudentDB) updateStudents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fmt.Println(vars)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	var student Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	student.Id = id

	if err := student.updatestudent(studentdb.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, student)
}

func (studentdb *StudentDB) deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	student := Student{Id: id}
	if err := student.deletestudent(studentdb.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "Coditationlication/json")
	w.WriteHeader(code)
	w.Write(response)
}
