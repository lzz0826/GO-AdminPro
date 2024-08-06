package utils

//
//import "gorm.io/gorm"
//
//type repository struct {
//	readDB  *gorm.DB
//	writeDB *gorm.DB
//
//	cacheRepo iface.ICacheRepository
//	mongoRepo iface.IMongoRepository
//	jPush     iface.IPushRepository
//}
//
//// Get 取得的資訊
//func (repo *repository) Get(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
//	if tx == nil {
//		tx = repo.readDB.WithContext(ctx)
//	}
//	tx = tx.Scopes(scopes...)
//	tx = opt.Preload(tx)
//	err := tx.Table(model.TableName()).Scopes(opt.Where).First(model).Error
//	if err != nil {
//		return errors.ConvertMySQLError(err)
//	}
//	return nil
//}
