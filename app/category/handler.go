package category

import (
	"inventaris/lib"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func (h CategoryHandler) ListCategory(w http.ResponseWriter, r *http.Request) {
	c := CategoryRepo{h.DB}
	id := lib.Trim(chi.URLParam(r, "id"))

	if id != "" {
		data, err := c.GetDataById(id)
		if err != nil {
			lib.SendMessageFail(w, 400, err.Error())
			return
		}
		lib.SendData(w, 201, data)
	} else {
		data, err := c.GetData()
		if err != nil {
			lib.SendMessageFail(w, 500, "Gagal mengambil kategori")
			return
		}
		if len(data) == 0 {
			lib.SendMessageFail(w, 404, "Data not found")
			return
		}
		lib.SendData(w, 201, data)
	}
}

func (h CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	c := CategoryRepo{h.DB}

	body := lib.ParseBody[struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}](r)

	if body.Name == "" {
		lib.SendMessageFail(w, 400, "Nama kategaori tidak boleh kosong")
		return
	}

	result, err := c.Insert(Category{
		Name:        body.Name,
		Description: body.Description,
	})

	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	lib.SendDataMessage(w, 201, result, "Kategori berhasil ditambahkan")
}

// body { name, description }
func (h CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	c := CategoryRepo{h.DB}

	id := chi.URLParam(r, "id") // Get id from params route

	body := lib.ParseBody[struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}](r)

	if body.Name == "" {
		lib.SendMessageFail(w, 400, "Nama kategaori tidak boleh kosong")
		return
	}

	result, err := c.GetDataById(id)
	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	result, err = c.Update(Category{
		ID:          id,
		Name:        body.Name,
		Description: body.Description,
	})

	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	lib.SendDataMessage(w, 201, result, "Kategori berhasil diperbarui")
}

func (h CategoryHandler) RemoveCategory(w http.ResponseWriter, r *http.Request) {
	c := CategoryRepo{h.DB}

	id := chi.URLParam(r, "id") // Get id from params route

	result, err := c.GetDataById(id)
	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	result, err = c.Delete(id)
	if err != nil {
		lib.SendMessageFail(w, 400, err.Error())
		return
	}

	lib.SendDataMessage(w, 201, result, "Kategori berhasil dihapus")
}
