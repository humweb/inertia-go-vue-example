package resources

import (
	"github.com/humweb/go-tables/tables"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserResource struct {
	tables.AbstractResource
}

func NewUserResource(db *gorm.DB, req *http.Request) *UserResource {
	r := &UserResource{
		tables.AbstractResource{
			DB:              db,
			Request:         req,
			HasGlobalSearch: true,
		},
	}

	r.Fields = r.GetFields()
	r.Filters = r.GetFilters()

	return r
}

func (u *UserResource) ApplyFilter(db *gorm.DB) {
	//
}

func (u *UserResource) GetFields() []*tables.Field {
	return []*tables.Field{
		tables.NewField("ID", tables.WithSortable()),
		tables.NewField("First name", tables.WithSortable(), tables.WithVisibility()),
		tables.NewField("Last name", tables.WithSortable(), tables.WithSearchable()),
		tables.NewField("Email", tables.WithSortable()),
		tables.NewField("Update at",
			tables.WithFieldComponent("date"),
			tables.WithSortable(),
			tables.WithMeta(map[string]interface{}{
				"relative":   true,
				"dateFormat": "MMM D, YYYY h:mm A",
			}),
		),
	}
}

func (u *UserResource) GetFilters() []*tables.Filter {
	return []*tables.Filter{
		tables.NewFilter("ID"),
	}
}

func (u *UserResource) WithGlobalSearch(db *gorm.DB, val string) {

	if v, err := strconv.Atoi(val); err == nil {
		db.Where("id = ?", v)
	} else {
		val = "%" + val + "%"
		db.Where("(LOWER(first_name) like LOWER(?) OR LOWER(last_name) like LOWER(?) OR LOWER(email) like LOWER(?))", val, val, val)
	}
}
