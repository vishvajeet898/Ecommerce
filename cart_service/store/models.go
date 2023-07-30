package store

type QueryFilter struct {
	Rows     string
	Table    string
	From     string
	JoinType string
	Join     string
	JoinOn   string
	Where    string
}
