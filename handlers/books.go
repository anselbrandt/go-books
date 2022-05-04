package handlers

import (
	"database/sql"
	"fmt"
	"go-books/data"
	"log"
	"net/http"
)

type Env struct {
	DB *sql.DB
}

// Define GetAll as a method on Env.
func (env *Env) GetAll(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	bks, err := data.AllBooks(env.DB)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
