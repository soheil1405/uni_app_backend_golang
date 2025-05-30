package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/article/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	usecase usecase.ArticleUseCase
}

func NewArticleHandler(usecase usecase.ArticleUseCase, e echo.Group) {
	articleHandler := &ArticleHandler{usecase}

	articlesRouteGroup := e.Group("/articles")
	articlesRouteGroup.POST("", articleHandler.CreateArticle)
	articlesRouteGroup.GET("/:id", articleHandler.GetArticleByID)
	articlesRouteGroup.PUT("/:id", articleHandler.UpdateArticle)
	articlesRouteGroup.DELETE("/:id", articleHandler.DeleteArticle)
	articlesRouteGroup.GET("", articleHandler.GetAllArticles)
	articlesRouteGroup.GET("/slug/:slug", articleHandler.GetArticleBySlug)
	articlesRouteGroup.GET("/author/:author_id", articleHandler.GetArticlesByAuthor)
	articlesRouteGroup.GET("/published", articleHandler.GetPublishedArticles)
	articlesRouteGroup.GET("/category/:category_id", articleHandler.GetArticlesByCategory)
	articlesRouteGroup.GET("/tag/:tag_id", articleHandler.GetArticlesByTag)
	articlesRouteGroup.GET("/search", articleHandler.SearchArticles)
	articlesRouteGroup.POST("/:id/publish", articleHandler.PublishArticle)
	articlesRouteGroup.POST("/:id/archive", articleHandler.ArchiveArticle)
	articlesRouteGroup.POST("/:id/like", articleHandler.LikeArticle)
	articlesRouteGroup.POST("/:id/unlike", articleHandler.UnlikeArticle)
	articlesRouteGroup.POST("/:id/view", articleHandler.ViewArticle)
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	var article models.Article
	if err := c.Bind(&article); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateArticle(c.Request().Context(), &article); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) GetArticleByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	article, err := h.usecase.GetArticleByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) UpdateArticle(c echo.Context) error {
	var article models.Article
	var err error
	if article.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := c.Bind(&article); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateArticle(c.Request().Context(), &article); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) DeleteArticle(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.DeleteArticle(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Article deleted"}, nil)
}

func (h *ArticleHandler) GetAllArticles(c echo.Context) error {
	var request models.FetchArticleRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	articles, paginate, err := h.usecase.GetAllArticles(c.Request().Context(), request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"articles": articles}, paginate)
}

func (h *ArticleHandler) GetArticleBySlug(c echo.Context) error {
	slug := c.Param("slug")
	article, err := h.usecase.GetArticleBySlug(c.Request().Context(), slug)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) GetArticlesByAuthor(c echo.Context) error {
	authorID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	articles, err := h.usecase.GetArticlesByAuthor(c.Request().Context(), uint(authorID))
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"articles": articles}, nil)
}

func (h *ArticleHandler) GetPublishedArticles(c echo.Context) error {
	articles, err := h.usecase.GetPublishedArticles(c.Request().Context())
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"articles": articles}, nil)
}

func (h *ArticleHandler) GetArticlesByCategory(c echo.Context) error {
	categoryID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	articles, err := h.usecase.GetArticlesByCategory(c.Request().Context(), uint(categoryID))
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"articles": articles}, nil)
}

func (h *ArticleHandler) GetArticlesByTag(c echo.Context) error {
	tagID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	articles, err := h.usecase.GetArticlesByTag(c.Request().Context(), uint(tagID))
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"articles": articles}, nil)
}

func (h *ArticleHandler) SearchArticles(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return helpers.Reply(c, http.StatusBadRequest, nil, nil, map[string]string{"error": "search query is required"})
	}
	articles, err := h.usecase.SearchArticles(c.Request().Context(), query)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"articles": articles}, nil)
}

func (h *ArticleHandler) PublishArticle(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.PublishArticle(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	article, err := h.usecase.GetArticleByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) ArchiveArticle(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.ArchiveArticle(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	article, err := h.usecase.GetArticleByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) LikeArticle(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.LikeArticle(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	article, err := h.usecase.GetArticleByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) UnlikeArticle(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UnlikeArticle(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	article, err := h.usecase.GetArticleByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}

func (h *ArticleHandler) ViewArticle(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.ViewArticle(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	article, err := h.usecase.GetArticleByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"article": article}, nil)
}
