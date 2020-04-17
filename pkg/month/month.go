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
}

// Month contains the current month.
type Month struct {
	current string
}

// NewMonth returns a new instance of Month.
func NewMonth() *Month {
	_, month, _ := time.Now().Date()
	return &Month{
		current: month.String(),
	}
}

func (m *Month) showMonth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, m)
}

func encodeJSON(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
