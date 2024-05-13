package main

import "fmt"

// 同一インタフェースの具象クラスを識別子で生成するパターン
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

func (g *Gun) String() string {
	return fmt.Sprintf("Gun: %s\nPower: %d", g.getName(), g.getPower())
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

type GunType int

const (
	AK47 GunType = iota
	Musket
)

func NewGun(gunType GunType) (IGun, error) {
	switch gunType {
	case AK47:
		return newAk47(), nil
	case Musket:
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := NewGun(AK47)
	musket, _ := NewGun(Musket)

	fmt.Println(ak47)
	fmt.Println(musket)
}
