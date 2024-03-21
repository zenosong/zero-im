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

func newAdmin(db *gorm.DB, opts ...gen.DOOption) admin {
	_admin := admin{}

	_admin.adminDo.UseDB(db, opts...)
	_admin.adminDo.UseModel(&model.Admin{})

	tableName := _admin.adminDo.TableName()
	_admin.ALL = field.NewAsterisk(tableName)
	_admin.ID = field.NewInt64(tableName, "id")
	_admin.CreatedAt = field.NewTime(tableName, "created_at")
	_admin.UpdatedAt = field.NewTime(tableName, "updated_at")
	_admin.DeletedAt = field.NewField(tableName, "deleted_at")
	_admin.Username = field.NewString(tableName, "username")
	_admin.Password = field.NewString(tableName, "password")
	_admin.APIToken = field.NewString(tableName, "api_token")
	_admin.Avatar = field.NewString(tableName, "avatar")
	_admin.GroupID = field.NewInt64(tableName, "group_id")
	_admin.IsSuper = field.NewBool(tableName, "is_super")

	_admin.fillFieldMap()

	return _admin
}

type admin struct {
	adminDo adminDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Username  field.String
	Password  field.String
	APIToken  field.String
	Avatar    field.String
	GroupID   field.Int64
	IsSuper   field.Bool

	fieldMap map[string]field.Expr
}

func (a admin) Table(newTableName string) *admin {
	a.adminDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a admin) As(alias string) *admin {
	a.adminDo.DO = *(a.adminDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *admin) updateTableName(table string) *admin {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Username = field.NewString(table, "username")
	a.Password = field.NewString(table, "password")
	a.APIToken = field.NewString(table, "api_token")
	a.Avatar = field.NewString(table, "avatar")
	a.GroupID = field.NewInt64(table, "group_id")
	a.IsSuper = field.NewBool(table, "is_super")

	a.fillFieldMap()

	return a
}

func (a *admin) WithContext(ctx context.Context) *adminDo { return a.adminDo.WithContext(ctx) }

func (a admin) TableName() string { return a.adminDo.TableName() }

func (a admin) Alias() string { return a.adminDo.Alias() }

func (a admin) Columns(cols ...field.Expr) gen.Columns { return a.adminDo.Columns(cols...) }

func (a *admin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *admin) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["username"] = a.Username
	a.fieldMap["password"] = a.Password
	a.fieldMap["api_token"] = a.APIToken
	a.fieldMap["avatar"] = a.Avatar
	a.fieldMap["group_id"] = a.GroupID
	a.fieldMap["is_super"] = a.IsSuper
}

func (a admin) clone(db *gorm.DB) admin {
	a.adminDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a admin) replaceDB(db *gorm.DB) admin {
	a.adminDo.ReplaceDB(db)
	return a
}

type adminDo struct{ gen.DO }

func (a adminDo) Debug() *adminDo {
	return a.withDO(a.DO.Debug())
}

func (a adminDo) WithContext(ctx context.Context) *adminDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminDo) ReadDB() *adminDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminDo) WriteDB() *adminDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminDo) Session(config *gorm.Session) *adminDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminDo) Clauses(conds ...clause.Expression) *adminDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminDo) Returning(value interface{}, columns ...string) *adminDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminDo) Not(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminDo) Or(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminDo) Select(conds ...field.Expr) *adminDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminDo) Where(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminDo) Order(conds ...field.Expr) *adminDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminDo) Distinct(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminDo) Omit(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminDo) Join(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminDo) LeftJoin(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminDo) RightJoin(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminDo) Group(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminDo) Having(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminDo) Limit(limit int) *adminDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminDo) Offset(offset int) *adminDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *adminDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminDo) Unscoped() *adminDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminDo) Create(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminDo) CreateInBatches(values []*model.Admin, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminDo) Save(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminDo) First() (*model.Admin, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Take() (*model.Admin, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Last() (*model.Admin, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Find() ([]*model.Admin, error) {
	result, err := a.DO.Find()
	return result.([]*model.Admin), err
}

func (a adminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Admin, err error) {
	buf := make([]*model.Admin, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminDo) FindInBatches(result *[]*model.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminDo) Attrs(attrs ...field.AssignExpr) *adminDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminDo) Assign(attrs ...field.AssignExpr) *adminDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminDo) Joins(fields ...field.RelationField) *adminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminDo) Preload(fields ...field.RelationField) *adminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminDo) FirstOrInit() (*model.Admin, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FirstOrCreate() (*model.Admin, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FindByPage(offset int, limit int) (result []*model.Admin, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminDo) Delete(models ...*model.Admin) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminDo) withDO(do gen.Dao) *adminDo {
	a.DO = *do.(*gen.DO)
	return a
}
