package log

import (
	"go.uber.org/zap"
	"testing"
)

var log0, _ = zap.NewProduction()

//var log0, _ = zap.NewDevelopment()

func Test_WithName(t *testing.T) {
	defer Sync() // used for record logger printer
	log0.Error("sd")
}

func Test_V(t *testing.T) {
	defer Sync() // used for record logger printer
	log0.Info("hello ")

	Debugw("Hello world!", "key", "value")
	Infow("Hello world!", "key", "value")
	Warnw("Hello world!", "key", "value")

}

func BenchmarkLog1(b *testing.B) {
	defer Sync()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1; i++ {
			Infow("Hello world!", "key", "value")
			//Warnw("Hello world!", "key", "value")
		}
	}
}
