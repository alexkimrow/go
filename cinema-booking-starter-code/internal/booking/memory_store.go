package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func () Book(b Booking) error
func () ListBookings(movieID string) []Booking
