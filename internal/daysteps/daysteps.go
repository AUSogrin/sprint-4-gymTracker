package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("slice isn`t equal to two")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, fmt.Errorf("conversion to number %v", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("the number of steps must be greater than zero")
	}
	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, 0, fmt.Errorf("no duration data avalible: %v", err)
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	const mInKm = 1000
	distance := (float64(steps) * StepLength) / mInKm
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d\nДистанция составила: %.2f км\nВы сожгли: %.2f ккал\n", steps, distance, calories)
}
