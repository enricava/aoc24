package day17

import (
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestB(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	input := strings.NewReader(
		`myrawstring`,
	)

	result := PartB(input)
	expected := 1

	if expected != result {
		log.WithFields(log.Fields{
			"expected": expected,
			"result":   result,
		}).Warn("Failed")

		t.Fail()
	}
}
