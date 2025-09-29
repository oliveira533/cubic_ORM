package utils

import (
	"reflect"
	"strings"
)

type FieldInfo struct {
	Name       string
	ColumnName string
	MataFields []MetaField
	Type       string
}

type MetaField struct {
	Title string
	Value *any
}

func MappingStruct(model any) ([]FieldInfo, string) {
	modelType := reflect.TypeOf(model)

	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	tableName := strings.ToLower(modelType.Name())

	var fields []FieldInfo

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		fields = append(fields, FieldInfo{
			Name:       field.Name,
			ColumnName: strings.ToLower(field.Name),
			MataFields: extractMetaFields(field.Tag.Get("cubic")),
			Type:       field.Type.String(),
		})
	}

	return fields, tableName
}

func extractMetaFields(tag string) []MetaField {
	slice := strings.Split(tag, ",")
	var metaFields []MetaField
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], "=") {
			aux := strings.Split(slice[i], "=")
			value := any(aux[1])
			metaFields = append(metaFields, MetaField{
				Title: aux[0],
				Value: &value,
			})
			continue
		}
		metaFields = append(metaFields, MetaField{
			Title: slice[i],
		})
	}

	return metaFields
}
