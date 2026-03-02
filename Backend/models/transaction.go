package models

type TransactionType string

// 收支类型枚举
const (
	Income  TransactionType = "INCOME"
	Expense TransactionType = "EXPENSE"
)

// 分类
type Category struct {
	ID        string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	Icon      string `gorm:"type:varchar(50)" json:"icon"`
	SortOrder int    `gorm:"default:0" json:"sort_order"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	IsDeleted bool   `gorm:"default:false" json:"is_deleted"`
}

// 收支对象
type Payee struct {
	ID        string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	Notes     string `gorm:"type:text" json:"notes"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	IsDeleted bool   `gorm:"default:false" json:"is_deleted"`
}

// 账户（支付方式）
type Account struct {
	ID        string  `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name      string  `gorm:"type:varchar(50);not null" json:"name"`
	Balance   float64 `gorm:"not null;default:0" json:"balance"`
	UpdatedAt int64   `gorm:"autoUpdateTime:milli" json:"updated_at"`
	IsDeleted bool    `gorm:"default:false" json:"is_deleted"`
}

// 交易记录
type Transaction struct {
	ID         string          `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Amount     float64         `gorm:"not null" json:"amount"`
	Type       TransactionType `gorm:"type:varchar(20);not null" json:"type" binding:"required,oneof=INCOME EXPENSE"`
	CategoryID string          `gorm:"type:varchar(36);index" json:"category_id"`
	Category   Category        `gorm:"foreignKey:CategoryID" json:"-"`
	PayeeID    string          `gorm:"type:varchar(36);index" json:"payee_id"`
	Payee      Payee           `gorm:"foreignKey:PayeeID" json:"-"`
	AccountID  string          `gorm:"type:varchar(36);index" json:"account_id"`
	Account    Account         `gorm:"foreignKey:AccountID" json:"-"`
	RecordTime int64           `gorm:"not null" json:"record_time"`
	Note       string          `gorm:"type:text" json:"note"`
	UpdatedAt  int64           `gorm:"autoUpdateTime:milli" json:"updated_at"`
	IsDeleted  bool            `gorm:"default:false" json:"is_deleted"`
}
