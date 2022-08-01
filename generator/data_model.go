package generator

type TableItemSchema struct {
	Order                                                                   int64
	ColSize                                                                 float64
	DataName, DataType, Title, ValueType                                    string
	Editable, Copyable, Ellipsis, Sorter, Search, HideInSearch, HideInTable bool
}

type TableSchema struct {
	ItemDataTypeName, ItemsDataUrl, ItemUpdateUrl string
	Items                                         []TableItemSchema
}
