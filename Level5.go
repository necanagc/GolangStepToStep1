package main

import (
	"slices"
	"strings"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	if p.Misses != 0 {
		p.Rating = (float64(p.Goals) + float64(p.Assists/2)) / float64(p.Misses)
	}
	p.Rating = float64(p.Goals) + float64(p.Assists)/2
}

func NewPlayer(name string, goals, misses, assists int) Player {
	player := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}

	player.calculateRating()

	return player
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Goals > b.Goals {
			return -1
		} else if a.Goals < b.Goals {
			return 1
		} else {
			return strings.Compare(a.Name, b.Name)
		}
	})

	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating > b.Rating {
			return -1
		} else if a.Rating < b.Rating {
			return 1
		} else {
			return strings.Compare(a.Name, b.Name)
		}
	})

	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {

		a1 := float64(a.Goals) / float64(a.Misses)
		b1 := float64(b.Goals) / float64(b.Misses)

		if a1 > b1 {
			return -1
		} else if a1 < b1 {
			return 1
		} else {
			return strings.Compare(a.Name, b.Name)
		}
	})

	return players
}
