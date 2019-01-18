package combinations

import (
	"math"
	"strings"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func next(position uint, initial []string, first []string, last []string, characters []string, separator string, returnChan *chan string) {
	if position == math.MaxUint64 {
		return
	}

	if Equal(initial, last) {
		return
	}

	for _, char := range characters {
		if initial[position] == char {
			next(position-1, append([]string(nil), initial...), first, last, characters, separator, returnChan)
			continue
		}
		initial[position] = char
		if !Equal(initial, first) {
			*returnChan <- strings.Join(initial, separator)
		}
		next(position-1, append([]string(nil), initial...), first, last, characters, separator, returnChan)
	}

}

func makeChan(characters []string, length uint) chan string {
	possibilities := uint(math.Pow(float64(len(characters)), float64(length)))

	return make(chan string, possibilities)
}

func generator(characters []string, length uint, separator string, returnChan *chan string) {

	first := make([]string, length)
	last := make([]string, length)

	for key := range first {
		first[key] = characters[0]
	}
	for key := range last {
		last[key] = characters[len(characters)-1]
	}
	*returnChan <- strings.Join(first, separator)

	next(length-1, first, append([]string(nil), first...), last, characters, separator, returnChan)

	close(*returnChan)
}
