package service

import (
	"fmt"
	"socialai/constants"
	"socialai/model"
	"socialai/repository"

	"github.com/olivere/elastic/v7"
)

// login
func CheckUser(username, password string) (bool, error) {
	// construct query
	query := elastic.NewBoolQuery()
    query.Must(elastic.NewTermQuery("username", username))
    query.Must(elastic.NewTermQuery("password", password))

	// call repo
	searchResult, err := repository.ESBackend.ReadFromES(query, constants.USER_INDEX)
    if err != nil {
        return false, err
    }

	if searchResult.TotalHits() > 0 {
        fmt.Printf("Login as %s\n", username)
        return true, nil
    }

    return false, nil

}

// option 1: search username -> User 2. compare User.password with password
// option 2: search username + password; found -> login success

// register
func AddUser(user *model.User) (bool, error) {
	query := elastic.NewTermQuery("username", user.Username)
    searchResult, err := repository.ESBackend.ReadFromES(query, constants.USER_INDEX)
    if err != nil {
        return false, err
    }

    if searchResult.TotalHits() > 0 {
        return false, nil
    }

    err = repository.ESBackend.SaveToES(user, constants.USER_INDEX, user.Username)
    if err != nil {
        return false, err
    }
    fmt.Printf("User is added: %s\n", user.Username)
    return true, nil
}