package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/muhtutorials/rss/db"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	result, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}
	insertedUserID, err := result.LastInsertId()
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retrieve user ID: %s", err))
		return
	}
	user, err := apiCfg.DB.GetUserByID(r.Context(), insertedUserID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retreive user: %s", err))
		return
	}
	respondWithJSON(w, 201, dbUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user db.User) {
	respondWithJSON(w, 200, dbUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user db.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), db.GetPostsForUserParams{
		UserID: user.ID, Limit: 100,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retreive user posts: %s", err))
		return
	}

	respondWithJSON(w, 200, dbPostsToPosts(posts))
}
