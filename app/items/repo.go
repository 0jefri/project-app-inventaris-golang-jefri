package items

import (
	"fmt"
	"time"

	"inventaris/app/category"
	"inventaris/lib"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type ItemsRepo struct {
	DB *gorm.DB
}

type ItemsRepoFilter struct {
	Id            string
	Limit         int
	Offset        int
	NoLimit       bool
	CategoryId    string
	Age           int
	ReplaceNeeded bool
}

type ItemsResult struct {
	Items
	Category            string             `json:"category"`
	CategoryData        *category.Category `json:"category_detail"`
	TotalUsageDays      int                `json:"total_usage_days"`
	ReplacementRequired bool               `json:"replacement_required"`
}

type ItemInvestmentResult struct {
	TotalInvestment  float64
	DepreciatedValue float64
}

var MaxDays = 100

func (h *ItemsRepo) buildJoinQuery(filter ItemsRepoFilter) *gorm.DB {
	query := h.DB.Model(&Items{})

	if filter.NoLimit {
		query = query.Order("purchase_date ASC")
	} else {
		query = query.
			Limit(filter.Limit).
			Offset(filter.Offset).
			Order("updated_at DESC")
	}

	if filter.Id != "" {
		query.Where("id = ?", filter.Id)
	}

	if filter.CategoryId != "" {
		query.Where("category_id = ?", filter.CategoryId)
	}

	cond := "CURRENT_DATE - Date(purchase_date) > INTERVAL '? days'"

	if filter.Age != 0 {
		query.Where(cond, filter.Age)
	}

	if filter.ReplaceNeeded {
		query.Where(cond, MaxDays)
	}

	return query
}

func (h *ItemsRepo) GetData(filter ItemsRepoFilter) (result []ItemsResult, count int64, err error) {
	data := []Items{}

	err = h.buildJoinQuery(filter).Count(&count).Find(&data).Error
	idCategory := []string{}
	today := time.Now()

	for i, _ := range data {
		days := lib.DiffDays(data[i].PurchaseDate, today)

		result = append(result, ItemsResult{
			Items:               data[i],
			TotalUsageDays:      days,
			ReplacementRequired: days > MaxDays,
		})

		if !slices.Contains(idCategory, data[i].IdCategory) {
			idCategory = append(idCategory, data[i].IdCategory)
		}

	}

	if len(idCategory) > 0 {
		c := category.CategoryRepo{h.DB}
		cat, _ := c.GetDataByMultiId(idCategory...)

		for i, _ := range result {
			id := result[i].Items.IdCategory
			if val, ok := cat[id]; ok {
				result[i].Category = val.Name
				result[i].CategoryData = &val
			}
		}
	}

	return
}

func (h *ItemsRepo) GetDataById(id string) (result ItemsResult, err error) {
	res, _, _ := h.GetData(ItemsRepoFilter{Id: id, Limit: 1})
	if len(res) == 0 {
		err = fmt.Errorf("barang tidak ditemukan")
	}
	for i, v := range res {
		if v.ID == id {
			result = res[i]
		}
	}
	return
}

func (h *ItemsRepo) Insert(data Items) (result ItemsResult, err error) {
	data.ID = uuid.NewString()
	err = h.DB.Create(&data).Error
	if err == nil {
		result, _ = h.GetDataById(data.ID)
	}

	fmt.Printf("err create: %v\n", err)
	return
}

func (h *ItemsRepo) Delete(id string) (result ItemsResult, err error) {
	err = h.DB.Where("id = ?", id).Delete(&Items{}).Error
	if err == nil {
		result, _ = h.GetDataById(id)
	}
	return
}

func (h *ItemsRepo) Update(data Items) (result ItemsResult, err error) {
	result, err = h.GetDataById(data.ID)
	if err == nil {
		query := h.DB.Model(&Items{}).Where("id = ?", data.ID)

		if data.IdCategory != "" {
			err = query.Update("id_category", data.IdCategory).Error
		}
		if data.Name != "" {
			err = query.Update("name", data.Name).Error
		}

		query.Update("photo_url", data.PhotoURL)
		query.Update("price", data.Price)
		query.Update("purchase_date", data.PurchaseDate)
		result, _ = h.GetDataById(data.ID)
	}
	return
}

func (h *ItemsRepo) TotalInvestment(depreciationRate int, id string) (result ItemInvestmentResult, err error) {
	cond := ""

	if id != "" {
		cond += fmt.Sprintf("WHERE id = '%s' ", id)
	}

	query := fmt.Sprintf(`
		WITH depreciation AS (
			SELECT 
				id,
				price,
				purchase_date,
				price - ((price * 0.9) / ?) * EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)) AS depreciated_value
			FROM   %s  %s
		)
		SELECT 
			SUM(price) AS total_investment,
			SUM(depreciated_value) AS depreciated_value
		FROM 
			depreciation `,
		Items{}.TableName(),
		cond,
	)

	err = h.DB.Raw(query, depreciationRate).Scan(&result).Error
	return
}
