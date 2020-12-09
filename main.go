package main

func (as Animals) ActionSwitch() {
	for _, a := range as {
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
	}
}

func (as Animals) ActionSwitchType() {
	for _, a := range as {
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
	}
}

func (as Animals) ActionMap() {
	for _, a := range as {
		handler := actionHandler[a.Type()]
		handler(a)
	}
}

var actionHandler = map[AnimalType]func(a Animal){
	AnimalTypeDog:    func(a Animal) { a.(Dog).Bark() },
	AnimalTypeCat:    func(a Animal) { a.(Cat).Meow() },
	AnimalTypeBird:   func(a Animal) { a.(Bird).Fly() },
	AnimalTypeBear:   func(a Animal) { a.(Bear).Roar() },
	AnimalTypeRabbit: func(a Animal) { a.(Rabbit).Jump() },
}

type Animal interface {
	Type() AnimalType
}

type Animals []Animal

type AnimalType int

const (
	AnimalTypeDog AnimalType = iota + 1
	AnimalTypeCat
	AnimalTypeBird
	AnimalTypeBear
	AnimalTypeRabbit
)

type Common struct {
	AnimalType AnimalType
}

func (c Common) Type() AnimalType {
	return c.AnimalType
}

type Dog struct {
	Common
}

func NewDog() Dog {
	return Dog{Common: Common{AnimalType: AnimalTypeDog}}
}

func (d Dog) Bark() {}

type Cat struct {
	Common
}

func NewCat() Cat {
	return Cat{Common: Common{AnimalType: AnimalTypeCat}}
}

func (c Cat) Meow() {}

type Bird struct {
	Common
}

func NewBird() Bird {
	return Bird{Common: Common{AnimalType: AnimalTypeBird}}
}

func (b Bird) Fly() {}

type Bear struct {
	Common
}

func NewBear() Bear {
	return Bear{Common: Common{AnimalType: AnimalTypeBear}}
}

func (b Bear) Roar() {}

type Rabbit struct {
	Common
}

func NewRabbit() Rabbit {
	return Rabbit{Common: Common{AnimalType: AnimalTypeRabbit}}
}

func (r Rabbit) Jump() {}
