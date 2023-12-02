package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
)

type Question struct {
	ID     string      `json:"id"`
	Title  string      `json:"title"`
	Type   string      `json:"type"`
	Answer interface{} `json:"answer"`
	Option []string    `json:"option"`
}

type Info struct {
	WorkID     string `json:"work_id"`
	ID         string `json:"id"`
	WorkName   string `json:"work_name"`
	WorkStatus string `json:"work_status"`
	WorkURL    string `json:"work_url"`
	CourseName string `json:"course_name"`
	IsRedo     string `json:"isRedo"`
	Score      string `json:"score"`
}

type Data struct {
	Questions []Question `json:"27835863"`
	Info      Info       `json:"info"`
}

func ReadFileToDB() {
	dir := "answer"

	files, err := getFilesInDir(dir)
	if err != nil {
		fmt.Println("无法读取目录:", err)
		return
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	fmt.Println(fileNames)
	filename := "answer/" + fileNames[0]
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "module/answer_save/db/answer.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS questions (id TEXT, title TEXT, type TEXT, answer TEXT, option TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO questions(id, title, type, answer, option) values(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	for _, question := range data.Questions {
		_, err = stmt.Exec(question.ID, question.Title, question.Type, question.Answer, question.Option)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getFilesInDir(dir string) ([]os.FileInfo, error) {
	var files []os.FileInfo

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, info)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
