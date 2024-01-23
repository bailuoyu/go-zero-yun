package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}{{$table := .}}
// {{Mapper $table.Name}} {{$table.Comment}}
type {{Mapper $table.Name}} struct {
{{range $table.ColumnsSeq}}{{$col := $table.GetColumn .}}   {{Mapper $col.Name}}    {{Type $col}}   `{{eTag $col}} json:"{{$col.Name}}"`    // {{eRemark $table $col}}
{{end}}
}

// TableName 表名
func (t {{Mapper $table.Name}}) TableName() string {
	return "{{$table.Name}}"
}
{{end}}
