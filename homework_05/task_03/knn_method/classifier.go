package k_nn_method

import (
	"sort"

	"HomeWork_05/Task_03/models"
)

type Neighbor struct {
	Object   models.Object
	Distance float64
}

// 'Classifier' implements the k-nearest neighbors algorithm
type Classifier struct {
	Data       []models.Object
	K          int
	Normalizer *models.Normalizer
}

// 'NewClassifier' creates a new KNN classifier with normalization
func NewClassifier(data []models.Object, k int) *Classifier {
	// Creating a normalizer based on training data
	normalizer := models.NewNormalizer(data)

	// Normalizing the training data
	normalizedData := normalizer.NormalizeAll(data)

	return &Classifier{
		Data:       normalizedData,
		K:          k,
		Normalizer: normalizer,
	}
}

// 'Classify' predicts a label for an unknown object
func (c *Classifier) Classify(x, y float64) string {
	// Create an unknown object
	unknown := models.Object{X: x, Y: y}

	// Normalize the unknown object
	normalizedUnknown := c.Normalizer.Normalize(&unknown)

	// Calculate distances to all known objects
	neighbors := make([]Neighbor, len(c.Data))
	for i, obj := range c.Data {
		neighbors[i] = Neighbor{
			Object:   obj,
			Distance: normalizedUnknown.Distance(&obj),
		}
	}

	// Sort by distance
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance < neighbors[j].Distance
	})

	// Count votes from k nearest neighbors
	votes := make(map[string]int)
	for i := 0; i < c.K && i < len(neighbors); i++ {
		label := neighbors[i].Object.Name
		votes[label]++
	}

	// Finding the most common label
	maxVotes := 0
	bestLabel := ""
	for label, count := range votes {
		if count > maxVotes {
			maxVotes = count
			bestLabel = label
		}
	}

	return bestLabel
}
