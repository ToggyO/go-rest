package repositories

type IBaseRepository[TEntity interface{}] interface {
	GetById(id int) *TEntity
	Create(model *TEntity) *TEntity
	Update(model *TEntity) *TEntity
	Delete(id int) *TEntity
}
