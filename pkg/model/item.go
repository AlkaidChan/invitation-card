package model

import (
	"github.com/alkaidchen/invitation-card/pkg/log"
	"gorm.io/gorm"
)

type Item struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Phone      string `gorm:"type:varchar(255)" json:"phone"`
	Attendance int    `gorm:"type:int(11)" json:"attendance"`
}

// CreateItem handle create article
func CreateItem(tx *gorm.DB, article *Item) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(article)
	return result.Error
}

// UpdateItemByID update article by id and data
func UpdateItemByID(tx *gorm.DB, article *Item) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Model(&article).Select("*").Omit("publish_time").Updates(article)
	return result.Error
}

// DeleteItemByID delete article by id
func DeleteItemByID(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Delete(&Item{}, id)
	return result.Error
}

func IncrementItemVisits(id uint) {
	log.Debug("AddItemView")
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	db.Model(&Item{}).Where("id = ?", id).Update("visits", gorm.Expr("visits + 1"))
}

// GetItemByID return article by id
func GetItemByID(id uint) (Item, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	article := Item{}
	result := db.First(&article, id)
	return article, result.Error
}

// GetItemBySlug return article by slug
func GetItemBySlug(slug string) (Item, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	article := Item{}
	result := db.First(&article, "`slug` = ?", slug)
	return article, result.Error
}

// ListAllItems return all articles
func ListAllItems(pageNum, pageSize int) ([]Item, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var articles []Item
	result := db.Order("publish_time DESC").Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles)
	return articles, result.Error
}

// ListLatestItems return latest articles
func ListLatestItems() ([]Item, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var articles []Item
	result := db.Order("publish_time DESC").Limit(8).Find(&articles)
	return articles, result.Error
}

// ListItemsByCategoryID return articles by categoryID
func ListItemsByCategoryID(cid uint, pageNum, pageSize int) ([]Item, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var articles []Item
	result := db.Where("id in (?)", db.Table("article_categories").Select("article_id").
		Where("category_id = ?", cid)).Order("publish_time DESC").Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles)
	return articles, result.Error
}

// CountAllItems return count of all articles
func CountAllItems() (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Item{}).Count(&count)
	return count, result.Error
}
