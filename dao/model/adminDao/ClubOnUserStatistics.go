package adminDao

type ClubOnUserStatistics struct {
	ClubId    int `gorm:"column:clubId;" json:"clubId"`
	NormalNum int `gorm:"column:normalNum;" json:"normalNum"`
	OpNum     int `gorm:"column:opNum;" json:"opNum"`
}
