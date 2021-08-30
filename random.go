package main

import (
	"errors"
)

var errMinCannotBeGreaterThanMax = errors.New("<min> cannot be greater than <max>")

type randomParser interface {
	ParseUpper(value string) (randomUpper, error)
}

type randomUpper interface {
	ParseLower(value string) (randomRange, error)
	DefaultLower() randomRange
}

type randomRange interface {
	IsLowerLargerThanUpper() bool
	PrintRandomValue()
}

var parsers = []randomParser{
	randomInt{},
	randomFloat{},
	randomDuration{},
	randomTime{},
	randomDate{},
}

var defaultParser randomParser = randomInt{}
var defaultRandomRange randomRange = randomInt{
	upper: 10,
	lower: 0,
}
