package cardinalities

import "github.com/steve-care-software/webx/domain/databases/entities"

type cardinality struct {
	entity entities.Entity
	min    uint
	pMax   *uint
}

func createCardinality(
	entity entities.Entity,
	min uint,
) Cardinality {
	return createCardinalityInternally(entity, min, nil)
}

func createCardinalityWithMax(
	entity entities.Entity,
	min uint,
	pMax *uint,
) Cardinality {
	return createCardinalityInternally(entity, min, pMax)
}

func createCardinalityInternally(
	entity entities.Entity,
	min uint,
	pMax *uint,
) Cardinality {
	out := cardinality{
		entity: entity,
		min:    min,
		pMax:   pMax,
	}

	return &out
}

// Entity returns the entity
func (obj *cardinality) Entity() entities.Entity {
	return obj.entity
}

// Min returns the minimum
func (obj *cardinality) Min() uint {
	return obj.min
}

// HasMax returns true if there is a max, false otherwise
func (obj *cardinality) HasMax() bool {
	return obj.pMax != nil
}

// Max returns the max, if any
func (obj *cardinality) Max() *uint {
	return obj.pMax
}
