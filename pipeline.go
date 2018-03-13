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

	//// Alternatively, process them in chunks.
	//_, tform1 := pipeline.SquareTransformLong(ctx, 1, source)
	//_, tform2 := pipeline.SquareTransformLong(ctx, 2, source)
	//_, tform3 := pipeline.SquareTransformLong(ctx, 3, source)
	//_, tform4 := pipeline.SquareTransformLong(ctx, 4, source)
	//
	//_, merged := pipeline.MergeInts(ctx, tform1, tform2, tform3, tform4)

	// Consume the values at the end.
	pipeline.SinkLong(ctx, merged)
}
