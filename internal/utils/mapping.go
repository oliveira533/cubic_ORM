package utils

import (
	"reflect"
	"strings"

	"github.com/oliveira533/cubic_ORM.git/internal/db"
)

func MappingStruct(model any) ([]db.FieldInfo, string) {
	modelType := reflect.TypeOf(model)

	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	tableName := strings.ToLower(modelType.Name())

	var fields []db.FieldInfo

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		fields = append(fields, db.FieldInfo{
			Name:       field.Name,
			ColumnName: strings.ToLower(field.Name),
			MataFields: extractMetaFields(field.Tag.Get("cubic")),
			Type:       field.Type.String(),
		})
	}

	return fields, tableName
}

func extractMetaFields(tag string) []db.MetaField {
	slice := strings.Split(tag, ",")
	var metaFields []db.MetaField
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], "=") {
			aux := strings.Split(slice[i], "=")
			value := any(aux[1])
			metaFields = append(metaFields, db.MetaField{
				Title: aux[0],
				Value: &value,
			})
			continue
		}
		metaFields = append(metaFields, db.MetaField{
			Title: slice[i],
		})
	}

	return metaFields
}
