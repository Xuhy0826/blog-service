package service

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ListArticleRequest struct {
	Title string `form:"title" binding:"max=25"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `json:"title" binding:"required,max=25"`
	Description   string `json:"description" binding:"max=200"`
	CoverImageUrl string `json:"cover_image_url" binding:""`
	Content       string `json:"content" binding:""`
	CreatedBy     string `json:"created_by" binding:"required"`
	State         uint8  `json:"state" binding:""`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `json:"title" binding:"required max=25"`
	Description   string `json:"description" binding:"max=200"`
	CoverImageUrl string `json:"cover_image_url" binding:""`
	Content       string `json:"content" binding:""`
	ModifiedBy    string `json:"modified_by" binding:"required"`
	State         uint8  `json:"state" binding:""`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int64, error) {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc *Service) GetArticleList(param *ListArticleRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Content, param.Description, param.CoverImageUrl, param.State, param.CreatedBy)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Content, param.Description, param.CoverImageUrl, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
