package service

import (
	"mime/multipart"
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

// TODO: add search by user and keywords (Boolean query)

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

// not atomic
func SavePost(post *model.Post, file multipart.File) error {
	// 1. save file to GCS and get the URL
	url, err := repository.GCSBackend.SaveToGCS(file, post.Id)
	if err != nil {
		return err
	}
	// 2. set the URL to post
	post.Url = url
	// 3. save post to ES
	return repository.ESBackend.SaveToES(post, constants.POST_INDEX, post.Id)
    // if error(latency)
    // 1. retry: save to GCS 3 times
    // 2. roll back: delete from GCS
    // if GCS fail: add a offline cleanup at end of day
}