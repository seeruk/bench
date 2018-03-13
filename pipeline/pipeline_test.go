package pipeline

import (
	"context"
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkPipeline(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	rand.Seed(time.Now().Unix())

	ctx, cfn := context.WithTimeout(context.Background(), time.Second)
	defer cfn()

	for i := 0; i < b.N; i++ {
		_, source := Source(ctx, 100)

		// Process them
		_, tform1 := SquareTransformLong(ctx, 1, source)
		_, tform2 := SquareTransformLong(ctx, 2, source)
		_, tform3 := SquareTransformLong(ctx, 3, source)
		_, tform4 := SquareTransformLong(ctx, 4, source)

		_, merged := MergeInts(ctx, tform1, tform2, tform3, tform4)

		SinkLong(ctx, merged)
	}
}

func BenchmarkPipelineBatch(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	rand.Seed(time.Now().Unix())

	ctx, cfn := context.WithTimeout(context.Background(), time.Second)
	defer cfn()

	for i := 0; i < b.N; i++ {
		_, source := Source(ctx, 100)
		_, merged := SquareTransformBatchLong(ctx, source)

		SinkLong(ctx, merged)
	}
}
