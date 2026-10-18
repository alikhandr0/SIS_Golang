package manager 

import (
	"testing"
	"Checkpoint1/internal/domain/task"
	"Checkpoint1/internal/storage"

	
)

func TestProcessTaskLifecycle(t *testing.T) {


	store := storage.NewMemoryStorage()
	m := NewTaskManager(store)

	id, err := m.CreateTask("[INFO] Test log process")
	if err != nil {

		t.Fatalf("не удалось создать задачу: %v", err)

	}

	

	err = m.ProcessTask(id)
	if err != nil {
		t.Errorf("ожидали успешную обработку, но получили ошибку: %v", err)
	}

	savedTask, _ := store.GetByID(id)
	if savedTask.Status != task.Completed {
		t.Errorf("ожидали статус %s, получили %s", task.Completed, savedTask.Status)
	}

	err = m.ProcessTask(id)
	if err != ErrInvalidState {
		t.Errorf("ожидали защиту от повторного запуска (%v), но получили: %v", ErrInvalidState, err)
	}
}