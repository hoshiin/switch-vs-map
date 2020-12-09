# Go Switch vs Map for handling interface Benchmarks
This benchmark measures the performance of handling interface with switch calling a function vs using a map of functions.
```golang
switch animal := a.(type) {
		case Dog:
			animal.Bark()
		case Cat:
			animal.Meow()
		case Bird:
			animal.Fly()
		case Bear:
			animal.Roar()
		case Rabbit:
			animal.Jump()
		}
```
vs
```golang
switch a.Type() {
		case AnimalTypeDog:
			a.(Dog).Bark()
		case AnimalTypeCat:
			a.(Cat).Meow()
		case AnimalTypeBird:
			a.(Bird).Fly()
		case AnimalTypeBear:
			a.(Bear).Roar()
		case AnimalTypeRabbit:
			a.(Rabbit).Jump()
		}
```
vs
```golang
handler := actionHandler[a.Type()]
```

# Running Benchmarks
go test -bench . 

# Results
This following results were produced from a CPU 2.2 GHz Intel Core i7 running `go1.14.7 darwin/amd64` on macOS Catalina.
```
Benchmark_5000_actionSwitch_5000_asc-4                    103582             10400 ns/op
Benchmark_5000_actionSwitchType_5000_asc-4                 46832             24090 ns/op
Benchmark_5000_actionMap_5000_asc-4                        22680             52334 ns/op
Benchmark_5000_actionSwitch_5000_random-4                  25519             50577 ns/op
Benchmark_5000_actionSwitchType_5000_random-4              18538             73040 ns/op
Benchmark_5000_actionMap_5000_random-4                     10000            101688 ns/op
Benchmark_10000_actionSwitch_10000_asc-4                   61953             18590 ns/op
Benchmark_10000_actionSwitchType_10000_asc-4               25437             48128 ns/op
Benchmark_10000_actionMap_10000_asc-4                      10000            103380 ns/op
Benchmark_10000_actionSwitch_10000_random-4                10000            144514 ns/op
Benchmark_10000_actionSwitchType_10000_random-4             8332            132059 ns/op
Benchmark_10000_actionMap_10000_random-4                    4606            217786 ns/op
Benchmark_50000_actionSwitch_50000_asc-4                   11432            101400 ns/op
Benchmark_50000_actionSwitchType_50000_asc-4                4981            258677 ns/op
Benchmark_50000_actionMap_50000_asc-4                       2224            569318 ns/op
Benchmark_50000_actionSwitch_50000_random-4                 1832            621763 ns/op
Benchmark_50000_actionSwitchType_50000_random-4             1922            622698 ns/op
Benchmark_50000_actionMap_50000_random-4                    1200            998241 ns/op
```
