package fan_in

import "sync"

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	// TODO: ваш код

	chout := make(chan int)

	var wg sync.WaitGroup

	output := func(ch <-chan int, id int) {
		defer wg.Done()
		for n := range ch {
			chout <- n
		}
	}

	wg.Add(len(channels))
	for id, ch := range channels {
		go output(ch, id)
	}

	go func() {
		wg.Wait()
		close(chout)
	}()
	return chout
}
