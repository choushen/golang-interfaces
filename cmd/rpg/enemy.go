package rpg

import "fmt"

type Enemy struct {
	Health      int
	AttackPower int
	Defense     int
	Weapon      string
}

func (e *Enemy) Attack() {
	fmt.Printf("Enemy attacks for %d damage\n", e.AttackPower)
}

func (e *Enemy) Defend() {
	fmt.Printf("Enemy protect itself for %d damage\n", e.Defense)
}

func (e *Enemy) SpecialAttack() {
	fmt.Printf("Enemy special attacks for %d damage\n", e.AttackPower*2)
}

type Goblin struct {
	Enemy
	Homeland string
}

type Orc struct {
	Enemy
	Tribe string
}
