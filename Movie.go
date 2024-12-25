package Movie

import "fmt"

type Movie struct {
	ID          int
	Title       string
	Genre       string
	IsAvailable bool
}

type Customer struct {
	ID           int
	Name         string
	RentedMovies []int
}

type Rental struct {
	RentalId   int
	MovieId    int
	CustomerId int
	RentDate   string
	ReturnDate string
}

type RentalStore struct {
	Movies    map[int]*Movie
	Customers map[int]*Customer
	Rentals   map[int]*Rental
}

func (RS *RentalStore) Init() {
	RS.Movies = make(map[int]*Movie)
	RS.Customers = make(map[int]*Customer)
	RS.Rentals = make(map[int]*Rental)
}

func (RS *RentalStore) AddMovies(ID int, Title, Genre string) {
	RS.Movies[ID].ID = ID
	RS.Movies[ID].Title = Title
	RS.Movies[ID].Genre = Genre
	RS.Movies[ID].IsAvailable = true
	//RS.Movies[ID] = {
	//	ID : ID
	//	Title : Title
	//	Genre : Genre
	//	IsAvailable : false
	//}
}

func (RS *RentalStore) RegisterCustomer(ID int, Name string, RentedMovies []int) {
	RS.Customers[ID].ID = ID
	RS.Customers[ID].Name = Name
	RS.Customers[ID].RentedMovies = RentedMovies
}

func (RS *RentalStore) RentMovieToCustomer(RentalId, MovieId int, CustomerId int, RentDate string, ReturnDate string) {
	RS.Rentals[RentalId].RentalId = RentalId
	RS.Rentals[RentalId].MovieId = MovieId
	RS.Rentals[RentalId].CustomerId = CustomerId
	RS.Rentals[RentalId].RentDate = RentDate
	RS.Rentals[RentalId].ReturnDate = ReturnDate
	RS.Movies[MovieId].IsAvailable = false
}

func (RS *RentalStore) ReturnMovie(MovieId int) {
	RS.Movies[MovieId].IsAvailable = true
}

func (RS *RentalStore) ViewAll() {
	fmt.Println("Available: ")
	for _, value := range RS.Movies {
		if value.IsAvailable {
			fmt.Printf("ID: '%v', Title: '%v', Genre '%v'", value.ID, value.Title, value.Genre)
		}
	}
	fmt.Println("Not Available: ")
	for _, value := range RS.Movies {
		if !value.IsAvailable {
			fmt.Printf("ID: '%v', Title: '%v', Genre '%v'", value.ID, value.Title, value.Genre)
		}
	}
}
