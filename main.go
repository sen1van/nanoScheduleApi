package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string][]string)
	fileurl := r.URL.Path[1:] + ".yaml"
	yamlData, err := os.ReadFile(fileurl)
	if err != nil {
		fmt.Fprint(w, "Ошибка получения данных")
	}
	err = yaml.Unmarshal(yamlData, &data)

	timeNow := time.Now()
	weekday := timeNow.Weekday().String()
	_, week := timeNow.ISOWeek()
	for _, i := range data[weekday] {
		if i[0] == '0' || int(i[0])-int('0') == week%2+1 {
			fmt.Fprintf(w, "%v \t%v\n", i[3:], timetable[i[1]-48])
		}

	}
	fmt.Fprintf(w, "\n%v", timeNow.Format("2006-01-02 15:04:05"))
}

var timetable map[byte]string

func main() {
	yamlData, err := os.ReadFile("timetable.yaml")
	if err != nil {
		panic("timetable.yaml reading error... I can't work without timetable")
	}
	err = yaml.Unmarshal(yamlData, &timetable)
	if err != nil {
		panic("timetable.yaml reading error... I can't work without timetable")
	}
	fmt.Printf("%v\n", timetable['1'])

	http.HandleFunc("/", handler)
	fmt.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
