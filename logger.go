package logger

import "fmt"

func Info(message string, params interface{}) {
	fmt.Println("Log info: ", message, "params: ", params)
}

func Error(message string, params interface{}) {
	fmt.Println("Log error: ", message, "params: ", params)
}

func Warning(message string, params interface{}) {
	fmt.Println("Log Warning: ", message, "params: ", params)
}

func Fatl(message string, params interface{}) {
	fmt.Println("Log Fatal: ", message, "params: ", params)
}
