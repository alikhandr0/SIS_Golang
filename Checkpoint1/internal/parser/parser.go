package parser 


import(
	"strings"
	"errors"
	"Checkpoint1/internal/domain/task"
)

var ErrEmptyLog = errors.New("log is empty")
var ErrInvalidFormat = errors.New("invalid log format")



func ParseLog(rawLog string) (task.Result, error){
	rawLog = strings.TrimSpace(rawLog)
	if rawLog == ""{
		return task.Result{}, ErrEmptyLog
	}

	closingBracketIndex := strings.Index(rawLog, "]")
	
	if !strings.HasPrefix(rawLog, "[") || closingBracketIndex == -1 {
		return task.Result{}, ErrInvalidFormat
	}

	level := rawLog[1:closingBracketIndex]
	message := strings.TrimSpace(rawLog[closingBracketIndex+1:])

	return task.Result{
		Level: level,
		Message: message,
		
	},nil 
}