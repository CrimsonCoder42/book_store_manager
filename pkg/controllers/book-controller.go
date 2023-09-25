var newBook models.Book  // Variable to hold a new Book object

// GetBook fetches all book records and returns as JSON.
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()  // Fetch all books
	res, _ := json.Marshal(newBooks)  // Convert the slice of books to JSON
	w.Header().Set("Content-Type", "application/json")  // Set the content type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status code to OK (200)
	w.Write(res)  // Write the JSON response
}

// GetBookById fetches a book by its ID and returns as JSON.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)  // Get URL variables
	bookId := vars["bookId"]  // Get the bookId from URL variables
	ID, err := strconv.ParseInt(bookId, 0, 0)  // Convert bookId to integer
	if err != nil {
		fmt.Println("Error while parsing")  // Log parsing error
	}
	bookDetails, _ := models.GetBookById(ID)  // Fetch book by ID
	res, _ := json.Marshal(bookDetails)  // Convert fetched book to JSON
	w.Header().Set("Content-Type", "application/json")  // Set the content type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status code to OK (200)
	w.Write(res)  // Write the JSON response
}

// CreateBook creates a new book and returns it as JSON.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}  // Initialize an empty Book struct
	utils.ParseBody(r, CreateBook)  // Parse the request body into the CreateBook variable
	b := CreateBook.CreateBook()  // Create the book in the database
	res, _ := json.Marshal(b)  // Convert the created book to JSON
	w.Header().Set("Content-Type", "application/json")  // Set the content type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status code to OK (200)
	w.Write(res)  // Write the JSON response
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

