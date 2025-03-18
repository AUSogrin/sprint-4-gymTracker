package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 3 {
		return 0, "", 0, fmt.Errorf("слайс не равен 3")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("преобразования в число %v", err)
	}
	duration, err := time.ParseDuration(s[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("данные о продолжительости отсутствуют %v", err)
	}
	activity := s[1]

	return steps, activity, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
func distance(steps int) float64 {
	return (float64(steps) * lenStep) / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return distance(steps) / duration.Hours()
}

// TrainingInfo возвращает строку с информацией о тренировке.
func TrainingInfo(data string, weight, height float64) string {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}
	var calories float64

	switch activity {
	case "Бег":
		calories = RunningSpentCalories(steps, weight, duration)
	case "Ходьба":
		calories = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "Неизвестная тренировка"
	}
	distance := distance(steps)
	meanSpeed := meanSpeed(steps, duration)

	return fmt.Sprintf("Тип активности: %s\nДистанция: %.2f км\nСредняя скорость: %.2f км/ч\nСожжено калорий: %.2f\n", activity, distance, meanSpeed, calories)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	meanSpeed := meanSpeed(steps, duration)

	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight

}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	meanSpeed := meanSpeed(steps, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH

}
