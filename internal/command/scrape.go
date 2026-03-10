package command

import (
	"os"
	"fmt"
	"github.com/google/uuid"
	"context"
	"time"
	"github.com/CromartyForth/gator/internal/config"
	"github.com/CromartyForth/gator/internal/database"
)


func scrapeFeed (s State) error {
	// get the next feed url from the db
	contextBackground := context.Background()
	url, err := s.Db.GetNextFeedToFetch(contextBackground)
	if err != nil {
		return fmt.Errorf("Error getting oldest feed from database: %v", err)
	}

	// retrieve actual feed
	


	return nil
}



/*
Write an aggregation function, I called mine scrapeFeeds. It should:
Get the next feed to fetch from the DB.
Mark it as fetched.
Fetch the feed using the URL (we already wrote this function)
Iterate over the items in the feed and print their titles to the console.
*/