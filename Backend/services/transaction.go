package services

import (
	"Drawer/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct{}

// ========交易服务=========

// 检查ID合法性
func IsValidID(model interface{}, id string) bool {
	if id == "" {
		return false
	}

	var count int64
	models.DB.Model(model).Where("id = ? AND is_deleted = false", id).Count(&count)
	return count > 0
}

// 创建交易记录
func (t *Transaction) CreateTransaction(tx *models.Transaction) error {
	if !IsValidID(&models.Category{}, tx.CategoryID) {
		return errors.New("非法的类型ID")
	}
	if !IsValidID(&models.Payee{}, tx.PayeeID) {
		return errors.New("非法的收支对象ID")
	}
	if !IsValidID(&models.Account{}, tx.AccountID) {
		return errors.New("非法的账户ID")
	}

	tx.ID = uuid.New().String()
	if tx.RecordTime == 0 {
		tx.RecordTime = time.Now().UnixMilli()
	}

	return models.DB.Create(tx).Error
}

// --------查询交易记录列表---------

// 查询参数结构体
type TransactionQueryParam struct {
	Type       models.TransactionType `json:"type" form:"type"`
	CategoryID string                 `json:"category_id" form:"category_id"`
	PayeeID    string                 `json:"payee_id" form:"payee_id"`
	AccountID  string                 `json:"account_id" form:"account_id"`
	StartTime  int64                  `json:"start_time" form:"start_time"`
	EndTime    int64                  `json:"end_time" form:"end_time"`
	Keyword    string                 `json:"keyword" form:"keyword"`
	MinAmount  float64                `json:"min_amount" form:"min_amount"`
	MaxAmount  float64                `json:"max_amount" form:"max_amount"`
}

// 查询结果分析结构体(统计数据)
type StatItem struct {
	TargetID string  `json:"target_id"`
	Amount   float64 `json:"amount"`
	Ratio    float64 `json:"ratio"`
}

// 收入/支出统计数据结构体
type BreakdownStats struct {
	Total      float64    `json:"total"`
	Ratio      float64    `json:"ratio"`
	ByCategory []StatItem `json:"by_category"`
	ByPayee    []StatItem `json:"by_payee"`
	ByAccount  []StatItem `json:"by_account"`
}

// 记账总数据结构体(结余\收入\支出)
type TransactionStats struct {
	Balance float64        `json:"balance"`
	Income  BreakdownStats `json:"income"`
	Expense BreakdownStats `json:"expense"`
}

// 交易记录列表与统计数据结构体
type TransactionQueryResult struct {
	List  []models.Transaction `json:"list"`
	Stats TransactionStats     `json:"stats"`
}

// 查询交易记录列表并计算统计数据
func (t *Transaction) QueryTransactions(param TransactionQueryParam) (*TransactionQueryResult, error) {
	query := models.DB.Model(&models.Transaction{}).Where("is_deleted = false")

	if param.Type != "" {
		query = query.Where("type = ?", param.Type)
	}
	if param.CategoryID != "" {
		query = query.Where("category_id = ?", param.CategoryID)
	}
	if param.PayeeID != "" {
		query = query.Where("payee_id = ?", param.PayeeID)
	}
	if param.AccountID != "" {
		query = query.Where("account_id = ?", param.AccountID)
	}
	if param.StartTime > 0 {
		query = query.Where("record_time >= ?", param.StartTime)
	}
	if param.EndTime > 0 {
		query = query.Where("record_time <= ?", param.EndTime)
	}
	if param.MinAmount > 0 {
		query = query.Where("amount >= ?", param.MinAmount)
	}
	if param.MaxAmount > 0 {
		query = query.Where("amount <= ?", param.MaxAmount)
	}
	if param.Keyword != "" {
		query = query.Where("Note LIKE ?", "%"+param.Keyword+"%")
	}

	var list []models.Transaction
	if err := query.Preload("Category").Preload("Payee").Preload("Account").Order("record_time DESC").Find(&list).Error; err != nil {
		return nil, err
	}

	result := &TransactionQueryResult{
		List:  list,
		Stats: calculateStats(list),
	}

	return result, nil
}

// 统计结果计算
func calculateStats(transactions []models.Transaction) TransactionStats {
	var stats TransactionStats

	incomeCategoryMap := make(map[string]float64)
	incomePayeeMap := make(map[string]float64)
	incomeAccountMap := make(map[string]float64)
	expenseCategoryMap := make(map[string]float64)
	expensePayeeMap := make(map[string]float64)
	expenseAccountMap := make(map[string]float64)

	for _, tx := range transactions {
		if tx.Type == models.Income {
			stats.Income.Total += tx.Amount
			incomeCategoryMap[tx.CategoryID] += tx.Amount
			incomePayeeMap[tx.PayeeID] += tx.Amount
			incomeAccountMap[tx.AccountID] += tx.Amount
		} else {
			stats.Expense.Total += tx.Amount
			expenseCategoryMap[tx.CategoryID] += tx.Amount
			expensePayeeMap[tx.PayeeID] += tx.Amount
			expenseAccountMap[tx.AccountID] += tx.Amount
		}
	}

	stats.Balance = stats.Income.Total - stats.Expense.Total
	totalFlow := stats.Income.Total + stats.Expense.Total
	if totalFlow > 0 {
		stats.Income.Ratio = stats.Income.Total / totalFlow
		stats.Expense.Ratio = stats.Expense.Total / totalFlow
	}

	stats.Income.ByCategory = buildStatItems(incomeCategoryMap, stats.Income.Total)
	stats.Income.ByPayee = buildStatItems(incomePayeeMap, stats.Income.Total)
	stats.Income.ByAccount = buildStatItems(incomeAccountMap, stats.Income.Total)
	stats.Expense.ByCategory = buildStatItems(expenseCategoryMap, stats.Expense.Total)
	stats.Expense.ByPayee = buildStatItems(expensePayeeMap, stats.Expense.Total)
	stats.Expense.ByAccount = buildStatItems(expenseAccountMap, stats.Expense.Total)

	return stats
}

// 计算比例并构建统计项列表
func buildStatItems(dataMap map[string]float64, total float64) []StatItem {
	items := make([]StatItem, 0, len(dataMap))
	for id, amount := range dataMap {
		ratio := 0.0
		if total > 0 {
			ratio = amount / total
		}
		items = append(items, StatItem{
			TargetID: id,
			Amount:   amount,
			Ratio:    ratio,
		})
	}
	return items
}

// 更新交易记录
func (t *Transaction) UpdateTransaction(id string, updates *models.Transaction) (*models.Transaction, error) {
	if updates.CategoryID != "" && !IsValidID(&models.Category{}, updates.CategoryID) {
		return nil, errors.New("非法的类型ID")
	}
	if updates.PayeeID != "" && !IsValidID(&models.Payee{}, updates.PayeeID) {
		return nil, errors.New("非法的收支对象ID")
	}
	if updates.AccountID != "" && !IsValidID(&models.Account{}, updates.AccountID) {
		return nil, errors.New("非法的账户ID")
	}

	var tx models.Transaction
	if err := models.DB.Where("id = ? AND is_deleted = false", id).First(&tx).Error; err != nil {
		return nil, err
	}
	updates.ID = tx.ID

	if err := models.DB.Model(&tx).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &tx, nil
}

// 删除交易记录（软删除）
func (t *Transaction) DeleteTransaction(id string) error {
	return models.DB.Model(&models.Transaction{}).Where("id = ?", id).Update("is_deleted", true).Error
}

// =======类型服务========

// 创建类型
func (t *Transaction) CreateCategory(cat *models.Category) error {
	cat.ID = uuid.New().String()
	return models.DB.Create(cat).Error
}

// 查询类型列表
func (t *Transaction) GetCategories() ([]models.Category, error) {
	var cats []models.Category
	err := models.DB.Where("is_deleted = false").Order("sort_order ASC, updated_at DESC").Find(&cats).Error
	return cats, err
}

// 根据ID查询类型
func (t *Transaction) GetCategoryByID(id string) (*models.Category, error) {
	var cat models.Category
	if err := models.DB.Where("id = ? AND is_deleted = false", id).First(&cat).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

// 更新类型
func (t *Transaction) UpdateCategory(id string, updates *models.Category) (*models.Category, error) {
	var cat models.Category
	if err := models.DB.Where("id = ? AND is_deleted = false", id).First(&cat).Error; err != nil {
		return nil, err
	}

	updates.ID = cat.ID
	if err := models.DB.Model(&cat).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &cat, nil
}

// 删除类型（软删除）
func (t *Transaction) DeleteCategory(id string) error {
	return models.DB.Model(&models.Category{}).Where("id = ?", id).Update("is_deleted", true).Error
}

// =======收支对象服务========

// 创建收支对象
func (t *Transaction) CreatePayee(payee *models.Payee) error {
	payee.ID = uuid.New().String()
	return models.DB.Create(payee).Error
}

// 查询收支对象列表
func (t *Transaction) GetPayees() ([]models.Payee, error) {
	var payees []models.Payee
	err := models.DB.Where("is_deleted = false").Order("updated_at DESC").Find(&payees).Error
	return payees, err
}

// 根据ID查询收支对象
func (t *Transaction) GetPayeeByID(id string) (*models.Payee, error) {
	var payee models.Payee
	if err := models.DB.Where("id = ? AND is_deleted = false", id).First(&payee).Error; err != nil {
		return nil, err
	}
	return &payee, nil
}

// 更新收支对象
func (t *Transaction) UpdatePayee(id string, updates *models.Payee) (*models.Payee, error) {
	var payee models.Payee
	if err := models.DB.Where("id = ? AND is_deleted = false", id).First(&payee).Error; err != nil {
		return nil, err
	}
	updates.ID = payee.ID
	if err := models.DB.Model(&payee).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &payee, nil
}

// 删除收支对象（软删除）
func (t *Transaction) DeletePayee(id string) error {
	return models.DB.Model(&models.Payee{}).Where("id = ?", id).Update("is_deleted", true).Error
}

// =======账户服务========

// 创建账户
func (t *Transaction) CreateAccount(account *models.Account) error {
	account.ID = uuid.New().String()
	return models.DB.Create(account).Error
}

// 查询账户列表
func (t *Transaction) GetAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := models.DB.Where("is_deleted = false").Order("updated_at DESC").Find(&accounts).Error
	return accounts, err
}

// 根据ID查询账户
func (t *Transaction) GetAccountByID(id string) (*models.Account, error) {
	var account models.Account
	if err := models.DB.Where("id = ? AND is_deleted = false", id).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

// 更新账户
func (t *Transaction) UpdateAccount(id string, updates *models.Account) (*models.Account, error) {
	var account models.Account
	if err := models.DB.Where("id = ? AND is_deleted = false", id).First(&account).Error; err != nil {
		return nil, err
	}
	updates.ID = account.ID
	if err := models.DB.Model(&account).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

// 删除账户（软删除）
func (t *Transaction) DeleteAccount(id string) error {
	return models.DB.Model(&models.Account{}).Where("id = ?", id).Update("is_deleted", true).Error
}
