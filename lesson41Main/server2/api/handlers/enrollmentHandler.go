package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"my_project/model"
	"net/http"
	"strconv"
	"strings"
)

// create enrollment - enrollment larni yaratamiz
func (handle *Handler) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	enrollment := model.Enrollment{}                   //  enrollment model ichidan strucni olib keladi
	err := json.NewDecoder(r.Body).Decode(&enrollment) // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		_, err := w.Write([]byte("error bad requeast"))
		if err != nil {
			return
		}
	}
	enrollmentJson, err := json.Marshal(&enrollment)
	if err != nil {
		_, err := w.Write([]byte("error bad requeast"))
		if err != nil {
			return
		}
		return
	}
	s, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/enrollment/create", bytes.NewBuffer(enrollmentJson))
	if err != nil {
		_, err := w.Write([]byte("error bad Internal server error"))
		if err != nil {
			return
		}
		return
	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error bad Internal server error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		_, err := w.Write([]byte("Ok success"))
		if err != nil {
			return
		}
		return
	} else {
		_, err := w.Write([]byte("error bad Internal server error"))
		if err != nil {
			return
		}
		return
	}
}

// updated qilinadi enrollments
func (handle *Handler) UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	enrollment := model.Enrollment{}                   //  enrollment model ichidan strucni olib keladi
	err := json.NewDecoder(r.Body).Decode(&enrollment) // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		_, err := w.Write([]byte("error bad request error"))
		if err != nil {
			return
		}
		return

	}
	enrollmentJson, err := json.Marshal(&enrollment)
	if err != nil {
		_, err := w.Write([]byte("error bad request error"))
		if err != nil {
			return
		}
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/enrollment/update/")

	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/enrollment/update/%s", id), bytes.NewBuffer(enrollmentJson))
	if err != nil {
		_, err := w.Write([]byte("error bad request error"))
		if err != nil {
			return
		}
		return
	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error bad request error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		_, err := w.Write([]byte("OK success"))
		if err != nil {
			return
		}
		return
	} else {
		_, err := w.Write([]byte("error internal server error error"))
		if err != nil {
			return
		}
		return
	}
}

func (handle *Handler) DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/enrollment/delete/")

	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/enrollment/delete/%s", id), nil)
	if err != nil {
		_, err := w.Write([]byte("error bad request error"))
		if err != nil {
			return
		}
		return
	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		_, err := w.Write([]byte("Ok success"))
		if err != nil {
			return
		}
		return
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
}

// filter shu qisimda boladi
func (handle *Handler) GetEnrollment(w http.ResponseWriter, r *http.Request) {
	enrollmentFilter := model.Filter{}
	//enrollmentFilter.User_id = gn.Query("id") // user id boyicha
	enrollmentFilter.UserId = r.URL.Query().Get("user_id")                 // user id si boyicha
	enrollmentFilter.CourseId = r.URL.Query().Get("course_id")             // course idsi boyicha
	enrollmentFilter.EnrollmentDate = r.URL.Query().Get("enrollment_date") // controllment date boyicha
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))                 // limit   query
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))               // offset query
	enrollmentFilter.Limit = limit                                         // limit enrollmentfiletrga tenlaymiz
	enrollmentFilter.Offset = offset
	//enrollmentFilterJson, err := json.Marshal(&enrollmentFilter)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/enrollments/get/?user_id=%s&course_id=%s&enrollment_date=%s&limit=%d&offset=%d", &enrollmentFilter.CourseId, &enrollmentFilter.UserId, &enrollmentFilter.EnrollmentDate, &enrollmentFilter.Limit, &enrollmentFilter.Offset), nil)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return

	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		enrollments := []model.Enrollment{}
		err = json.NewDecoder(response.Body).Decode(&enrollments)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		err := json.NewEncoder(w).Encode(&enrollments)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}

}

// get by id - search qiladi id boyicha

func (handle *Handler) GetEnrollmentById(w http.ResponseWriter, r *http.Request) {
	// enrollem get by id si boyicha search qiladi
	//enrollment, err := handle.EnrollmentRepo.GetById(gn.Param("id"))
	id := strings.TrimPrefix("/api/enrollment/id/", r.URL.Path)
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/enrollment/id/%s", id), nil)
	if err != nil {
		_, err := w.Write([]byte("error bad request  error"))
		if err != nil {
			return
		}
		return
	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		enrollments := []model.Enrollment{}
		err = json.NewDecoder(response.Body).Decode(&enrollments)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		err = json.NewDecoder(response.Body).Decode(&enrollments)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}

}
