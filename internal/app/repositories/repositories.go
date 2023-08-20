package repositories

import (
	"math"
	"sync"

	"gorm.io/gorm"

	"eql/internal/app/entities/schema"
)

type Repository struct {
	mutex sync.Mutex
}

// NewRepository new repository
func NewRepository() Repository {
	return Repository{}
}

// PageForm page info interface
type PageForm interface {
	GetPage() int
	GetSize() int
	GetQuery() string
	GetSort() string
	GetReverse() bool
	GetOrderBy() string
}

type PagePatientForm interface {
	GetPage() int
	GetSize() int
	GetQuery() string
	GetHn() string
	GetFname() string
	GetCid() string
	GetStartDate() string
	GetEndDate() string
	GetPatientColorId() string
	GetSort() string
	GetReverse() bool
	GetOrderBy() string
}

const (
	// DefaultPage default page in page query
	DefaultPage int = 1
	// DefaultSize default size in page query
	DefaultSize int = 20
)

// FindAllAndPageInformation get page information
func (r *Repository) FindAllAndPageInformation(db *gorm.DB, pageForm PageForm, entities interface{}, selector ...string) (*schema.PageInformation, error) {
	var count int64

	db.Model(entities).Count(&count)

	if pageForm.GetOrderBy() != "" {

		db = db.Order(pageForm.GetOrderBy())

	} else if pageForm.GetSort() != "" {

		order := pageForm.GetSort()
		if pageForm.GetReverse() {
			order = order + " desc"
		}
		db = db.Order(order)
	}

	if pageForm.GetQuery() != "" {
		db = db.Where(pageForm.GetQuery())
	}

	page := pageForm.GetPage()
	if pageForm.GetPage() < 1 {
		page = DefaultPage
	}

	limit := pageForm.GetSize()
	if pageForm.GetSize() == 0 {
		limit = DefaultSize
	}

	var offset int
	if page != 1 {
		offset = (page - 1) * limit
	}

	if len(selector) > 0 {
		db = db.Select(selector)
	}

	if err := db.
		Limit(limit).
		Offset(offset).
		Find(entities).Error; err != nil {
		return nil, err
	}

	return &schema.PageInformation{
		Page:     page,
		Size:     limit,
		Count:    count,
		LastPage: int(math.Ceil(float64(count) / float64(limit))),
	}, nil
}

func (r *Repository) FindAllPageInfoJoinTable(db *gorm.DB, pageForm PageForm, entities interface{}, selector ...string) (*schema.PageInformation, error) {
	var count int64

	db.Count(&count)

	if pageForm.GetOrderBy() != "" {

		db = db.Order(pageForm.GetOrderBy())

	} else if pageForm.GetSort() != "" {

		order := pageForm.GetSort()
		if pageForm.GetReverse() {
			order = order + " desc"
		}
		db = db.Order(order)
	}

	if pageForm.GetQuery() != "" {
		db = db.Where(pageForm.GetQuery())
	}

	page := pageForm.GetPage()
	if pageForm.GetPage() < 1 {
		page = DefaultPage
	}

	limit := pageForm.GetSize()
	if pageForm.GetSize() == 0 {
		limit = DefaultSize
	}

	var offset int
	if page != 1 {
		offset = (page - 1) * limit
	}

	if len(selector) > 0 {
		db = db.Select(selector)
	}

	if err := db.
		Limit(limit).
		Offset(offset).
		Find(entities).Error; err != nil {
		return nil, err
	}

	return &schema.PageInformation{
		Page:     page,
		Size:     limit,
		Count:    count,
		LastPage: int(math.Ceil(float64(count) / float64(limit))),
	}, nil
}
