package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
	"time"
)

func (d *Dao) CountArticle(title string, state uint8) (int64, error) {
	article := model.Article{Title: title, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticle(id uint32) (*model.Article, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	return article.GetArticle(d.engine)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title, content, description, coverUrl string, state uint8, createdBy string) error {
	article := model.Article{
		Title: title,
		State: state,
		Content: content,
		Description: description,
		CoverImageUrl: coverUrl,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title, content, description, coverUrl string, state uint8, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
		"modified_on": time.Now(),
	}
	if title != "" {
		values["title"] = title
	}
	if content != "" {
		values["content"] = content
	}
	if coverUrl != "" {
		values["cover_image_url"] = coverUrl
	}
	if description != "" {
		values["description"] = description
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	return article.Delete(d.engine)
}
