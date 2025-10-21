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
	"time"

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
