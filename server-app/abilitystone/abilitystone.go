package abilitystone

import (
	"math/rand"
	"time"
)

type Line uint8

const (
	Undefined Line = iota
	Line1
	Line2
	Line3
)

type AbilityStone struct {
	Abilities          map[Line][]bool `json:"abilities"`
	CurrentProbability int             `json:"currentProbability"`
}

func (a AbilityStone) GetCurrentAbilities() map[Line][]bool {
	return a.Abilities
}

func (a AbilityStone) GetCurrentProbabilities() int {
	return a.CurrentProbability
}

func (a *AbilityStone) RollAbilityLine(line Line) {
	a.Abilities[line] = append(a.Abilities[line], a.roll())
}

func (a *AbilityStone) roll() bool {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomRoller := rand.New(randomSource)
	rolledValue := randomRoller.Intn(100)

	if rolledValue <= a.CurrentProbability {
		if a.CurrentProbability > 25 {
			a.CurrentProbability -= 10
		}
		return true
	}
	if a.CurrentProbability < 75 {
		a.CurrentProbability += 10
	}
	return false
}

func NewAbilityStone() *AbilityStone {
	abilities := make(map[Line][]bool)
	abilities[Line1] = make([]bool, 0, 10)
	abilities[Line2] = make([]bool, 0, 10)
	abilities[Line3] = make([]bool, 0, 10)

	currentProbability := 75

	return &AbilityStone{
		abilities,
		currentProbability,
	}
}
