package model

const TABLE_ITEM_OPT_FORM = "form"
const TABLE_ITEM_OPT_ACTION = "action"
const TABLE_ITEM_OPT_REDIRECT = "redirect"
const TABLE_ITEM_OPT_EDIT = "edit"

type TableItemSchema struct {
	Width, Order                                                            int64
	ColSize                                                                 float64
	DataName, DataType, Title, ValueType                                    string
	Editable, Copyable, Ellipsis, Sorter, Search, HideInSearch, HideInTable bool
}

type TableItemOptionSchema struct {
	Key, Title, Type, Url string // type: edit,action,form,redirect
}

type TableSchema struct {
	ItemDataTypeName, ItemsDataUrl, ItemUpdateUrl, ItemDeleteUrl, ItemCreateUrl, RowKey string
	Items                                                                               []TableItemSchema
	ItemOptions                                                                         []TableItemOptionSchema
	ItemForms                                                                           []ModalFormSchema
	ToolBarForms                                                                        []ModalFormSchema
	BatchOptButtons                                                                     []BatchOptButtonSchema
}

type BatchOptButtonSchema struct {
	Url, Title string
}

type ButtonSchema struct {
	Title, Type, Size string
}
type FormSchema struct {
	Title, SubmitUrl string
	FormFields       []FormFieldSchema
}
type FormFieldSchema struct {
	Group                                      []FormFieldSchema
	Component, Name, Label, Width, Placeholder string
}

type ModalFormSchema struct {
	Key    string
	Button ButtonSchema
	Form   FormSchema
}
