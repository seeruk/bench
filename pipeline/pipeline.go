package pipeline

import (
	"context"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

func Source(ctx context.Context, n int) (context.Context, chan int) {
	log.Println("Source: Started")

	// Single-threaded source.
	is := make(chan int, n)

	go func() {
		for i := 0; i < n; i++ {
			n := randBetween(0, math.MaxInt16)
			log.Printf("Made %d: %d", i, n)
			is <- n
			time.Sleep(2 * time.Millisecond)
		}

		close(is)
	}()

	return ctx, is
}

func SquareTransformBatchLong(ctx context.Context, in chan int) (context.Context, chan int) {
	log.Println("Started batch function")

	source := make(chan int, 256)
	out := make(chan int)
	ctxc := make(chan context.Context, 1)

	go func() {
		// Wait for our tracing "local" context before doing any work.
		lctx := <-ctxc

		// This could be turned into a loop and slice.
		_, tform1 := SquareTransformLong(lctx, 1, source)
		_, tform2 := SquareTransformLong(lctx, 2, source)
		_, tform3 := SquareTransformLong(lctx, 3, source)
		_, tform4 := SquareTransformLong(lctx, 4, source)

		// Merge results, before sending them on their way.
		_, merged := MergeInts(ctx, tform1, tform2, tform3, tform4)

		for m := range merged {
			out <- m
		}

		close(out)
	}()

	go func() {
		first := true

		for {
			select {
			case <-ctx.Done():
				return
			case i, ok := <-in:
				if !ok {
					close(source)
					return
				}

				if first {
					// Make a new context, it'd come from the span really.
					log.Println("Starting the real work now.")

					ctxc <- ctx
					first = false
				}

				source <- i
			}
		}
	}()

	return ctx, out
}

func SquareTransformLong(ctx context.Context, n int, in chan int) (context.Context, chan int) {
	log.Printf("SquareTransformLong %d: Started\n", n)

	out := make(chan int)

	var c int

	go func() {
		defer close(out)

		for {
			c++

			select {
			case <-ctx.Done():
				return
			case i, ok := <-in:
				if !ok {
					return
				}

				log.Printf("SquareTransformLong %d: %d: %d\n", n, c, i)

				time.Sleep(time.Duration(randBetween(500, 1000)) * time.Nanosecond)

				out <- i * i
			}
		}
	}()

	return ctx, out
}

func MergeInts(ctx context.Context, iss ...<-chan int) (context.Context, <-chan int) {
	log.Println("MergeInts: Started")

	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(iss))
	for _, is := range iss {
		go func(is <-chan int) {
			for i := range is {
				out <- i
			}
			wg.Done()
		}(is)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return ctx, out
}

func SinkLong(ctx context.Context, in <-chan int) {
	log.Println("SinkLong: Started")

	var c int

	for {
		c++

		select {
		case <-ctx.Done():
			return
		case i, ok := <-in:
			if !ok {
				return
			}

			if c == 1 {
				// Can start a span here for this process?
				log.Println("First end result:")
			}

			log.Printf("End result %d: %d\n", c, i)
		}
	}
}

func randBetween(min, max int) int {
	return rand.Intn(max-min) + min
}
