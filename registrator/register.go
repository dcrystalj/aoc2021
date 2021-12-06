package registrator

import "fmt"


var register map[string]func([]string) string

func Register(day string, f func([]string) string) {
	register[day] = f
}

func Run(day string, lines []string) string {
	f, exists := register[day]
	if !exists {
		fmt.Printf("Day %s not found or not registered.", day)
		panic("")
	}
	return f(lines)
}

func init() {
	register = make(map[string]func([]string) string, 0)
}
