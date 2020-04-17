package month

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server is the backend month server for MonthAI.
type Server struct {
	mux *mux.Router

	month *Month
}

// NewServer creates and returns a new Server.
func NewServer() *Server {
	mux := mux.NewRouter()
	month := NewMonth()

	return &Server{
		mux:   mux,
		month: month,
	}
}

// ServeHTTP satisfies the http.Handler interface for Server.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.routes()

	s.mux.ServeHTTP(w, r)
}

// Routes registers the routes for the Server.
func (s *Server) routes() {
	s.mux.HandleFunc("/", s.month.showMonth)
	s.mux.HandleFunc("/next", s.month.nextMonth)
}

// Month contains the current month.
type Month struct {
	current time.Month
}

// NewMonth returns a new instance of Month.
func NewMonth() *Month {
	_, month, _ := time.Now().Date()
	return &Month{
		current: month,
	}
}

func (m *Month) showMonth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, m.current.String())
}

func (m *Month) nextMonth(w http.ResponseWriter, r *http.Request) {
	var next time.Month

	switch m.current {
	case time.January:
		next = time.February
	case time.February:
		next = time.March
	case time.March:
		next = time.April
	case time.April:
		next = time.May
	case time.May:
		next = time.June
	case time.June:
		next = time.July
	case time.July:
		next = time.August
	case time.August:
		next = time.September
	case time.September:
		next = time.October
	case time.October:
		next = time.November
	case time.November:
		next = time.December
	case time.December:
		next = time.January
	}

	fmt.Println(w, next.String())
}

func encodeJSON(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
