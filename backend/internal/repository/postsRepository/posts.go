package postsRepository

import (
	"blog-app-backend/internal/models"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func GetAllPosts(ctx *gin.Context) ([]models.Post, error) {
	db := ctx.MustGet("DB").(*pgxpool.Pool)

	sql, _, err := psql.Select("*").From("posts").ToSql()
	if err != nil {
		fmt.Printf("Error generating SQL: %s\n", err)
		return nil, err
	}

	query, err := db.Query(context.Background(), sql)
	if err != nil {
		fmt.Printf("Error querying database: %s\n", err)
		return nil, err
	}
	defer query.Close()

	posts, err := pgx.CollectRows(query, pgx.RowToStructByName[models.Post])
	if err != nil {
		fmt.Printf("CollectRows error: %v\n", err)
		return nil, err
	}

	return posts, nil
}

func GetPostById(ctx *gin.Context, postId int) (*models.Post, error) {
	db := ctx.MustGet("DB").(*pgxpool.Pool)
	sql, args, err := psql.Select("*").From("posts").Where(squirrel.Eq{"post_id": postId}).ToSql()

	post := models.Post{}

	if err != nil {
		fmt.Printf("Error generating SQL: %s\n", err)
		return nil, err
	}
	query, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		fmt.Printf("Error querying database: %s\n", err)
		return nil, err
	}
	defer query.Close()

	post, err = pgx.CollectOneRow(query, pgx.RowToStructByName[models.Post])
	if err != nil {
		fmt.Printf("CollectRows error: %v\n", err)
		return nil, err
	}

	return &post, nil
}

func CreateNewPost(ctx *gin.Context, newPost models.CreatePostRequest) (*models.Post, error) {
	db := ctx.MustGet("DB").(*pgxpool.Pool)
	sql, args, err := psql.Insert("posts").Columns("title", "content", "author").Values(newPost.Title, newPost.Content, newPost.Author).Suffix("RETURNING \"post_id\"").ToSql()

	//post := models.Post{}
	var newPostId int
	if err != nil {
		fmt.Printf("Error generating SQL: %s\n", err)
		return nil, err
	}

	query := db.QueryRow(context.Background(), sql, args...)
	err = query.Scan(&newPostId)
	if err != nil {
		fmt.Printf("Error creating a row.")
		return nil, err
	}

	post, err := GetPostById(ctx, newPostId)
	if err != nil {
		fmt.Printf("Error fetching post on return.")
		return nil, err
	}

	return post, nil
}

func DeletePostById(ctx *gin.Context, postId int) (bool, error) {
	db := ctx.MustGet("DB").(*pgxpool.Pool)
	sql, args, err := psql.Delete("posts").Where(squirrel.Eq{"post_id": postId}).ToSql()

	if err != nil {
		fmt.Printf("Error generating SQL: %s\n", err)
		return false, err
	}

	query := db.QueryRow(ctx, sql, args...)
	err = query.Scan()
	if err != nil {
		fmt.Printf("Error deleting a post.")
	}

	return true, nil

}
