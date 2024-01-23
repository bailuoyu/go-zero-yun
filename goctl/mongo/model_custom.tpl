package model

{{if .Cache}}import (
    "github.com/zeromicro/go-zero/core/stores/cache"
    monc "go-zero-yun/pkg/monckit"
){{else}}import mon "go-zero-yun/pkg/monkit"{{end}}

{{if .Easy}}
const {{.Type}}CollectionName = "{{.snakeType}}"
{{end}}

var _ {{.Type}}Model = (*custom{{.Type}}Model)(nil)

type (
    // {{.Type}}Model is an interface to be customized, add more methods here,
    // and implement the added methods in custom{{.Type}}Model.
    {{.Type}}Model interface {
        {{.lowerType}}Model
    }

    custom{{.Type}}Model struct {
        *default{{.Type}}Model
    }
)


// New{{.Type}}Model returns a model for the mongo.
{{if .Easy}}func New{{.Type}}Model({{if .Cache}}, c cache.CacheConf{{end}}) {{.Type}}Model {
    db := MonDb()
    conn := {{if .Cache}}monc{{else}}mon{{end}}.MustNewModel(db, {{.Type}}CollectionName{{if .Cache}}, c{{end}})
    return &custom{{.Type}}Model{
        default{{.Type}}Model: newDefault{{.Type}}Model(conn),
    }
}{{else}}func New{{.Type}}Model(collection string{{if .Cache}}, c cache.CacheConf{{end}}) {{.Type}}Model {
    db := MonDb()
    conn := {{if .Cache}}monc{{else}}mon{{end}}.MustNewModel(db, collection{{if .Cache}}, c{{end}})
    return &custom{{.Type}}Model{
        default{{.Type}}Model: newDefault{{.Type}}Model(conn),
    }
}{{end}}
