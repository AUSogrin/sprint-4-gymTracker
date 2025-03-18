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
		return 0, 0, fmt.Errorf("слайс не равен двум")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, fmt.Errorf("преобразования в число %v", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("кол-во шагов должно быть юольше нуля")
	}
	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, 0, fmt.Errorf("данные о продолжительности отсутствуют: %v", err)
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := (float64(steps) * StepLength) / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d\nДистанция составила: %.2f км\nВы сожгли: %.2f ккал\n", steps, distance, calories)
}
