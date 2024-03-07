package test

import (
	"testing"

	"github.com/schoolboybru/hockey-data/cmd/data"
)

func TestGetGame(t *testing.T) {
	d := data.GameDayData{}

	err := d.Get()

	if err != nil {
		t.Error(err)
	}
}
