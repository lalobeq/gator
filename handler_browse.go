package main

import (
	"context"
	"github.com/lalobec/gator/internal/database"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	
	if len(cmd.arguments) == 2 {
		parseNum, err := strconv.ParseInt(cmd.arguments[1], 10, 32)
		if err != nil {
			fmt.Printf("You must provide a positive integer: %v", err)
		}
		limit = int32(parseNum)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: limit,
		})
	if err != nil {
		return nil
	}
	for _, post := range posts {
		fmt.Println("------------------------------------------------------")
		fmt.Println(post.Title)
		fmt.Println("------------------------------------------------------")
		fmt.Println(post.Description)
		fmt.Printf("Link: %s \n", post.Url)
		fmt.Println()
		fmt.Println()
	}

	return nil
}
