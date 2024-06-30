package graft

import (
	_ "database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	_ "reflect"
	_ "strings"
)

type Graft struct {
	*pgxpool.Pool
}

func NewGraft() *Graft {
	db, err := New()
	if err != nil {
		panic("db connection issue")
	}
	return db
}

func (g *Graft) create(model interface{}) {
	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model)

	if t.Kind() != reflect.Struct {
		fmt.Println("Create expects a struct")
		return
	}

	var columns []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		columns = append(columns, t.Field(i).Name)
		values = append(values, v.Field(i).Interface())
	}

	query := fmt.Sprintf("INSERT INTO %s")
	stmt, err := g.db.Prepare(query)
	if err != nil {
		fmt.Println("Failed to prepare Query")
		return
	}
	defer stmt.Close()


}

func (g *Graft) update(map[string]any) {

}

func (g *Graft) delete(map[string]any) {

}

func (g *Graft) edit(map[string]any) {

}

