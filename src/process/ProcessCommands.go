package process

import (
	"YadroImpulse_2024/src/models"
	my "YadroImpulse_2024/src/uility"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var clubMap models.ClubMap
var clubInfo models.ClubInfo
var tableInfo models.TableInfo

func ProcessEvents(clubMapTemp models.ClubMap, clubInfoTemp models.ClubInfo, filepath string) {
	clubMap = clubMapTemp
	clubInfo = clubInfoTemp
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(clubInfo.OpenTime)
	scanner := bufio.NewScanner(file)
	for i := 0; i < 3; i++ {
		scanner.Scan()
	}

	for scanner.Scan() {
		tempStr := scanner.Text()
		fmt.Println(tempStr)

		fields := strings.Fields(tempStr)
		action, _ := strconv.Atoi(fields[1])

		switch action {
		case 1:
			time := fields[0]
			name := fields[2]
			HandleCome(time, name)
		case 2:
			time := fields[0]
			name := fields[2]
			tableNumber, _ := strconv.Atoi(fields[3])
			HandleTakeTable(time, name, tableNumber)
		case 3:
			time := fields[0]
			name := fields[2]
			HandleWait(time, name)
		case 4:
			time := fields[0]
			name := fields[2]
			HandleLeft(time, name)
		}
	}
	QueueForKick()
	ShowStats()
}
func HandleCome(time, name string) {
	if my.ContainsValue(clubMap.TablesMap, name) || my.ContainsUser(clubMap.UsersQueue, name) {
		fmt.Println(time + " 13 YouShallNotPass")
	} else if time < clubInfo.OpenTime || time > clubInfo.CloseTime {
		fmt.Println(time + " 13 NotOpenYet")
	} else {
		clubMap.UsersQueue = append(clubMap.UsersQueue, name)
	}
}

func HandleTakeTable(time, name string, tableNumber int) {
	if !my.ContainsValue(clubMap.TablesMap, name) && !my.ContainsUser(clubMap.UsersQueue, name) {
		fmt.Println(time + " 13 ClientUnknown")
	} else if !my.ContainsTable(clubMap.TablesFree, tableNumber) {
		fmt.Println(time + " 13 PlaceIsBusy")
	} else if my.ContainsValue(clubMap.TablesMap, name) && !my.ContainsUser(clubMap.UsersQueue, name) {
		UpdateTableStats(time, name)
		my.RemoveTableFree(&clubMap.TablesFree, tableNumber)
		my.AddTableFree(&clubMap.TablesFree, clubMap.TablesMap[name].TableId)
		UpdateTableInfo(tableNumber, time, name)
	} else {
		my.RemoveTableFree(&clubMap.TablesFree, tableNumber)
		my.RemoveUser(&clubMap.UsersQueue, name)
		UpdateTableInfo(tableNumber, time, name)

	}

}

func HandleWait(time, name string) {
	if len(clubMap.TablesFree) > 0 {
		fmt.Println(time + " 13 ICanWaitNoLonger!")
	} else if len(clubMap.UsersQueue) > clubInfo.NumberOfTables {
		my.RemoveUser(&clubMap.UsersQueue, name)
		fmt.Println(time + " 11 " + name)
	}
}

func HandleLeft(time, name string) {
	if !my.ContainsValue(clubMap.TablesMap, name) && !my.ContainsUser(clubMap.UsersQueue, name) {
		fmt.Println(time + " 13 ClientUnknown")
	} else if my.ContainsValue(clubMap.TablesMap, name) {

		UpdateTableStats(time, name)
		table := clubMap.TablesMap[name].TableId
		my.AddTableFree(&clubMap.TablesFree, table)
		delete(clubMap.TablesMap, name)
		if len(clubMap.UsersQueue) != 0 {
			nameFirst := clubMap.UsersQueue[0]
			HandleTakeTable(time, nameFirst, table)
			fmt.Println(time + " 12 " + nameFirst + " " + strconv.Itoa(table))
		}
	} else {
		my.RemoveUser(&clubMap.UsersQueue, name)
	}

}
func UpdateTableInfo(tableNumber int, time, name string) {
	tableInfo.TableId = tableNumber
	tableInfo.TimeStart = time
	clubMap.TablesMap[name] = tableInfo

}
func HandleKick(time, name string) {
	if len(clubMap.TablesFree) != clubInfo.NumberOfTables && !my.ContainsUser(clubMap.UsersQueue, name) {
		UpdateTableStats(time, name)
		table := clubMap.TablesMap[name].TableId

		my.AddTableFree(&clubMap.TablesFree, table)
		delete(clubMap.TablesMap, name)
	}
	fmt.Println(time + " 11 " + name)
}
func UpdateTableStats(time, name string) {
	table := clubMap.TablesMap[name].TableId
	timeStart := clubMap.TablesMap[name].TimeStart
	minutesOnTable := my.GetPassedHours(time, timeStart)
	timeOnTable := (minutesOnTable + 59) / 60
	cost := clubInfo.CostPerHour * timeOnTable

	tableStats := clubMap.TablesProfit[table]
	tableStats.TotalProfit += cost
	tableStats.TotalMinutes += minutesOnTable
	clubMap.TablesProfit[table] = tableStats
}
func QueueForKick() {
	names := make([]string, 0, len(clubMap.TablesMap))
	for key := range clubMap.TablesMap {
		names = append(names, key)
	}
	queueForKick := append(names, clubMap.UsersQueue...)
	sort.Strings(queueForKick)
	if len(clubMap.TablesFree) != clubInfo.NumberOfTables {
		for _, key := range queueForKick {
			HandleKick(clubInfo.CloseTime, key)
		}
	}
	clubMap.UsersQueue = nil
	fmt.Println(clubInfo.CloseTime)
}
func ShowStats() {
	var keys []int
	for table, _ := range clubMap.TablesProfit {
		keys = append(keys, table)
	}
	sort.Ints(keys)
	for _, table := range keys {
		stats := clubMap.TablesProfit[table]
		fmt.Println(strconv.Itoa(table) + " " + strconv.Itoa(stats.TotalProfit) + " " + my.ConvertToTime(stats.TotalMinutes))
	}
}
