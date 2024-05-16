package process

import (
	"YadroImpulse_2024/src/models"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ProcessFile(filename string, info *models.ClubInfo) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		os.Exit(1)
	}
	defer file.Close()
	intRegex := regexp.MustCompile(`^\d+$`)
	timeRegex := regexp.MustCompile(`([01]?[0-9]|2[0-3]):([0-5]\d) ([01]?[0-9]|2[0-3]):([0-5]\d)$`)
	eventRegex := regexp.MustCompile(`(?:0[0-9]|1[0-9]|2[0-3]):[0-5][0-9] ([1-4]) ([a-z0-9_-]+(?: \d+)?)$`)
	regexList := []*regexp.Regexp{
		intRegex,
		timeRegex,
		intRegex,
	}
	scanner := bufio.NewScanner(file)
	for i, regex := range regexList {
		scanner.Scan()
		line := scanner.Text()
		if regex.MatchString(line) {
			switch i {
			case 0:
				info.NumberOfTables, err = strconv.Atoi(line)
			case 1:
				timeRange := strings.Split(line, " ")
				info.OpenTime = timeRange[0]
				info.CloseTime = timeRange[1]
			case 2:
				info.CostPerHour, err = strconv.Atoi(line)
			}
		} else {
			fmt.Println(line)
			os.Exit(1)
		}

	}

	time := ""
	for scanner.Scan() {
		line := scanner.Text()
		if eventRegex.MatchString(line) {
			tempFotTime := strings.SplitN(line, " ", 2)
			tempForTable := strings.Fields(tempFotTime[1])
			if time <= tempFotTime[0] {
				time = tempFotTime[0]
			} else {
				fmt.Println(line)
				os.Exit(1)
			}
			if len(tempForTable) > 2 {
				table, err := strconv.Atoi(tempForTable[len(tempForTable)-1])
				if err != nil {
					os.Exit(1)
				}
				if table > info.NumberOfTables || table < 1 {
					fmt.Println(line)
					os.Exit(1)
				}
			}
		} else {
			fmt.Println(line)
			os.Exit(1)
		}
	}
}
