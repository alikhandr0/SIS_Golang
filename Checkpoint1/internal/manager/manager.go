package manager

import(
	"errors"
	"time"
	"Checkpoint1/internal/domain/task"
	"Checkpoint1/internal/parser"
	
)


var ErrInvalidState = errors.New("task cannot be processed in its current state ")

type Storage interface{
	Save(t *task.Task) error
	GetByID(id int) ( *task.Task, error)
}

type TaskManager struct{
	store Storage
}

func NewTaskManager(store Storage) *TaskManager{
	return &TaskManager{
		store: store,
	}
}

func (m *TaskManager ) CreateTask(rawLog string ) (int, error){
	t := &task.Task{
		RawLog: rawLog,
		Status: task.Pending,
		CreatedAt: time.Now(),
	}

	err := m.store.Save(t)
	if err != nil{
		return 0, err
	}

	return t.ID, nil
}

func (m *TaskManager) ProcessTask(id int) error {
	t, err := m.store.GetByID(id)
	if err != nil{
		return err
	}

	if t.Status != task.Pending{
		return ErrInvalidState 
	}

	t.Status = task.Processing

	res, err := parser.ParseLog(t.RawLog)

	if err != nil{
		return err
	}

	t.Result = res
	t.Status = task.Completed

	return nil 
}

