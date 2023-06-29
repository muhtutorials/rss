package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/muhtutorials/rss/db"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID    int64     `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	result, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %s", err))
		return
	}
	insertedFeedFollowID, err := result.LastInsertId()
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retrieve feed follow ID: %s", err))
		return
	}
	feedFollow, err := apiCfg.DB.GetFeedFollowByID(r.Context(), insertedFeedFollowID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retreive feed follow: %s", err))
		return
	}
	respondWithJSON(w, 201, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetUserFeedFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollows, err := apiCfg.DB.GetUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user feed follows: %s", err))
		return
	}

	respondWithJSON(w, 201, dbFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := strconv.ParseInt(feedFollowIDStr, 10, 64)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Couldn't parse feed follow ID: %s", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %s", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}