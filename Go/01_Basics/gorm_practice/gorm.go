package main

/*
Practice Task: "Movie Collection App"

Requirements:
1. Connect to a SQLite database named "movies.db".
   - If the connection fails, print an error and stop the program.
   - If successful, print a confirmation message in the console.

2. Define a struct Movie with the following fields:
   - ID (primary key)
   - Title (cannot be null)
   - Director
   - Year
   - Rating (float64)
   - CreatedAt, UpdatedAt, DeletedAt

3. Use GORM struct tags to:
   - mark the primary key,
   - prevent null values for the Title,
   - and create an index on DeletedAt.

4. Perform an automatic migration to create the "movies" table.

5. Implement the basic CRUD operations:
   - Add at least two movies.
   - Read and print all movies.
   - Update the rating of one movie.
   - Delete one movie (use both soft and permanent delete).
*/

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Movie represents a movie record in the database.
type Movie struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Director  string
	Year      uint
	Rating    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {

	// 1. Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	fmt.Println("Database connection successful.")

	// 2. Auto-migrate the schema
	if err := db.AutoMigrate(&Movie{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migration completed.")

	// 3. Add movies
	movie1 := &Movie{
		Title:    "Inception",
		Director: "Christopher Nolan",
		Year:     2010,
		Rating:   8.8,
	}

	movie2 := &Movie{
		Title:    "The Matrix",
		Director: "Lana & Lilly Wachowski",
		Year:     1999,
		Rating:   8.7,
	}

	db.Create(&movie1)
	db.Create(&movie2)
	fmt.Println("Two movies added to the database.")

	// 4. Read and print all movies
	var allMovies []Movie
	if err := db.Find(&allMovies).Error; err != nil {
		log.Println("Error reading movies:", err)
	}
	fmt.Println("\nAll movies before update:")
	for _, m := range allMovies {
		fmt.Printf("→ [%d] %s (%d) — Rating: %.1f\n", m.ID, m.Title, m.Year, m.Rating)
	}

	// 5. Update the rating of one movie
	db.Model(&movie1).Update("rating", 9.9)
	fmt.Printf("\nUpdated rating for '%s' to %.1f.\n", movie1.Title, 9.9)

	// 6. Read again and print updated list
	var updatedMovies []Movie
	db.Find(&updatedMovies)
	fmt.Println("\nAll movies after update:")
	for _, m := range updatedMovies {
		fmt.Printf("→ [%d] %s (%d) — Rating: %.1f\n", m.ID, m.Title, m.Year, m.Rating)
	}

	// 7. Soft delete and hard delete
	db.Delete(&movie1)            // Soft delete
	db.Unscoped().Delete(&movie2) // Hard delete
	fmt.Println("\nDeleted one movie (soft) and one permanently (hard).")

}
