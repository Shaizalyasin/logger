package logger

func Info(data Interface{}){
fmt.Println("Log info: ", data)
}

func Error(data Interface{}){
fmt.Println("Log error: ", data)
}

func Warning(data Interface{}){
fmt.Println("Warning info: ", data)
}

func Fatal(data Interface{}){
	fmt.Println("Log fatal: ", data)
}
