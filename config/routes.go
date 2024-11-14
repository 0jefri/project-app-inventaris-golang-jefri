package config

import (
	"inventaris/app/category"
	"inventaris/app/items"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *chi.Mux {
	r := Router()

	ch := category.CategoryHandler{db}
	r.Route("/api/category", func(r chi.Router) {
		r.Get("/", ch.ListCategory) // list data category
		r.Post("/", ch.AddCategory) // create categoy

		r.Get("/{id}", ch.ListCategory)      // detail category by id
		r.Put("/{id}", ch.UpdateCategory)    // update category by id
		r.Delete("/{id}", ch.RemoveCategory) // remove category by id
	})

	ph := items.ItemsHandler{db}
	r.Route("/api/items", func(r chi.Router) {
		r.Get("/", ph.ListItems)          // list data items
		r.Post("/", ph.AddItems)          // creae items
		r.Get("/{id}", ph.ListItems)      // detail items by id
		r.Put("/{id}", ph.UpdateItems)    // update items by id
		r.Delete("/{id}", ph.RemoveItems) // remove item by id

		r.Get("/replacement-needed", ph.ListItemsReplacementNeeded) // list items needed to replacement
		r.Get("/investment", ph.ListItemsInvestment)                // total investment and depreciated value
		r.Get("/investment/{id}", ph.ListItemsInvestment)           // detail investment and depreciated by id
	})

	return r
}
