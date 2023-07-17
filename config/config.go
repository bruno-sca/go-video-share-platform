package config

var (
	logger *Logger
)


func Init() error {
	// Add necessary initialization logic here
	
	return nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}