package pokeapi

import (
	"math"
	"math/rand/v2"
)

// Calculate catch probability (0-100%) based on base experience
func calculateCatchProbability(baseExp int) float64 {
	maxExp := 608 // Current known maximum base experience
	normalized := float64(baseExp) / float64(maxExp)
	return 80 * (1 - math.Pow(normalized, 1.5)) // Curved difficulty
}

// CatchPokemon - Attempts to catch a Pokemon
func (p *Pokemon) AttemptCatch() bool {
	roll := rand.Float64() * 100 // Random number 0-100
	catchThreshold := calculateCatchProbability(p.BaseExperience)
	return roll <= catchThreshold
}
