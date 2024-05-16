package models

type ClubInfo struct {
	NumberOfTables int    // число столов
	CostPerHour    int    // цена за час
	OpenTime       string // время открытия
	CloseTime      string // время закрытия
}
