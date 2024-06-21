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

// function deleted lessons -lesson ni updated qiladi

func (handle *Handler) CreateLesson(w http.ResponseWriter, r *http.Request) {
	lesson := model.Lesson{}
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		_, err := w.Write([]byte("error bad request  error"))
		if err != nil {
			return
		}
	}
	//err = handle.LessonRepo.CreateLesson(lesson)
	lessonJson, err := json.Marshal(&lesson)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	s, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/lesson/create", bytes.NewBuffer(lessonJson))
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte("Ok success"))
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
	}
	if response.StatusCode == http.StatusCreated || response.StatusCode == 200 {
		_, err := w.Write([]byte("ok success"))
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
}

// function updated lessons -lesson ni updated qiladi
func (handle *Handler) UpdateLesson(w http.ResponseWriter, r *http.Request) {
	lesson := model.Lesson{}
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	id := strings.TrimPrefix(r.URL.Path, "api/lesson/update")
	lessonJson, err := json.Marshal(&lesson)
	if err != nil {
		return
	}
	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/lesson/update/%s", id), bytes.NewBuffer(lessonJson))
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		_, err := w.Write([]byte("Ok success "))
		if err != nil {
			return
		}
	}

}

/* bu functionda deleteLesson da lesson ni delete qiliadi id boyicha delete qilish vazifasini oz ichiga oladi*/
func (handle *Handler) DeleteLesson(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "api/lesson/delete")

	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/lesson/delete/%s", id), nil)
	if err != nil {
		_, err := w.Write([]byte("error bad request  error"))
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte(" Ok success "))
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
	}
	if response.StatusCode == 200 {
		_, err := w.Write([]byte("Ok"))
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
}

/* filter yoki getAll lessons ga nisbatana bu method ishlatiladi  qaysi fildan query dan kelsa shu paramni search qilib topib beradi */
func (handle *Handler) GetLesson(w http.ResponseWriter, r *http.Request) {
	lessonFilter := model.Filter{}
	lessonFilter.CourseId = r.URL.Query().Get("course_id") // course id boyicha
	lessonFilter.Title = r.URL.Query().Get("title")        // title boyicha
	lessonFilter.Content = r.URL.Query().Get("content")    // content boyicha
	limit, err := strconv.Atoi(r.URL.Query().Get("limit")) // bu yerda parse qilib beradi  limit stringdan intga
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	lessonFilter.Limit = limit
	lessonFilter.Offset = offset
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return

	}
	client := http.Client{}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/lesson/get/?course_id=%s&title=%s&content=%s&limit=%d&offset=%d", &lessonFilter.CourseId, &lessonFilter.Title, &lessonFilter.Content, &lessonFilter.Limit, &lessonFilter.Offset), nil)
	response, err := client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	if response.StatusCode == 200 {
		lessons := []model.Lesson{}
		err = json.NewDecoder(response.Body).Decode(&lessons)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
		}
		json.NewEncoder(w).Encode(&lessons)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}

// getlesson by id search qiladi
func (handle *Handler) GetLessonById(w http.ResponseWriter, r *http.Request) {
	//lesson, err := handle.LessonRepo.GetById(gn.Param("id"))
	id := strings.TrimPrefix(r.URL.Path, "api/lesson/delete")

	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("localhost:8080/api/lesson/id/%s", id), nil)
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
	}
	if response.StatusCode == 200 {
		lessons := []model.Lesson{}
		err = json.NewDecoder(response.Body).Decode(&lessons)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
		} else {
			json.NewEncoder(w).Encode(&lessons)
		}
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}
