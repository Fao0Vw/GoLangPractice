package main

import (
	"log"
	"os"
)

func main() {
	// log.Println("hello")
	// log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	// log.Println("hello")
	// if len(os.Args) == 1 {
	// 	fmt.Println("Program required one of more arguments!")
	// 	return
	// }
	// switch os.Args[1] {
	// case "panic":
	// 	log.Panic("Panic")

	// case "fatal":
	// 	log.Fatalln("Fatal")
	// default:
	// 	fmt.Println("All is good")
	// }
	// logFile, _ := os.Create("info.log")
	// log.SetOutput(logFile)
	// log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	// log.Println("Event.")
	// logFile.Close()
	// time.Sleep(3 * time.Second)
	// os.Remove("info.log")
	flags := log.LstdFlags | log.Lshortfile
	// logInfo := log.New(os.Stdout, "INFO:\t", log.LstdFlags)
	// logWarn := log.New(os.Stdout, "WARN:\t", flags)
	// logError := log.New(os.Stderr, "ERROR:\t", flags)
	fileInfo, _ := os.OpenFile("log_info_file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fileWarn, _ := os.OpenFile("log_warn_file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fileError, _ := os.OpenFile("log_error_file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logInfo := log.New(fileInfo, "INFO:\t", log.LstdFlags)
	logWarn := log.New(fileWarn, "WARN:\t", flags)
	logError := log.New(fileError, "ERROR:\t", flags)
	logInfo.Println("information")
	logWarn.Println("Warning")
	logError.Println("ERROR!")
	defer func() {
		fileInfo.Close()
		fileWarn.Close()
		fileError.Close()
	}()
}
