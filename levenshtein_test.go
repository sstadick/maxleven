package maxleven

import (
	"fmt"
	"testing"

	arbovm "github.com/arbovm/levenshtein"
)

var distanceTests = []struct {
	first  string
	second string
	wanted int
}{
	{"a", "a", 0},
	{"ab", "ab", 0},
	{"ab", "aa", 1},
	{"ab", "aa", 1},
	{"ab", "aaa", 2},
	{"bbb", "a", 3},
	{"kitten", "sitting", 3},
	{"a", "", 1},
	{"", "a", 1},
	{"aa", "aü", 1},
	{"Fön", "Föm", 1},
}
var distanceTestsMaxDist = []struct {
	first   string
	second  string
	wanted  int
	maxDist int
}{
	{"a", "a", 0, 2},
	{"ab", "ab", 0, 2},
	{"ab", "aa", 1, 2},
	{"ab", "aa", 1, 2},
	{"ab", "aaa", 2, 2},
	{"bbb", "a", -1, 2},
	{"kitten", "sitting", -1, 2},
	{"a", "", 1, 2},
	{"", "a", 1, 2},
	{"aa", "aü", 1, 2},
	{"Fön", "Föm", 1, 1},
	{"kitten", "sitting", 3, 4},
	{"a", "", -1, 0},
	{"", "a", -1, 0},
	{"aa", "aü", -1, 0},
	{"Fön", "Föm", 1, 1},
}

func TestDistance(t *testing.T) {

	for index, distanceTest := range distanceTests {
		result := LevDistance([]rune(distanceTest.first), []rune(distanceTest.second), 100)
		if result != distanceTest.wanted {
			output := fmt.Sprintf("%v \t distance of %v and %v should be %v but was %v.",
				index, distanceTest.first, distanceTest.second, distanceTest.wanted, result)
			t.Errorf(output)
		}
	}
	for index, distanceTest := range distanceTestsMaxDist {
		result := LevDistance([]rune(distanceTest.first), []rune(distanceTest.second), distanceTest.maxDist)
		if result != distanceTest.wanted {
			output := fmt.Sprintf("maxdist: %v \t distance of %v and %v should be %v but was %v.",
				index, distanceTest.first, distanceTest.second, distanceTest.wanted, result)
			t.Errorf(output)
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += LevDistance([]rune(s1), []rune(s2), 100)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkDistanceOriginal(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += arbovm.Distance(s1, s2)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}
