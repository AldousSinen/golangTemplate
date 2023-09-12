package utils

import (
	"log"
	"os"
	"strings"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Separator     *log.Logger
)

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func SystemLogger(class, folder, filename, process string, logContent map[string]interface{}) {
	currentTime := time.Now()
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/"+filename+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")

	InfoLogger.Println(class + ": - - - -  " + strings.ToUpper(process) + " - - - -")
	for k, v := range logContent {
		InfoLogger.Printf("%v: %v: %v \n", class, k, v)
	}

	defer file.Close()
}
