package catav15

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func ManipulateCreature(db *gorm.DB, column string, id []int, key string, value interface{}) error {
	model := Creature{}
	return manipulate(db, column, id, key, value, model)
}

func ManipulateCreatureTemplate(db *gorm.DB, column string, id []int, key string, value interface{}) error {
	model := CreatureTemplate{}
	return manipulate(db, column, id, key, value, model)
}

func manipulate(db *gorm.DB, column string, id []int, key string, value interface{}, model interface{}) error {
	realType := getFieldType(columnKeyToStructKey(key), model)
	if realType == nil {
		return fmt.Errorf("bad key/column for creature_template: %s", key)
	}

	var typeCastProblem bool

	// Based on the column type...
	switch realType.Kind() {

	// ... if it's a string ...
	case reflect.String:
		// ... try to convert the user's provided type to a string ...
		realValue, OK := value.(string)
		if !OK {
			// ... error if it cannot be converted to that type
			typeCastProblem = true
			break
		}

		// ... update the table based on the id id
		for _, i := range id {
			db.Model(&model).Where(fmt.Sprintf("%s = ?", column), i).UpdateColumn(key, realValue)
		}

	case reflect.Int:
		realValue, OK := value.(int64)
		if !OK {
			typeCastProblem = true
			break
		}

		for _, i := range id {
			db.Model(&model).Where(fmt.Sprintf("%s = ?", column), i).UpdateColumn(key, realValue)
		}

	case reflect.Float64:
		realValue, OK := value.(float64)
		if !OK {
			typeCastProblem = true
			break
		}

		for _, i := range id {
			db.Model(&model).Where(fmt.Sprintf("%s = ?", column), i).UpdateColumn(key, realValue)
		}

	// Error handling case, basically...
	default:
		return fmt.Errorf("unknown data type for key: %s", key)
	}

	if typeCastProblem {
		return fmt.Errorf("id %d: unable to convert %v to required type %s - stopping", id, value, realType.Name())
	}

	return nil
}
