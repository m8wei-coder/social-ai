package service

import (
	"reflect"
	"socialai/constants"
	"socialai/model"
	"socialai/repository"

	"github.com/olivere/elastic/v7"
)

func SearchPostsByUser(user string) ([]model.Post, error) {
   // 1. create a query
   query := elastic.NewTermQuery("user", user)

   // 2. call repo
   searchResult, err := repository.ESBackend.ReadFromES(query, constants.POST_INDEX)
   if err != nil {
       return nil, err
   }

   return getPostFromSearchResult(searchResult), nil
}

func SearchPostsByKeywords(keywords string) ([]model.Post, error) {
   //option1:return nothing
   // if keywords == "" {
   //  return nil, nil
   // }

   // 1. create a query
   query := elastic.NewMatchQuery("message", keywords)
   query.Operator("AND")
   //option 2, return all
   if keywords == "" {
       query.ZeroTermsQuery("all")
   }

   // 2. call repo
   searchResult, err := repository.ESBackend.ReadFromES(query, constants.POST_INDEX)
   if err != nil {
       return nil, err
   }

   return getPostFromSearchResult(searchResult), nil
}

func getPostFromSearchResult(searchResult *elastic.SearchResult) []model.Post {
   var ptype model.Post
   var posts []model.Post

   for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
       p := item.(model.Post)
       posts = append(posts, p)
   }
   return posts
}