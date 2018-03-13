package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/seeruk/bench/pipeline"
)

func main() {
	rand.Seed(time.Now().Unix())

	ctx, cfn := context.WithTimeout(context.Background(), 30*time.Second)
	defer cfn()

	// Produce some values.
	_, source := pipeline.Source(ctx, 100)

	// Process them all in one place.
	_, merged := pipeline.SquareTransformBatchLong(ctx, source)

	// Consume the values at the end.
	pipeline.SinkLong(ctx, merged)
}
