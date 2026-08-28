package models

import "math"

// 'Object' represents a data point with a label and two features
type Object struct {
	Name string
	X, Y float64
}

// 'Distance' calculates the Euclidean distance between two objects
func (o *Object) Distance(other *Object) float64 {
	dx := o.X - other.X
	dy := o.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// 'Normalizer' stores parameters for minimax normalization
type Normalizer struct {
	MinX, MaxX float64
	MinY, MaxY float64
}

// 'NewNormalizer' calculates the normalization parameters for a set of objects
func NewNormalizer(objects []Object) *Normalizer {
	if len(objects) == 0 {
		return &Normalizer{0, 1, 0, 1}
	}

	n := &Normalizer{
		MinX: objects[0].X,
		MaxX: objects[0].X,
		MinY: objects[0].Y,
		MaxY: objects[0].Y,
	}

	for _, obj := range objects {
		n.MinX = min(n.MinX, obj.X)
		n.MaxX = max(n.MaxX, obj.X)
		n.MinY = min(n.MinY, obj.Y)
		n.MaxY = max(n.MaxY, obj.Y)
	}

	return n
}

// 'Normalize' normalizes an object using the formula (x - min) / (max - min)
func (n *Normalizer) Normalize(obj *Object) Object {
	normalized := Object{Name: obj.Name}

	// Normalization X
	if n.MaxX-n.MinX != 0 {
		normalized.X = (obj.X - n.MinX) / (n.MaxX - n.MinX)
	} else {
		normalized.X = 0
	}

	// Normalization Y
	if n.MaxY-n.MinY != 0 {
		normalized.Y = (obj.Y - n.MinY) / (n.MaxY - n.MinY)
	} else {
		normalized.Y = 0
	}

	return normalized
}

// 'NormalizeAll' normalizes the entire data set
func (n *Normalizer) NormalizeAll(object []Object) []Object {
	result := make([]Object, len(object))
	for i, obj := range object {
		result[i] = n.Normalize(&obj)
	}
	return result
}
