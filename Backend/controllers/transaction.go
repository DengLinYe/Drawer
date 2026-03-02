package controllers

import (
	"Drawer/models"
	"Drawer/services"
	"Drawer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
}

// 创建交易记录
func (t *Transaction) CreateTransaction(c *gin.Context) {
	var tx models.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	if err := services.Service.CreateTransaction(&tx); err != nil {
		utils.ErrorDefault(c, err.Error())
		return
	}

	utils.Success(c, tx)
}

// 查询交易记录列表
func (t *Transaction) QueryTransactions(c *gin.Context) {
	var param services.TransactionQueryParam
	if err := c.ShouldBindQuery(&param); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "查询参数解析失败")
		return
	}

	result, err := services.Service.QueryTransactions(param)
	if err != nil {
		utils.ErrorDefault(c, "查询交易记录失败")
		return
	}

	utils.Success(c, result)
}

// 更新交易记录
func (t *Transaction) UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var updates models.Transaction
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	tx, err := services.Service.UpdateTransaction(id, &updates)
	if err != nil {
		utils.ErrorDefault(c, err.Error())
		return
	}

	utils.Success(c, tx)
}

// 删除交易记录
func (t *Transaction) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	if err := services.Service.DeleteTransaction(id); err != nil {
		utils.ErrorDefault(c, "删除交易记录失败")
		return
	}

	utils.Success(c, map[string]string{"id": id})
}

// 创建类型
func (t *Transaction) CreateCategory(c *gin.Context) {
	var cat models.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	if err := services.Service.CreateCategory(&cat); err != nil {
		utils.ErrorDefault(c, "创建类型失败")
		return
	}

	utils.Success(c, cat)
}

// 查询类型列表
func (t *Transaction) GetCategories(c *gin.Context) {
	cats, err := services.Service.GetCategories()
	if err != nil {
		utils.ErrorDefault(c, "获取类型列表失败")
		return
	}

	utils.Success(c, cats)
}

// 根据ID查询类型
func (t *Transaction) GetCategory(c *gin.Context) {
	id := c.Param("id")
	cat, err := services.Service.GetCategoryByID(id)
	if err != nil {
		utils.ErrorDefault(c, "分类不存在或已被删除")
		return
	}
	utils.Success(c, cat)
}

// 更新类型
func (t *Transaction) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var updates models.Category
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	cat, err := services.Service.UpdateCategory(id, &updates)
	if err != nil {
		utils.ErrorDefault(c, "更新类型失败")
		return
	}

	utils.Success(c, cat)
}

// 删除类型
func (t *Transaction) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := services.Service.DeleteCategory(id); err != nil {
		utils.ErrorDefault(c, "删除类型失败")
		return
	}

	utils.Success(c, map[string]string{"id": id})
}

// 创建收支对象
func (t *Transaction) CreatePayee(c *gin.Context) {
	var payee models.Payee
	if err := c.ShouldBindJSON(&payee); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	if err := services.Service.CreatePayee(&payee); err != nil {
		utils.ErrorDefault(c, "创建收支对象失败")
		return
	}

	utils.Success(c, payee)
}

// 查询收支对象列表
func (t *Transaction) GetPayees(c *gin.Context) {
	payees, err := services.Service.GetPayees()
	if err != nil {
		utils.ErrorDefault(c, "获取收支对象列表失败")
		return
	}

	utils.Success(c, payees)
}

// 根据ID查询收支对象
func (t *Transaction) GetPayee(c *gin.Context) {
	id := c.Param("id")
	payee, err := services.Service.GetPayeeByID(id)
	if err != nil {
		utils.ErrorDefault(c, "收支对象不存在或已被删除")
		return
	}
	utils.Success(c, payee)
}

// 更新收支对象
func (t *Transaction) UpdatePayee(c *gin.Context) {
	id := c.Param("id")
	var updates models.Payee
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	payee, err := services.Service.UpdatePayee(id, &updates)
	if err != nil {
		utils.ErrorDefault(c, "更新收支对象失败")
		return
	}

	utils.Success(c, payee)
}

// 删除收支对象
func (t *Transaction) DeletePayee(c *gin.Context) {
	id := c.Param("id")
	if err := services.Service.DeletePayee(id); err != nil {
		utils.ErrorDefault(c, "删除收支对象失败")
		return
	}

	utils.Success(c, map[string]string{"id": id})
}

// 创建支付方式
func (t *Transaction) CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	if err := services.Service.CreateAccount(&account); err != nil {
		utils.ErrorDefault(c, "创建账户失败")
		return
	}

	utils.Success(c, account)
}

// 查询支付方式列表
func (t *Transaction) GetAccounts(c *gin.Context) {
	accounts, err := services.Service.GetAccounts()
	if err != nil {
		utils.ErrorDefault(c, "获取账户列表失败")
		return
	}

	utils.Success(c, accounts)
}

// 根据ID查询支付方式
func (t *Transaction) GetAccount(c *gin.Context) {
	id := c.Param("id")
	account, err := services.Service.GetAccountByID(id)
	if err != nil {
		utils.ErrorDefault(c, "账户不存在或已被删除")
		return
	}
	utils.Success(c, account)
}

// 更新支付方式列表
func (t *Transaction) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var updates models.Account
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.Error(c, http.StatusBadRequest, 400, "参数解析失败")
		return
	}

	account, err := services.Service.UpdateAccount(id, &updates)
	if err != nil {
		utils.ErrorDefault(c, "更新账户失败")
		return
	}

	utils.Success(c, account)
}

// 删除支付方式
func (t *Transaction) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := services.Service.DeleteAccount(id); err != nil {
		utils.ErrorDefault(c, "删除账户失败")
		return
	}

	utils.Success(c, map[string]string{"id": id})
}
