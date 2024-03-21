// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"zero/rpc/dao/model"
)

func newChatSetting(db *gorm.DB, opts ...gen.DOOption) chatSetting {
	_chatSetting := chatSetting{}

	_chatSetting.chatSettingDo.UseDB(db, opts...)
	_chatSetting.chatSettingDo.UseModel(&model.ChatSetting{})

	tableName := _chatSetting.chatSettingDo.TableName()
	_chatSetting.ALL = field.NewAsterisk(tableName)
	_chatSetting.ID = field.NewInt64(tableName, "id")
	_chatSetting.Name = field.NewString(tableName, "name")
	_chatSetting.Title = field.NewString(tableName, "title")
	_chatSetting.GroupID = field.NewInt64(tableName, "group_id")
	_chatSetting.Value = field.NewString(tableName, "value")
	_chatSetting.Options = field.NewString(tableName, "options")
	_chatSetting.Type = field.NewString(tableName, "type")
	_chatSetting.CreatedAt = field.NewTime(tableName, "created_at")
	_chatSetting.UpdatedAt = field.NewTime(tableName, "updated_at")

	_chatSetting.fillFieldMap()

	return _chatSetting
}

type chatSetting struct {
	chatSettingDo chatSettingDo

	ALL       field.Asterisk
	ID        field.Int64
	Name      field.String
	Title     field.String
	GroupID   field.Int64
	Value     field.String
	Options   field.String
	Type      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (c chatSetting) Table(newTableName string) *chatSetting {
	c.chatSettingDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chatSetting) As(alias string) *chatSetting {
	c.chatSettingDo.DO = *(c.chatSettingDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chatSetting) updateTableName(table string) *chatSetting {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Name = field.NewString(table, "name")
	c.Title = field.NewString(table, "title")
	c.GroupID = field.NewInt64(table, "group_id")
	c.Value = field.NewString(table, "value")
	c.Options = field.NewString(table, "options")
	c.Type = field.NewString(table, "type")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *chatSetting) WithContext(ctx context.Context) *chatSettingDo {
	return c.chatSettingDo.WithContext(ctx)
}

func (c chatSetting) TableName() string { return c.chatSettingDo.TableName() }

func (c chatSetting) Alias() string { return c.chatSettingDo.Alias() }

func (c chatSetting) Columns(cols ...field.Expr) gen.Columns { return c.chatSettingDo.Columns(cols...) }

func (c *chatSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chatSetting) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["title"] = c.Title
	c.fieldMap["group_id"] = c.GroupID
	c.fieldMap["value"] = c.Value
	c.fieldMap["options"] = c.Options
	c.fieldMap["type"] = c.Type
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c chatSetting) clone(db *gorm.DB) chatSetting {
	c.chatSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chatSetting) replaceDB(db *gorm.DB) chatSetting {
	c.chatSettingDo.ReplaceDB(db)
	return c
}

type chatSettingDo struct{ gen.DO }

func (c chatSettingDo) Debug() *chatSettingDo {
	return c.withDO(c.DO.Debug())
}

func (c chatSettingDo) WithContext(ctx context.Context) *chatSettingDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chatSettingDo) ReadDB() *chatSettingDo {
	return c.Clauses(dbresolver.Read)
}

func (c chatSettingDo) WriteDB() *chatSettingDo {
	return c.Clauses(dbresolver.Write)
}

func (c chatSettingDo) Session(config *gorm.Session) *chatSettingDo {
	return c.withDO(c.DO.Session(config))
}

func (c chatSettingDo) Clauses(conds ...clause.Expression) *chatSettingDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chatSettingDo) Returning(value interface{}, columns ...string) *chatSettingDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chatSettingDo) Not(conds ...gen.Condition) *chatSettingDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chatSettingDo) Or(conds ...gen.Condition) *chatSettingDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chatSettingDo) Select(conds ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chatSettingDo) Where(conds ...gen.Condition) *chatSettingDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chatSettingDo) Order(conds ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chatSettingDo) Distinct(cols ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chatSettingDo) Omit(cols ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chatSettingDo) Join(table schema.Tabler, on ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chatSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chatSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chatSettingDo) Group(cols ...field.Expr) *chatSettingDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chatSettingDo) Having(conds ...gen.Condition) *chatSettingDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chatSettingDo) Limit(limit int) *chatSettingDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chatSettingDo) Offset(offset int) *chatSettingDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chatSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *chatSettingDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chatSettingDo) Unscoped() *chatSettingDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chatSettingDo) Create(values ...*model.ChatSetting) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chatSettingDo) CreateInBatches(values []*model.ChatSetting, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chatSettingDo) Save(values ...*model.ChatSetting) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chatSettingDo) First() (*model.ChatSetting, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatSetting), nil
	}
}

func (c chatSettingDo) Take() (*model.ChatSetting, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatSetting), nil
	}
}

func (c chatSettingDo) Last() (*model.ChatSetting, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatSetting), nil
	}
}

func (c chatSettingDo) Find() ([]*model.ChatSetting, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChatSetting), err
}

func (c chatSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChatSetting, err error) {
	buf := make([]*model.ChatSetting, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chatSettingDo) FindInBatches(result *[]*model.ChatSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chatSettingDo) Attrs(attrs ...field.AssignExpr) *chatSettingDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chatSettingDo) Assign(attrs ...field.AssignExpr) *chatSettingDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chatSettingDo) Joins(fields ...field.RelationField) *chatSettingDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chatSettingDo) Preload(fields ...field.RelationField) *chatSettingDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chatSettingDo) FirstOrInit() (*model.ChatSetting, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatSetting), nil
	}
}

func (c chatSettingDo) FirstOrCreate() (*model.ChatSetting, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatSetting), nil
	}
}

func (c chatSettingDo) FindByPage(offset int, limit int) (result []*model.ChatSetting, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c chatSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chatSettingDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chatSettingDo) Delete(models ...*model.ChatSetting) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chatSettingDo) withDO(do gen.Dao) *chatSettingDo {
	c.DO = *do.(*gen.DO)
	return c
}
