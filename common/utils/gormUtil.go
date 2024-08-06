package utils

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

/**
  SELECT
*/

// WithSelect 指定要選擇的字段
func WithSelect(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	}
}

// WithRawQuery 執行原生 SQL 查詢
func WithRawQuery(query string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Raw(query, args...)
	}
}

/**
  WHERE
*/

// WithRawCondition 應用自定義原生 SQL 條件
// example:
// rawCondition := "age > ? AND name LIKE ?"
// args := []interface{}{25, "%a%"}
func WithRawCondition(condition string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(condition, args...)
	}
}

// WithOmit 排除結果中的特定字段
func WithOmit(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(fields...)
	}
}

// WithSoftDelete 應用軟刪除過濾器
func WithSoftDelete() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at IS NULL")
	}
}

// WithDateRange 應用日期範圍過濾器
func WithDateRange(startTime, endTime time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("created_time BETWEEN ? AND ?", startTime, endTime)
	}
}

// WithLike 應用 LIKE 過濾器
func WithLike(column, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" LIKE ?", "%"+value+"%")
	}
}

// WithNotIn 應用 NOT IN 過濾器
func WithNotIn(column string, values []interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Not(column+" IN ?", values)
	}
}

// WithLimit 設置返回記錄的限制
func WithLimit(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit > 0 {
			return db.Limit(limit)
		}
		return db
	}
}

// WithNotNull 過濾指定列不為 NULL 的記錄
func WithNotNull(column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column + " IS NOT NULL")
	}
}

// WithNull 過濾指定列為 NULL 的記錄
func WithNull(column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column + " IS NULL")
	}
}

// WithIn 過濾指定列在給定值範圍內的記錄
func WithIn(column string, values []interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" IN ?", values)
	}
}

// WithOr 在查詢中應用 OR 條件
func WithOr(query string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Or(query, args...)
	}
}

// WithBetween 過濾指定列在兩個值之間的記錄
func WithBetween(column string, min, max interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" BETWEEN ? AND ?", min, max)
	}
}

// WithExists 在查詢中應用 EXISTS 子句
func WithExists(subQuery string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("EXISTS ("+subQuery+")", args...)
	}
}

/**
  UPDATE
*/

// WithUpdates 應用記錄更新
func WithUpdates(updates map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Updates(updates)
	}
}

/**
  JOIN
*/

// WithJoin 添加聯接子句
func WithJoin(joinTable, joinCondition string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(joinTable + " ON " + joinCondition)
	}
}

// WithJoinAndSelect 同時添加聯接和選擇字段
func WithJoinAndSelect(joinTable, joinCondition string, fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(joinTable + " ON " + joinCondition).Select(fields)
	}
}

/**
  GROUP BY
*/

// WithGroupBy 應用 GROUP BY 子句
func WithGroupBy(columns ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) > 0 {
			groupByClause := strings.Join(columns, ", ")
			return db.Group(groupByClause)
		}
		return db
	}
}

/**
  ORDER BY
*/

// WithOrderBySQL 應用自定義 ORDER BY 子句
func WithOrderBySQL(orderSQL string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(orderSQL)
	}
}

// WithSorting 應用多個排序子句
func WithSorting(orderClauses ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(orderClauses) > 0 {
			return db.Order(strings.Join(orderClauses, ", "))
		}
		return db
	}
}

// WithConditionalOrder 僅在條件為真時應用 ORDER BY 子句
func WithConditionalOrder(condition bool, orderClause string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if condition {
			return db.Order(orderClause)
		}
		return db
	}
}

// WithLimitAndOrderBy 同時設置限制和排序
func WithLimitAndOrderBy(limit int, orderClause string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(orderClause).Limit(limit)
	}
}

// WithOffsetAndOrderBy 同時設置偏移量和排序
func WithOffsetAndOrderBy(offset int, orderClause string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(orderClause).Offset(offset)
	}
}

/**
  PAGE
*/

// WithPagination 應用分頁功能，使用頁面和限制
func WithPagination(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

/**
  OTHERS
*/

// WithLimitAndOffset 同時應用限制和偏移量
func WithLimitAndOffset(limit, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit > 0 && offset >= 0 {
			return db.Limit(limit).Offset(offset)
		}
		return db
	}
}

// WithHaving 應用 HAVING 子句
func WithHaving(havingClause string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Having(havingClause)
	}
}

// WithHavingCondition 在 HAVING 子句中添加條件
func WithHavingCondition(condition string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Having(condition, args...)
	}
}

// WithDistinct 應用 DISTINCT 子句
func WithDistinct() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Distinct()
	}
}

// WithSelectDistinct 選擇特定字段並使用 DISTINCT
func WithSelectDistinct(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields).Distinct()
	}
}

/**
  动态拼接
*/
//MAP 空直判断 忽略空值 (构造用指针)
func BuildNotNullMap(obj interface{}) map[string]interface{} {
	insertMap := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return insertMap
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		// Check for nil pointers
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}

		gormTag := fieldType.Tag.Get("gorm")
		columnName := parseColumnName(gormTag)
		if columnName == "" {
			columnName = fieldType.Name
		}

		// Dereference pointer types
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		insertMap[columnName] = fieldValue.Interface()
	}

	return insertMap
}

