package items

import (
	"math"
	"net/http"
	"strconv"
	"time"

	// "inventaris/app/category"
	"inventaris/app/category"
	"inventaris/lib"

	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
)

type ItemsHandler struct {
	DB *gorm.DB
}

type bodyItemsHandler struct {
	IdCategory   string    `json:"category_id"`
	Name         string    `json:"name"`
	PhotoURL     string    `json:"photo_url"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
}

func (h ItemsHandler) ListItems(w http.ResponseWriter, r *http.Request) {
	it := ItemsRepo{h.DB}
	id := lib.Trim(chi.URLParam(r, "id"))

	if id != "" {
		data, err := it.GetDataById(id)
		if err != nil {
			lib.SendMessageFail(w, 400, err.Error())
			return
		}
		lib.SendData(w, 201, data)
	} else {

		limit, offset := lib.GetLimitOffset(r)

		data, count, err := it.GetData(ItemsRepoFilter{
			Limit:      limit,
			Offset:     offset,
			CategoryId: r.URL.Query().Get("category_id"),
			Age:        lib.ParseToInt(r.URL.Query().Get("age")),
		})

		if err != nil {
			lib.SendMessageFail(w, 500, err.Error())
			return
		}

		if len(data) == 0 {
			lib.SendMessageFail(w, 404, "Data not found")
			return
		}

		lib.SendResponse(w, 201, map[string]any{
			"success":     true,
			"page":        math.Ceil(float64(offset)/float64(limit)) + 1,
			"limit":       limit,
			"total_item":  count,
			"total_pages": math.Ceil(float64(count) / float64(limit)),
			"data":        data,
		})
	}
}

func (h ItemsHandler) ListItemsReplacementNeeded(w http.ResponseWriter, r *http.Request) {
	it := ItemsRepo{h.DB}

	data, count, err := it.GetData(ItemsRepoFilter{NoLimit: true})

	if err != nil {
		lib.SendMessageFail(w, 500, err.Error())
		return
	}

	if len(data) == 0 {
		lib.SendMessageFail(w, 404, "Data not found")
		return
	}

	lib.SendResponse(w, 201, map[string]any{
		"success":    true,
		"total_item": count,
		"data":       data,
	})
}

func (h ItemsHandler) ListItemsInvestment(w http.ResponseWriter, r *http.Request) {
	it := ItemsRepo{h.DB}
	id := lib.Trim(chi.URLParam(r, "id"))
	dep_rate := 10

	res := map[string]any{
		"depreciation_rate": dep_rate,
		"depreciated_value": 0,
	}

	if id != "" {
		result, err := it.GetDataById(id)
		if err != nil {
			lib.SendMessageFail(w, 400, err.Error())
			return
		}
		res["item_id"] = result.ID
		res["name"] = result.Name
		res["initial_price"] = result.Price
		res["item_detail"] = result
	}

	data, err := it.TotalInvestment(10, id)
	if err != nil {
		lib.SendMessageFail(w, 500, "Gagal menghitung total investasi")
		return
	}

	if id == "" {
		res["total_investment"] = data.TotalInvestment
	}

	res["depreciated_value"] = data.DepreciatedValue

	lib.SendData(w, 201, res)
}

func (h ItemsHandler) GetFormData(r *http.Request) Items {
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	purchaseDate, _ := lib.ParseTime(r.FormValue("purchase_date"))

	return Items{
		IdCategory:   r.FormValue("category_id"),
		Name:         r.FormValue("name"),
		PhotoURL:     r.FormValue("photo_url"),
		Price:        price,
		PurchaseDate: purchaseDate,
	}
}

func (h ItemsHandler) AddItems(w http.ResponseWriter, r *http.Request) {
	c := category.CategoryRepo{h.DB}
	it := ItemsRepo{h.DB}

	body := h.GetFormData(r)

	if body.Name == "" {
		lib.SendMessageFail(w, 400, "Nama items tidak boleh kosong")
		return
	}

	_, err := c.GetDataById(body.IdCategory)
	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	result, err := it.Insert(body)

	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	lib.SendDataMessage(w, 201, result, "Barang berhasil ditambahkan")
}

func (h ItemsHandler) UpdateItems(w http.ResponseWriter, r *http.Request) {
	c := category.CategoryRepo{h.DB}
	it := ItemsRepo{h.DB}

	body := h.GetFormData(r)
	body.ID = lib.Trim(chi.URLParam(r, "id"))

	if body.Name == "" {
		lib.SendMessageFail(w, 400, "Nama items tidak boleh kosong")
		return
	}

	_, err := c.GetDataById(body.IdCategory)
	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	result, err := it.Update(body)

	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	lib.SendDataMessage(w, 201, result, "Barang berhasil diperbarui")
}

func (h ItemsHandler) RemoveItems(w http.ResponseWriter, r *http.Request) {
	it := ItemsRepo{h.DB}

	id := chi.URLParam(r, "id") // Get id from params route

	result, err := it.GetDataById(id)
	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	result, err = it.Delete(id)
	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	lib.SendDataMessage(w, 201, result, "Barang berhasil dihapus")
}
