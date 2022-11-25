package database

func getAllTables() []interface{} {
	return []interface{}{
		new(User),
		new(DataTable),
		new(Project),
		new(WebPage),
		// Code generated Begin; DO NOT EDIT.
		// Code generated End; DO NOT EDIT.
	}
}
