package logit

import(	
	"log"
	"encoding/json"
	"os"
	"bytes"
//	"fmt"
//	"time"
)

type Logger struct{
	LogFile string
	LogLevel string		
}

type Config struct{
	settings map[string]string
	filename string
}

var logger *Logger
var config = new(Config)


func (*Logger) Initialize(config *Config){
	logger.LogFile = config.settings["LogFile"]
	logger.SetOutput()
//	logger.SetPrefix()
}

func NewLogger(pathToConfigFile string)(*Logger){	
	if pathToConfigFile == "" {
		pathToConfigFile = "./logConfig.json"
	}

	f, err := os.Open(pathToConfigFile)
	checkError(err)
	
	b := new(bytes.Buffer)
	_, err = b.ReadFrom(f)
	checkError(err)

//	fmt.Printf("Bytes %v", string(b.Bytes()))
	err = json.Unmarshal(b.Bytes(), &config.settings)
	checkError(err)	

	logger = new(Logger)
	logger.Initialize(config)	

	//fmt.Printf("Print Config %v \n", config.settings)
	return logger
}
	
func (l *Logger) SetOutput(){
	f, err := os.OpenFile(config.settings["logFile"], os.O_APPEND|os.O_WRONLY, 0666)
 	checkError(err)

	log.SetOutput(f)
}


func (*Logger) Debug(msg string){	
	log.Printf("[Debug] %v", msg)
}

func (*Logger) Info(msg string){
	log.Printf("[Info] %v", msg)
}

func (*Logger) Error(msg string){	
	log.Printf("[Error] %v", msg)
}

// func main(){
// 	newLogger := NewLogger()
// 	fmt.Println(newLogger)

// 	for {
// 		time.Sleep(1 * time.Second)
// 		newLogger.Error("Hello World")
// 	}
// }


func checkError(err error){
	if err != nil{
		log.Printf("[ERROR] %v \n" , err)
		os.Exit(1)
	}
}

