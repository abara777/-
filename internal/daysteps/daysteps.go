package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	str := strings.Split(data, ",")

	if len(str) < 2 {
		return 0, 0 * time.Hour, errors.New("неправильный формат строки")
	}

	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, 0 * time.Hour, err
	}

	t, err := time.ParseDuration(str[1])

	if err != nil {
		return 0, 0 * time.Hour, err
	}

	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, _, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm

	//burnt := WalkingSpentCalories()

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, 0.00)

}
