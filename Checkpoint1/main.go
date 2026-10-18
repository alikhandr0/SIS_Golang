package main

import (
	"fmt"
	"Checkpoint1/internal/manager"
    "Checkpoint1/internal/storage"
	)

func main(){
	store := storage.NewMemoryStorage()

	taskManager := manager.NewTaskManager(store)

	fmt.Println("=== тест1: успешный парсинг ===")

	id1, err := taskManager.CreateTask("[ERROR] database connection lost")

	if err != nil{
		fmt.Println("ошибка создания:", err)
		return
	}

	fmt.Printf("создана задача с ID: %d\n", id1)

	err= taskManager.ProcessTask(id1)
	if err!= nil{
		fmt.Println("ошибка оброботки:", err)
	}else {
		fmt.Println("задача успешно робработана")
		t, _ := store.GetByID(id1)
		fmt.Printf("статус: %s\n", t.Status)
		fmt.Printf("уровень: %s | сообщение: %s\n\n", t.Result.Level, t.Result.Message)

		
	}

	fmt.Println("=== тест2: ошибка валидации ===")

	id2, _ := taskManager.CreateTask("INFO server started successtully ")
	fmt.Printf("создана задача с ID: %d\n", id2) 

	err = taskManager.ProcessTask(id2)
	if err != nil{
		fmt.Printf("ожидаемая ошибка обработки: %v\n", err)
	}

}