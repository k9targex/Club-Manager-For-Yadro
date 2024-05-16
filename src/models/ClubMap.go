package models

type ClubMap struct {
	UsersQueue   []string
	TablesFree   []int
	TablesMap    map[string]TableInfo
	TablesProfit map[int]TableStats
}
