package main

import (
	"YadroImpulse_2024/src/models"
	"YadroImpulse_2024/src/process"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Путь к файлу не указан.")
		os.Exit(1)
	}
	filepath := os.Args[1]
	//filepath := "C:/YadroImpulse_2024/data/test2.txt"
	clubInfo := models.ClubInfo{}
	process.ProcessFile(filepath, &clubInfo)

	tablesFree := make([]int, 0)
	tablesMap := make(map[string]models.TableInfo)
	tablesProfit := make(map[int]models.TableStats)
	for i := 1; i <= clubInfo.NumberOfTables; i++ {
		tablesProfit[i] = models.TableStats{TotalProfit: 0, TotalMinutes: 0}
		//tablesProfit[i]
		tablesFree = append(tablesFree, i)
	}
	clubMap := models.ClubMap{
		UsersQueue:   make([]string, 0),
		TablesFree:   tablesFree,
		TablesMap:    tablesMap,
		TablesProfit: tablesProfit}
	process.ProcessEvents(clubMap, clubInfo, filepath)

}
