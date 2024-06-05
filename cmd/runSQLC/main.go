package main

import (
	"context"
	"database/sql"

	"github.com/DiegoSenaa/go-sql/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.NewString(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "5b050c4e-486e-47a6-ae81-55f7074be761",
		Name:        "Backend",
		Description: sql.NullString{String: "Nova descrição", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.DeleteCategory(ctx, "5b050c4e-486e-47a6-ae81-55f7074be761")
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		print("pós delete")
		println(category.ID, category.Name, category.Description.String)
	}
}
