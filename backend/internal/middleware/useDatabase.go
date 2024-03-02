package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UseDatabase(db *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := db.Ping(ctx)
		if err != nil {
			fmt.Printf("Cannot connect to database: %s\n", err)
			os.Exit(1)
		}
		ctx.Set("DB", db)
		ctx.Next()
	}
}
