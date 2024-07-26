package rpg

type Mob interface {
	Attack
	Defend
}

type Attack interface {
	Attack()
	SpecialAttack()
}

type Defend interface {
	Defend()
}

func CounterAttack(m Mob) {
	m.Defend()
	m.Attack()
}
