package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetStudentInfo(t *testing.T) {
	id := "学生1"
	stu := GetStudentById(id)
	if len(stu.Name) == 0 {
		t.Failed()
	} else {
		fmt.Println("+v\n", stu)
	}
}

func TestGetName(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:2345/get_name?student_id=学生2")
	if err != nil {
		fmt.Println("Error:", err)
		t.Fail()
	} else {
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			t.Fail()
		} else {
			fmt.Println(string(bytes))
		}
	}
}

func TestGetAge(t *testing.T) {
	//url.Values{}  map[string][]string
	queryParams := url.Values{"student_id": []string{"学生2"}}
	resp, err := http.PostForm("http://127.0.0.1:2345/get_age", queryParams) // 修正为传递两个参数
	if err != nil {
		fmt.Println("Error:", err)
		t.Fail()
	} else {
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			t.Fail()
		} else {
			fmt.Println(string(bytes))
		}
	}
}

func TestGetHeight(t *testing.T) {
	client := http.Client{}
	reader := strings.NewReader(`{"student_id":"学生2"}`)
	request, err := http.NewRequest("POST", "http://127.0.0.1:2345/get_height", reader)
	if err != nil {
		fmt.Println("Error:", err)
		t.Fail()
	}

	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
		t.Fail()
	} else {
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			t.Fail()
		} else {
			//fmt.Println(string(bytes))
			var stu Student
			err = json.Unmarshal(bytes, &stu)
			if err != nil {
				fmt.Println("Error:", err)
				t.Fail()
			} else {
				fmt.Printf("%+v\n", stu)
			}
		}
	}
}
