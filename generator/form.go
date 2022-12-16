package generator

import (
	"lemocoder/model"
	"strings"
)

const TABLE_ITEM_OPT_KEY_UPDATE = "update"

const FORM_COMPONENT_TEXT = "ProFormText"
const FORM_COMPONENT_TEXT_AREA = "ProFormTextArea"
const FORM_COMPONENT_DIGIT = "ProFormDigit"
const FORM_COMPONENT_SELECT = "ProFormSelect"

func getDbTypeComponentsMap() map[string]string {
	return map[string]string{
		TYPE_DB_BIGINT:   FORM_COMPONENT_DIGIT,
		TYPE_DB_FLOAT:    FORM_COMPONENT_DIGIT,
		TYPE_DB_INT:      FORM_COMPONENT_SELECT,
		TYPE_DB_SMALLINT: FORM_COMPONENT_DIGIT,
		TYPE_DB_STRING:   FORM_COMPONENT_TEXT,
		TYPE_DB_TEXT:     FORM_COMPONENT_TEXT_AREA,
	}
}

func GetCreateForm(fields []model.TableItemSchema, postUrl string) model.ModalFormSchema {
	form := model.ModalFormSchema{Key: "create", Button: model.ButtonSchema{Title: "添加", Type: "primary"}}
	form.Form = getFormSchema(fields, postUrl, false)
	return form
}

func GetUpdateForm(fields []model.TableItemSchema, postUrl string) model.ModalFormSchema {
	form := model.ModalFormSchema{Key: TABLE_ITEM_OPT_KEY_UPDATE, Button: model.ButtonSchema{Title: "编辑", Type: "primary"}}
	form.Form = getFormSchema(fields, postUrl, false)
	return form
}

func getFormSchema(fields []model.TableItemSchema, postUrl string, isUpdate bool) model.FormSchema {
	var formFields []model.FormFieldSchema
	title := "新增数据"
	if isUpdate {
		title = "更新数据"
		formFields = append(formFields, model.FormFieldSchema{Name: "ID", Component: FORM_COMPONENT_TEXT})
	}
	fieldComponentsMap := getDbTypeComponentsMap()
	for _, field := range fields {
		fdpType := strings.ToUpper(field.DataType)
		fcomponentName, ok := fieldComponentsMap[fdpType]
		if !ok {
			fcomponentName = FORM_COMPONENT_TEXT
		}
		formFields = append(formFields, model.FormFieldSchema{Component: fcomponentName, Name: field.DataName, Label: field.Title})
	}
	return model.FormSchema{Title: title, SubmitUrl: postUrl, FormFields: formFields}
}
