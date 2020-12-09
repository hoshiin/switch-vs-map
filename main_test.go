package main

import (
	"math/rand"
	"os"
	"testing"
)

var ascInputs5000 Animals
var randInputs5000 Animals
var ascInputs10000 Animals
var randInputs10000 Animals
var ascInputs50000 Animals
var randInputs50000 Animals
var animals = Animals{NewDog(), NewCat(), NewBird(), NewBear(), NewRabbit()}

func TestMain(m *testing.M) {
	ascInputs5000 = makeAscInputs(5000)
	randInputs5000 = makeRandInputs(5000)
	ascInputs10000 = makeAscInputs(10000)
	randInputs10000 = makeRandInputs(10000)
	ascInputs50000 = makeAscInputs(50000)
	randInputs50000 = makeRandInputs(50000)

	os.Exit(m.Run())
}

func makeRandInputs(num int) Animals {
	randInputs := make(Animals, num)
	for i := 0; i < num; i++ {
		randInputs[i] = animals[rand.Int()%5]
	}
	return randInputs
}

func makeAscInputs(num int) Animals {
	ascInputs := make(Animals, num)
	for i := 0; i < num; i++ {
		ascInputs[i] = animals[i%5]
	}
	return ascInputs
}

func Benchmark_5000_actionSwitch_5000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs5000.ActionSwitch()
	}
}

func Benchmark_5000_actionSwitchType_5000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs5000.ActionSwitchType()
	}
}

func Benchmark_5000_actionMap_5000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs5000.ActionMap()
	}
}

func Benchmark_5000_actionSwitch_5000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs5000.ActionSwitch()
	}
}

func Benchmark_5000_actionSwitchType_5000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs5000.ActionSwitchType()
	}
}

func Benchmark_5000_actionMap_5000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs5000.ActionMap()
	}
}

func Benchmark_10000_actionSwitch_10000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs10000.ActionSwitch()
	}
}

func Benchmark_10000_actionSwitchType_10000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs10000.ActionSwitchType()
	}
}

func Benchmark_10000_actionMap_10000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs10000.ActionMap()
	}
}

func Benchmark_10000_actionSwitch_10000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs10000.ActionSwitch()
	}
}

func Benchmark_10000_actionSwitchType_10000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs10000.ActionSwitchType()
	}
}

func Benchmark_10000_actionMap_10000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs10000.ActionMap()
	}
}

func Benchmark_50000_actionSwitch_50000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs50000.ActionSwitch()
	}
}

func Benchmark_50000_actionSwitchType_50000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs50000.ActionSwitchType()
	}
}

func Benchmark_50000_actionMap_50000_asc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ascInputs50000.ActionMap()
	}
}

func Benchmark_50000_actionSwitch_50000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs50000.ActionSwitch()
	}
}

func Benchmark_50000_actionSwitchType_50000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs50000.ActionSwitchType()
	}
}

func Benchmark_50000_actionMap_50000_random(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randInputs50000.ActionMap()
	}
}
