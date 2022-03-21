package repositories

type IBaseRepository[TEntity interface{}] interface {
	GetById(id int) *TEntity
	Create(entity *TEntity) *TEntity
	Update(entity *TEntity) *TEntity
	Delete(id int) *TEntity
}
