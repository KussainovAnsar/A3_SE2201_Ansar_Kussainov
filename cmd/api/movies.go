package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DataDavD/snippetbox/greenlight/internal/data"
	"github.com/DataDavD/snippetbox/greenlight/internal/validator"
)

// createMovieHandler handles "POST /v1/movies" endpoint. For now, we just return a plain-text
// placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Declare an anonymous struct to hold the information that we expect to be in the HTTP
	// request body (not that the field names and types in the struct are a subset of the Movie
	// struct). This struct will be our *target decode destination*.
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	// Use the readJSON() helper to decode the request body into the struct.
	// If this returns an error we send the client the error message along with
	// a 400 Bad Request status code.
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Movie struct.
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	// Initialize a new Validator instance.
	v := validator.New()

	// Call the ValidateMovie() function and return a response containing the errors if any of
	// the checks fail.
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Call the Insert() method on our movies model, passing in a pointer to the validated movie
	// struct. This will create a record in the database and update the movie struct with the
	// system-generated information.
	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.infoLog.Println("inside err if statement after movies insert")
		app.serverErrorResponse(w, r, err)
		return
	}
	// When sending an HTTP response,
	// we want to include a Location header to let the client know which URL they can find the
	// newly created resource at. We make an empty http.Header map and then use the Set()
	// method to add a new Location header,
	// interpolating the system-generated ID for our new movie in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	// Write a JSON response with a 201 Created status code, the movie data in the response body,
	// and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
	if err != nil {
		app.infoLog.Println("inside err if statement after movies insert")
		app.serverErrorResponse(w, r, err)
		return
	}
}

// showMovieHandler handles "Get /v1/movies/:id" endpoint. For now, it returns plain-text
// placeholder response using the interpolated "id" parameter from the current URL
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, any interpolated URL Parameters will be stored
	// in the request context. We can use the ParamsFromContext() function to retrieve a slice
	// containing these paremter names and values.
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	// Call the Get() method to fetch the data for a specific movie.
	// We also need to use the errors.Is()
	// function to check if it returns a data.ErrRecordNotFound error,
	// in which case we send a 404 Not Found response to the client.
	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Create an envelope{"movie": movie} instance and pass it to writeJSON(), instead of passing
	// the plain movie struct.
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	return
}
