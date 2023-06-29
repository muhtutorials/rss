package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/muhtutorials/rss/db"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		UserID    int64     `json:"user_id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		CreatedAt time.Time `json:"created_at"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	result, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		UserID: user.ID,
		Name:      params.Name,
		Url: params.URL,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %s", err))
		return
	}
	insertedFeedID, err := result.LastInsertId()
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retrieve feed ID: %s", err))
		return
	}
	feed, err := apiCfg.DB.GetFeedByID(r.Context(), insertedFeedID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retreive feed: %s", err))
		return
	}
	respondWithJSON(w, 201, dbFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't retrieve feeds: %s", err))
		return
	}
	respondWithJSON(w, 201, dbFeedsToFeeds(feeds))
}