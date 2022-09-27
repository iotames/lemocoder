package generator

type TableItemSchema struct {
	Order                                                                   int64
	ColSize                                                                 float64
	DataName, DataType, Title, ValueType                                    string
	Editable, Copyable, Ellipsis, Sorter, Search, HideInSearch, HideInTable bool
}

type TableItemOptionSchema struct {
	Key, Title, Type, Url string // type edit,act,redirect
}

type TableSchema struct {
	ItemDataTypeName, ItemsDataUrl, ItemUpdateUrl, ItemDeleteUrl, ItemCreateUrl string
	Items                                                                       []TableItemSchema
	ItemOptions                                                                 []TableItemOptionSchema
	ItemForms                                                                   []ModalFormSchema
	ToolBarForms                                                                []ModalFormSchema
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
