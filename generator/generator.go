package generator

import (
	"math/rand"
	"strconv"
	"strings"
)

func GeneratePassword(length int, lowercase bool, uppercase bool, symbols bool, numbers bool) string {
	if length <= 0 {
		panic("Length cannot be less than or equals to 0")
	}
	if !lowercase && !uppercase && !symbols && !numbers {
		panic("All parameters cannot be false")
	}
	password := make([]string, 0, length)
	params := []string{}

	if lowercase {
		password = append(password, generate_random_value(MIN_LOWERCASE, MAX_LOWERCASE))
		params = append(params, LOWERCASE)
	}

	if uppercase {
		password = append(password, generate_random_value(MIN_UPPERCASE, MAX_UPPERCASE))
		params = append(params, UPPERCASE)
	}

	if numbers {
		password = append(password, strconv.Itoa(rand.Intn(10)))
		params = append(params, NUMBERS)
	}

	if symbols {
		password = append(password, generate_random_value(MIN_SYMBOL, MAX_LOWERCASE))
		params = append(params, SYMBOLS)
	}

	for len(password) < length {
		r := rand.Intn(len(params))
		param := params[r]
		switch param {
		case LOWERCASE:
			if lowercase {
				password = append(password, generate_random_value(MIN_LOWERCASE, MAX_LOWERCASE))
			}
		case UPPERCASE:
			if uppercase {
				password = append(password, generate_random_value(MIN_UPPERCASE, MAX_UPPERCASE))
			}
		case NUMBERS:
			if numbers {
				password = append(password, strconv.Itoa(rand.Intn(10)))
			}
		case SYMBOLS:
			if symbols {
				password = append(password, generate_random_value(MIN_SYMBOL, MAX_SYMBOL))
			}
		}
	}

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return strings.Join(password, "")
}

func generate_random_value(lower int, upper int) string {
	range_length := upper - lower + 1
	result := rand.Intn(range_length) + lower
	return string(rune(result))
}
