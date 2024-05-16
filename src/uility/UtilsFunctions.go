package uility

import (
	"YadroImpulse_2024/src/models"
	"fmt"
	"strconv"
	"strings"
)

func ContainsValue(m map[string]models.TableInfo, value string) bool {
	for name, _ := range m {
		if name == value {
			return true
		}
	}
	return false
}
func ContainsUser(users []string, user string) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}
func ContainsTable(tables []int, value int) bool {
	for _, item := range tables {
		if item == value {
			return true
		}
	}
	return false
}
func RemoveTableFree(slice *[]int, value int) {
	for i := 0; i < len(*slice); i++ {
		if (*slice)[i] == value {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return
		}
	}
}
func AddTableFree(slice *[]int, value int) {
	if ContainsTable(*slice, value) {
		return
	}
	*slice = append(*slice, value)
}
func RemoveUser(users *[]string, user string) {
	for i, u := range *users {
		if u == user {
			*users = append((*users)[:i], (*users)[i+1:]...)
			return
		}
	}
}

func GetPassedHours(time1, time2 string) int {
	parts1 := strings.Split(time1, ":")
	parts2 := strings.Split(time2, ":")

	hours1, _ := strconv.Atoi(parts1[0])
	hours2, _ := strconv.Atoi(parts2[0])
	minutes1, _ := strconv.Atoi(parts1[1])
	minutes2, _ := strconv.Atoi(parts2[1])

	totalMinutes1 := hours1*60 + minutes1
	totalMinutes2 := hours2*60 + minutes2

	diffMinutes := totalMinutes1 - totalMinutes2

	return diffMinutes
}
func ConvertToTime(minutes int) string {
	hours := minutes / 60
	minutesRem := minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutesRem)
}
