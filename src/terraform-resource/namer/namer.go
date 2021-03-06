package namer

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

//go:generate counterfeiter . Namer

type Namer interface {
	RandomName() string
}

func New() Namer {
	return &adjNounNamer{}
}

type adjNounNamer struct{}

func (n adjNounNamer) RandomName() string {
	return fmt.Sprintf("%s-%s", randomdata.Adjective(), randomdata.Noun())
}
