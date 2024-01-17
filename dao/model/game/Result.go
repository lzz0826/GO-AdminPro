package game

import "time"
import "gorm.io/gorm"
import "github.com/google/uuid"

// Result 開獎結果 /**
type Result struct {
	ID           string    `gorm:"column:id;type:varchar(40);primary_key;comment:'開獎結果id'"`
	GameID       string    `gorm:"column:game_id;type:varchar(40);not null;comment:'遊戲id'"`
	DrawMethod   int       `gorm:"column:draw_method;type:int;default:1;comment:'開獎方式\n1. 自動\n2. 手動'"`
	ResultType   int       `gorm:"column:result_type;type:int;default:1;comment:'開獎種類\n1. 自有\n2. 外部非官方\n3. 官方'"`
	ResultSource string    `gorm:"column:result_source;type:varchar(40);comment:'開獎來源'"`
	IssueID      string    `gorm:"column:issue_id;type:varchar(40);not null;comment:'期號'"`
	WinOutcomes  string    `gorm:"column:win_outcomes;type:text;comment:'勝market::outcome'"`
	LoseOutcomes string    `gorm:"column:lose_outcomes;type:text;comment:'輸market::outcome'"`
	TieOutcomes  string    `gorm:"column:tie_outcomes;type:text;comment:'平market::outcome'"`
	Memo         string    `gorm:"column:memo;type:text;not null;comment:'贏家的相關說明'"`
	Result       string    `gorm:"column:result;type:varchar(90);not null;comment:'結果'"`
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	IsEnabled    bool      `gorm:"column:is_enabled;type:tinyint(1);not null;default:true;comment:'是否開啟'"`
}

func (Result) TableName() string {
	return "game_result"
}

func (g *Result) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