// BuildNullMap 包含指针空值处理的 map (构造用指针)
func BuildNullMap(obj interface{}) map[string]interface{} {
	insertMap := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return insertMap
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		// 解析 gorm 标签获取列名
		gormTag := fieldType.Tag.Get("gorm")
		columnName := parseColumnName(gormTag)
		if columnName == "" {
			columnName = fieldType.Name
		}

		// 检查指针类型字段并处理空值
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			insertMap[columnName] = nil
		} else {
			// 处理指针类型字段的解引用
			if fieldValue.Kind() == reflect.Ptr {
				fieldValue = fieldValue.Elem()
			}
			insertMap[columnName] = fieldValue.Interface()
		}
	}

	return insertMap
}

// 返回构造中非空字段的名称和值 #预设值也会加进条件(构造不使用指针)
func GetNonEmptyFields(obj interface{}) map[string]interface{} {
	fields := map[string]interface{}{}
	refValue := reflect.ValueOf(obj)
	refType := reflect.TypeOf(obj)

	if refValue.Kind() != reflect.Struct {
		panic("obj must be a struct")
	}

	for i := 0; i < refValue.NumField(); i++ {
		field := refType.Field(i)
		value := refValue.Field(i)

		// 获取字段的gorm tag
		gormTag := field.Tag.Get("gorm")
		if gormTag == "" {
			continue // 跳过没有gorm tag字段
		}

		columnName := field.Tag.Get("column")
		if columnName == "" {
			continue // 跳过没有column tag字段
		}

		// 根据字段类型设置条件
		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if value.Int() != 0 {
				fields[columnName] = value.Interface()
			}
		case reflect.String:
			if value.String() != "" {
				fields[columnName] = value.Interface()
			}
		case reflect.Float32, reflect.Float64:
			if value.Float() != 0 {
				fields[columnName] = value.Interface()
			}
		case reflect.Bool:
			fields[columnName] = value.Interface()
			// 添加其他类型的处理
		}
	}

	fmt.Println(fields)

	return fields
}

// WhereConditions (构造不使用指针)
func WhereConditions(obj interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		refValue := reflect.ValueOf(obj)
		refType := reflect.TypeOf(obj)

		if refValue.Kind() != reflect.Struct {
			panic("obj must be a struct")
		}

		for i := 0; i < refValue.NumField(); i++ {
			field := refType.Field(i)
			value := refValue.Field(i)

			// 获取字段的gorm tag
			gormTag := field.Tag.Get("gorm")
			if gormTag == "" {
				continue // 跳过没有gorm tag字段?
			}

			columnName := field.Tag.Get("column")
			if columnName == "" {
				continue // 跳过没有column tag字段
			}

			switch value.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if value.Int() != 0 {
					db = db.Where(columnName+" = ?", value.Interface())
				}
			case reflect.String:
				if value.String() != "" {
					db = db.Where(columnName+" = ?", value.Interface())
				}
			case reflect.Struct:
				if field.Type == reflect.TypeOf(time.Time{}) {
					timeValue := value.Interface().(time.Time)
					if !timeValue.IsZero() {
						db = db.Where(columnName+" = ?", value.Interface())
					}
				}
			}
		}

		return db
	}

}

// Criterion 包含查询条件的结构体
//
//		criteria := []Criterion{
//		  {Condition: "age", Operation: "gt", Value: 18},
//		  {Condition: "name", Operation: "eq", Value: "John"},
//		  {Condition: "created_at", Operation: "between", Value: startTime, SecondValue: endTime},
//		  {Condition: "role", Operation: "in", ListValue: []interface{}{"admin", "user"}},
//	    }
type Criterion struct {
	Condition   string        // 查询条件字服串 例如 "age > ?"
	Operation   string        // 操作符例如 "eq", "gt", "in"
	Value       interface{}   // 值 可以式单个值或值的集合
	SecondValue interface{}   // 第二个直 用于 betweenValue 操作
	ListValue   []interface{} // 值的列表 用于 listValue 操作
}

// BuildDynamicQuery 根据传入得 Criterion 结构体动态构建查询条件
func BuildDynamicQuery(db *gorm.DB, criteria []Criterion) *gorm.DB {
	for _, criterion := range criteria {
		switch criterion.Operation {
		case "eq":
			db = db.Where(criterion.Condition, criterion.Value)
		case "gt":
			db = db.Where(criterion.Condition+" > ?", criterion.Value)
		case "lt":
			db = db.Where(criterion.Condition+" < ?", criterion.Value)
		case "between":
			db = db.Where(criterion.Condition+" BETWEEN ? AND ?", criterion.Value, criterion.SecondValue)
		// 处理直为列表的状态
		case "in":
			// ?????????
			db = db.Where(criterion.Condition+" IN (?)", criterion.ListValue)
		default:
			// 其他操作符的处理
		}
	}
	return db
}

// parseColumnName 解析 GORM 标签中的列名
func parseColumnName(tag string) string {
	parts := strings.Split(tag, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "column:") {
			return strings.TrimPrefix(part, "column:")
		}
	}
	return ""
}
