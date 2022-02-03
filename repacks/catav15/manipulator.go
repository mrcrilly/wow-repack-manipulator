package catav15

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func ManipulateCreature(db *gorm.DB, entry int, key string, value interface{}) error {
	return errors.New("not implemented")
}

func ManipulateCreatureTemplate(db *gorm.DB, entry int, key string, value interface{}) error {
	model := CreatureTemplate{}
	return manipulate(db, entry, key, value, model)
}

func manipulate(db *gorm.DB, entry int, key string, value interface{}, model interface{}) error {
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

		// ... other update the table based on the entry id
		db.Model(&model).Where("entry = ?", entry).UpdateColumn(key, realValue)

	case reflect.Int:
		realValue, OK := value.(int64)
		if !OK {
			typeCastProblem = true
			break
		}

		db.Model(&model).Where("entry = ?", entry).UpdateColumn(key, realValue)

	case reflect.Float64:
		realValue, OK := value.(float64)
		if !OK {
			typeCastProblem = true
			break
		}

		db.Model(&model).Where("entry = ?", entry).UpdateColumn(key, realValue)

	// Error handling case, basically...
	default:
		return fmt.Errorf("unknown data type for key: %s", key)
	}

	if typeCastProblem {
		return fmt.Errorf("entry %d: unable to convert %v to required type %s - stopping", entry, value, realType.Name())
	}

	return nil
}
