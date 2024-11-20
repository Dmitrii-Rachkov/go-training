package main

import (
	"fmt"
	"sync"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	// Добавляем wg для ожидания завершения всех горутин
	wg := &sync.WaitGroup{}
	for r := 0; r < 5; r++ {
		// Устанавливаем счётчик горутин 1
		wg.Add(1)
		go func(r int) {
			// В конце выполнения горутины даём сигнал о завершении
			defer wg.Done()

			mutex.Lock()
			state[r] = r + 1
			mutex.Unlock()
		}(r)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println(state)
}
