package model

const TABLE_ITEM_OPT_FORM = "form"
const TABLE_ITEM_OPT_ACTION = "action"
const TABLE_ITEM_OPT_REDIRECT = "redirect"
const TABLE_ITEM_OPT_EDIT = "edit"

// TableItemSchema 数据表格行数据结构描述
//
// ValueType和搜索表单: ProTable 会根据列来生成一个 Form，用于筛选列表数据。
// Form 的列是根据 valueType 来生成不同的类型,详细的值类型请查看通用配置。
// valueType 是 ProComponents 的灵魂，ProComponents 会根据 valueType 来映射成不同的表单项。
// https://procomponents.ant.design/components/schema#valuetype
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
	Searchable                                                                                 bool
	Title, ItemDataTypeName, ItemsDataUrl, ItemUpdateUrl, ItemDeleteUrl, ItemCreateUrl, RowKey string
	Items                                                                                      []TableItemSchema
	ItemOptions                                                                                []TableItemOptionSchema
	ItemForms                                                                                  []ModalFormSchema
	ToolBarForms                                                                               []ModalFormSchema
	BatchOptButtons                                                                            []BatchOptButtonSchema
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
