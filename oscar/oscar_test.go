package oscar_test

import (
	"testing"

	"github.com/hello/go-by-yod/oscar"
)

func TestActorWhoWonMoreThanOneAward(t *testing.T) {
	oscar.ActorWhoWonMoreThanOneAward("./oscar_age_male.csv")
}
