package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/* create course method dida gin contextni oladi  bu yerda request ni qabul qilib create methoddga berib yuboradi u method table yozadi*/
func (handle *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	course := model.Course{}                       //  enrollment model ichidan strucni olib keladi
	err := json.NewDecoder(r.Body).Decode(&course) // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		_, err := w.Write([]byte("error bad request"))
		if err != nil {
			return
		}
	}
	courseJson, err := json.Marshal(&course)
	if err != nil {
		_, err := w.Write([]byte("error bad request"))
		if err != nil {
			return
		}
		return
	}
	s, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/course/create", bytes.NewBuffer(courseJson))
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

// updated course bu updated qiliadi
func (handle *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	course := model.Course{}                       //  enrollment model ichidan strucni olib keladi
	err := json.NewDecoder(r.Body).Decode(&course) // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		_, err := w.Write([]byte("error bad request error"))
		if err != nil {
			return
		}
		return

	}
	courseJson, err := json.Marshal(&course)
	if err != nil {
		_, err := w.Write([]byte("error bad request error"))
		if err != nil {
			return
		}
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/course/update/")

	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/course/update/%s", id), bytes.NewBuffer(courseJson))
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

// delete course qilinadi
func (handle *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/course/delete/")

	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/course/delete/%s", id), nil)
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

// getcourse filter yoki getAt course
func (handle *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {
	courseFilter := model.Filter{} /// course course filter modeldan chaqiriladi
	courseFilter.Title = r.URL.Query().Get("title")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))   // limitni  int ga parse qiliadi
	offset, err := strconv.Atoi(r.URL.Query().Get("offset")) // offsetni int ga parse qiliadi
	courseFilter.Limit = limit                               // limit  bunda filterda boladi
	courseFilter.Offset = offset                             // offset bunda filterda boladi
	courseFilter.Description = r.URL.Query().Get("description")
	//courseFilterJson, err := json.Marshal(&courseFilter)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/course/get/?title=%s&description=%s&limit=%d&offset=%d", &courseFilter.Title, &courseFilter.Description, &courseFilter.Limit, &courseFilter.Offset), nil)
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
		courses := []model.Course{}
		err = json.NewDecoder(response.Body).Decode(&courses)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		err := json.NewEncoder(w).Encode(&courses)
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

// course id si orqali qidirib lessons larni olib keladi
func (handle *Handler) GetLessonsByCourseId(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/course/lessons/")

	//}
	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/course/lessons/%s", id), nil)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
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
		lessons := []model.Lesson{}
		err = json.NewDecoder(response.Body).Decode(&lessons)
		json.NewEncoder(w).Encode(&lessons)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}

func (handle *Handler) GetUserWithEnrollmentByCourseId(w http.ResponseWriter, r *http.Request) {
	// course id,user lar slice ,err qaytaradi  course da enrollment bilen inner join
	//qilinidai user_id topiladi keyin user_id boyicha user ni topadi
	//courseId, users, err := handle.CourseRepo.GetUserByCourseIdWithEnrollment(gn.Param("id"))
	courseId := strings.TrimPrefix(r.URL.Path, "/api/course/enrollments/")

	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/course/enrollments/%s", courseId), nil)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		users := []model.User{}
		err = json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		err = json.NewDecoder(response.Body).Decode(&users)
		json.NewEncoder(w).Encode(&users)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return

	}

}

// course larni idsi boyicha qidiradi
func (handle *Handler) GetCourseById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/course/id/")

	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/course/id/%s", id), nil)
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
		courses := []model.Course{}
		err = json.NewDecoder(response.Body).Decode(&courses)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		json.NewEncoder(w).Encode(&courses)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
}

// show popularni courselarni qidiradi
func (handle *Handler) ShowPopularCourse(w http.ResponseWriter, r *http.Request) {
	//startTime := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	//endTime := time.Date(2030, 12, 31, 23, 59, 59, 999, time.UTC)

	time1 := r.URL.Query().Get("start_time")
	time2 := r.URL.Query().Get("end_time")

	startTime1 := strings.Split(time1, "-")
	startYear, err := strconv.Atoi(startTime1[0])
	startMonth, err := strconv.Atoi(startTime1[1])
	startDay, err := strconv.Atoi(startTime1[2])
	if err != nil {
		panic(err)
	}
	endTime1 := strings.Split(time2, "-")

	endYear, err := strconv.Atoi(endTime1[0])
	endMonth, err := strconv.Atoi(endTime1[1])
	endDay, err := strconv.Atoi(endTime1[2])
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

	startTime := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(endYear, time.Month(endMonth), endDay, 00, 00, 00, 0, time.UTC)

	//courses, err := handle.CourseRepo.GetPopularyCourse(startTime, endTime)

	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/?start_time=%s&end_time=%s", startTime, endTime), nil)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		courses := []model.Course{}
		err = json.NewDecoder(response.Body).Decode(&courses)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		json.NewEncoder(w).Encode(&courses)

	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}
