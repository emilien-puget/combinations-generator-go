package combinations

import (
	"math"
	"strings"
)

func next(position uint, initial []string, first []string, last []string, characters []string, separator string, returnChan *chan string) {
	if position == math.MaxUint64 {
		return
	}

	for _, char := range characters {
		if initial[position] != char {
			initial[position] = char
			*returnChan <- strings.Join(initial, separator)
		}
		next(position-1, append([]string(nil), initial...), first, last, characters, separator, returnChan)
	}
}

func MakeChan(characters []string, length uint) chan string {
	possibilities := uint(math.Pow(float64(len(characters)), float64(length)))

	return make(chan string, possibilities)
}

func Generator(characters []string, length uint, separator string, returnChan *chan string) {

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
