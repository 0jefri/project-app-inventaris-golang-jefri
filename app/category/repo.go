package category

import (
	"fmt"
	"inventaris/lib"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func (h *CategoryRepo) GetData() (result []Category, err error) {
	err = h.DB.Order("updated_at DESC").Find(&result).Error
	return
}

func (h *CategoryRepo) GetDataById(id string) (result Category, err error) {
	err = h.DB.First(&result, "id = ?", id).Error
	if result.ID == "" {
		err = fmt.Errorf("kategori tidak ditemukan")
	}
	return
}

func (h *CategoryRepo) GetDataByMultiId(ids ...string) (result map[string]Category, err error) {
	condId := []string{}
	dataId := []string{}
	for i, _ := range ids {
		if !slices.Contains(dataId, ids[i]) {
			condId = append(condId, "?")
			dataId = append(dataId, ids[i])
		}
	}

	data := []Category{}
	cond := fmt.Sprintf("id IN (%s)", strings.Join(condId, ", "))
	err = h.DB.
		Where(cond, lib.ToSliceAny(dataId)...).
		First(&data).Error

	if len(data) == 0 {
		err = fmt.Errorf("kategori tidak ditemukan")
	}

	result = map[string]Category{}
	for i, _ := range data {
		result[data[i].ID] = data[i]
	}

	return
}

func (h *CategoryRepo) Insert(data Category) (result Category, err error) {
	data.ID = uuid.NewString()
	err = h.DB.Create(&data).Error
	if err == nil {
		result, _ = h.GetDataById(data.ID)
	}
	return
}

func (h *CategoryRepo) Delete(id string) (result Category, err error) {
	err = h.DB.Where("id = ?", id).Delete(&result).Error
	if err == nil {
		result, _ = h.GetDataById(id)
	}
	return
}

func (h *CategoryRepo) Update(data Category) (result Category, err error) {
	result, err = h.GetDataById(data.ID)
	if err == nil {
		if data.Name != "" {
			err = h.DB.Model(&result).Update("name", data.Name).Error
		}
		if data.Description != "" {
			err = h.DB.Model(&result).Update("description", data.Description).Error
		}
		result, _ = h.GetDataById(data.ID)
	}
	return
}
