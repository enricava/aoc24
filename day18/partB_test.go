package day18

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
		`5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
	)

	result := solveB(input, 7, 10000, 15)
	expected := 20

	if expected != result {
		log.WithFields(log.Fields{
			"expected": expected,
			"result":   result,
		}).Warn("Failed")

		t.Fail()
	}
}
