package task

import (
	"time"
)

//status of task life sicle
const(
	Pending string = "PENDING"
	Processing string = "PROCESSING"
	Completed string = "COMPLETED"
)

type Task struct{
	ID int 	
	Status string 
	RawLog string		// 
	CreatedAt time.Time	// time
	Result Result 
	
}


type Result struct{
	Level string		// level of criticality: ERROr, INFO
	Message string		// cleared log texT
	
}




