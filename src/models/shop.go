package models

import (
	"math/rand"
	"time"
)

// Shop は、Shop のモデル。
type Shop struct {
	ID        string    `dynamo:"id" json:"id"`
	Name      string    `dynamo:"name" json:"name"`
	URL       string    `dynamo:"url" json:"url"`
	Memo      string    `dynamo:"memo" json:"memo"`
	CreatedAt time.Time `dynamo:"created_at" json:"created_at"`
}

// ChoiceShopsRandomly は、Shop をランダムに選択する。
func ChoiceShopsRandomly(shops []Shop) []Shop{
	const numOfChoice = 3

	alreadyChosen := make([]int, 0)
	chosen := make([]Shop, 0)
	for {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(len(shops))

		if !isAlreadyExistence(j, alreadyChosen) {
			chosen = append(chosen, shops[j])
		}

		alreadyChosen = append(alreadyChosen, j)

		if len(chosen) == numOfChoice {
			break
		}
	}
	return chosen
}

func isAlreadyExistence(a int, s []int) bool {
	for _, v := range s {
		if a == v {
			return true
		}
	}
	return false
}
