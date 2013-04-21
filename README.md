Logit
=====

Logit is a very basic and simple logging library which helps users log to a file.


Installation
=============

        go get github.com/supersid/logit

Usage
=====

Assuming that you want a log file called `development.log` in your current directory you would need to
follow the next few steps.

        touch development.log (This will be automated in subsequent releases but for now lets build this)
        
        touch logConfig.json

Now load the logConfig.json and update it with the following code
        
        {
           "logFile":"siddharth.log",
	         "logLevel":"debug"
        }

Note the keys need to be strings.


In your application 

      package main
```go
      import(
        "github.com/supersid/logit"
      	"fmt"
      	"time"
      )
      
      func main(){
      	fmt.Println("Welcome to Logit test application")
      	logger := logit.NewLogger("./logConfig.json")
      
      	for{
      		time.Sleep(2 * time.Second)
      		logger.Debug("Debugging my Code")
      	}
      }
```
