package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"my_project/models"
	"net/http"
)

func main() {
	CreateUser()
	UpdateUser()
	DeleteUser()
	GetAllUsers()
}

func CreateUser() {
	user := models.User{}
	fmt.Print("1-Name:")
	_, err := fmt.Scanf("%s", &user.Name)
	if err != nil {
		return
	}
	fmt.Print("2-Age:\t")
	_, err = fmt.Scanf("%d", &user.Age)
	if err != nil {
		return
	}
	fmt.Print("3-email:\t")
	_, err = fmt.Scanf("%s", &user.Email)
	if err != nil {
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return
	}
	resp, err := http.Post("http://localhost:8080/user/create", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		fmt.Println("Failed :", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode == 201 || resp.StatusCode == 200 {
		fmt.Println("success")
	}

}

func UpdateUser() {
	user := models.User{}

	fmt.Print("1-Name: ")
	_, err := fmt.Scanf("%s", &user.Name)
	if err != nil {
		return
	}

	fmt.Print("2-Age: ")
	_, err = fmt.Scanf("%d", &user.Age)
	if err != nil {
		return
	}

	fmt.Print("3-Email: ")
	_, err = fmt.Scanf("%s", &user.Email)
	if err != nil {
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return
	}
	fmt.Print("Id: ")
	var id string
	_, err = fmt.Scanf("%s", &id)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/user/update/%d", id), bytes.NewBuffer(userJson))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Failed", err)
		}
	}()

	if resp.StatusCode == http.StatusOK {
		fmt.Println(" successfully")
	} else {
		fmt.Printf("Failed  Status code: %d\n", resp.StatusCode)
	}
}

func DeleteUser() {
	var id int

	fmt.Print("Id: ")
	_, err := fmt.Scanf("%s", &id)
	if err != nil {
		fmt.Println("id -> uuid")
		return
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/user/delete/%d", id), nil)
	if err != nil {
		fmt.Println("Failed :", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed :", err)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Failed ", err)
		}
	}()

	if resp.StatusCode == http.StatusOK {
		fmt.Println(" successfully")
	} else {
		fmt.Printf("Failed Status code: %d\n", resp.StatusCode)
	}
}

func GetAllUsers() {
	url := "http://localhost:8080/users/get"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed ", err)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Failed ", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed Status code: %d\n", resp.StatusCode)
		return
	}

	users := []models.User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Println("Failed ", err)
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s\n", user.Id, user.Name, user.Age, user.Email)
	}
}
