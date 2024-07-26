package rpg

import "fmt"

type Enemy struct {
	health  int
	attack  int
	defence int
	weapon  string
}

func (e *Enemy) Attack() {
	fmt.Printf("Enemy attacks for %d damage\n", e.attack)
}

func (e *Enemy) Defend() {
	fmt.Printf("Enemy protect itself for %d damage\n", e.defence)
}

func (e *Enemy) SpecialAttack() {
	fmt.Printf("Enemy special attacks for %d damage\n", e.attack*2)
}

type Goblin struct {
	Enemy
	homeland string
}

type Orc struct {
	Enemy
	tribe string
}
