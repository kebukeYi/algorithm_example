package storage

import (
	"context"
	"sync"
)

func WorkerPool(
	ctx context.Context,
	jobs []int,
	workers int,
	handle func(int),
) error {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case job, ok := <-ch:
					if !ok {
						return
					}
					handle(job)

				case <-ctx.Done():
					return
				}
			}
		}()
	}

	for _, job := range jobs {
		select {
		case ch <- job:
		case <-ctx.Done():
			close(ch)
			wg.Wait()
			return ctx.Err()
		}
	}

	close(ch)
	wg.Wait()
	return nil
}
