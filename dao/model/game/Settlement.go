package game

import "time"
import "gorm.io/gorm"
import "github.com/google/uuid"

// Settlement 結算 /**
type Settlement struct {
	ID           string    `gorm:"column:id;type:varchar(40);primary_key;comment:'結算id'"`
	GameID       string    `gorm:"column:game_id;type:varchar(40);not null;comment:'遊戲id'"`
	IssueID      string    `gorm:"column:issue_id;type:varchar(40);not null;comment:'期號id'"`
	WinOutcomes  string    `gorm:"column:win_outcomes;type:text;not null;comment:'勝market::outcome'"`
	LoseOutcomes string    `gorm:"column:lose_outcomes;type:text;not null;comment:'負market::outcome'"`
	TieOutcomes  string    `gorm:"column:tie_outcomes;type:text;not null;comment:'平market::outcome'"`
	Result       string    `gorm:"column:result;type:varchar(255);comment:'結果'"`
	ResultID     string    `gorm:"column:result_id;type:varchar(40);comment:'結果id'"`
	Memo         string    `gorm:"column:memo;type:text;comment:'備註'"`
	Details      string    `gorm:"column:details;type:text;comment:'其他雜項'"`
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
}

func (Settlement) TableName() string {
	return "game_settlement"
}

func (g *Settlement) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
