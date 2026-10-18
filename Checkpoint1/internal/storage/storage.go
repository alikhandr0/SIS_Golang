package storage

import(
	"errors"
	"Checkpoint1/internal/domain/task"
)

var ErrTaskNotFound = errors.New("task not found")

type MemoryStorage struct{
	tasks map[int]*task.Task
	lastID int		//counter for generating unique  ID
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[int]*task.Task),
		lastID: 0, 
		
	}
}

// save take the task, give to him new id and save on map
func (s *MemoryStorage) Save(t *task.Task) error{
	s.lastID = s.lastID + 1
	t.ID = s.lastID
	s.tasks[t.ID] = t
	
	return nil
	
}

//GetByID locking for task by ID
func (s *MemoryStorage) GetByID(id int) (*task.Task, error){
	v, ok := s.tasks[id]

	if ok == false{
		return nil, ErrTaskNotFound
	}
	
	return v, nil
	
}


