package main

import (
	"github.com/choushen/homeless-golang-code/cmd/rpg"
)

func main() {
	// This is a comment
	//fmt.Println("Hello, World!")

	// Call the TrelloAPI function
	//api.TrelloAPI()

	orcCharacter := rpg.Orc{
		Enemy: rpg.Enemy{
			Health:      100,
			AttackPower: 10,
			Defense:     5,
			Weapon:      "Axe",
		},
		Tribe: "Horde",
	}

	rpg.CounterAttack(&orcCharacter)

	goblinCharacter := rpg.Goblin{
		Enemy: rpg.Enemy{
			Health:      50,
			AttackPower: 5,
			Defense:     2,
			Weapon:      "Dagger",
		},
		Homeland: "Forest",
	}

	rpg.CounterAttack(&goblinCharacter)
}
