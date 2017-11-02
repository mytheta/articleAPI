package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"

)

type ArticleApi struct {
    *revel.Controller
}


func (c ArticleApi) GetArticles() revel.Result {
  // articleモデルを利用
  articles := []models.Article{}
  // Idが降順になるように取得
  DB.Order("id desc").Find(&articles)

  response := JsonResponse{}
  response.Response = articles // 結果を格納してあげる

  return c.RenderJSON(response)
}

func (c ArticleApi) GetArticle() revel.Result {

    response := JsonResponse{}
    response.Response = "single article"

    return c.RenderJSON(response)
}

func (c ArticleApi) PostArticle() revel.Result {

    response := JsonResponse{}
    response.Response = "post article"

    return c.RenderJSON(response)
}

func (c ArticleApi) PutArticle() revel.Result {

    response := JsonResponse{}
    response.Response = "put article"

    return c.RenderJSON(response)
}

func (c ArticleApi) DeleteArticle() revel.Result {

    response := JsonResponse{}
    response.Response = "delete article"

    return c.RenderJSON(response)
}
