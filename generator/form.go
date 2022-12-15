package generator

import (
	"lemocoder/model"
)

const TABLE_ITEM_OPT_KEY_UPDATE = "update"

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
	title := "新增数据"
	if isUpdate {
		title = "更新数据"
	}
	var formFields []model.FormFieldSchema
	form := model.FormSchema{Title: title, SubmitUrl: postUrl}
	for _, field := range fields {
		// TODO ADD formFields
		switch field.DataType {
		case TYPE_DB_STRING:
			formFields = append(formFields, model.FormFieldSchema{Component: ""})
		case TYPE_DB_TEXT:
			formFields = append(formFields, model.FormFieldSchema{Component: ""})
		}
	}
	form.FormFields = formFields
	return form
}
