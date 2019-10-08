package main 

import (
	"fmt"
	"reflect"
	"strings"
)

type XyPayBill1 struct {
	Id              int64           `db:"I_ID"`
	FileName        string          `db:"CH_HISTORY_FILE_NAME"`
}

type XyPayBill2 struct {
	Id              int64           `db:"I_ID"`
	FileName        string          `db:"CH_HISTORY_FILE_NAME"`
}


func main(){
	st := new(XyPayBill1)
	query := BulkInsert(st)
	fmt.Println("query=", query)

}

func BulkInsert(list ...interface{}) (query string) {
	if len(list) == 0 {
		return
	}
	var (
		fields     = getFields(list[0])
		fieldStr   = strings.Join(fields, ",")
		values     []interface{}
		holderStr  = "(?"
		holderList []string
	)
	for i := 1; i < len(fields); i++ {
		holderStr += ",?"
	}
	holderStr += ")"
	for _, v := range list {
		holderList = append(holderList, holderStr)

		structV := New(v)
		structV.TagName = "db"
		for _, f := range structV.Fields() {
			values = append(values, f.Value())
		}
	}

	query = fmt.Sprintf("INSERT INTO (%s) VALUES %s", fieldStr, strings.Join(holderList, ","))
	return
}

type Struct struct {
	raw     interface{}
	value   reflect.Value
	TagName string
}

// New returns a new *Struct with the struct s. It panics if the s's kind is
// not struct.
func New(s interface{}) *Struct {
	return &Struct{
		raw:     s,
		value:   strctVal(s),
		TagName: "structs",
	}
}

func getFields(i interface{}) (fields []string) {
	fields = getFieldsT(reflect.TypeOf(i))
	return
}

func getFieldsT(ty reflect.Type) (fields []string) {
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	if ty.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < ty.NumField(); i++ {
		f := ty.Field(i)
		k := f.Type.Kind()
		if k == reflect.Struct && f.Anonymous {
			fields = append(fields, getFields(f.Type)...)
		} else {
			val, ok := f.Tag.Lookup("db")
			if ok {
				fields = append(fields, val)
			}
		}
	}

	return fields
}

func strctVal(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	return v
}
func (s *Struct) Fields() []*Field {
	return getFields(s.value, s.TagName)
}

func (s *Struct)getFields(v reflect.Value, tagName string) []*Field {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	var fields []*Field

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if tag := field.Tag.Get(tagName); tag == "-" {
			continue
		}

		f := &Field{
			field: field,
			value: v.FieldByName(field.Name),
		}

		fields = append(fields, f)

	}

	return fields
}