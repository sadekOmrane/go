package dict
	
import (
	"encoding/json"
	"net/http"
)
type Dict struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type Server struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Dict `json:"data"`
}

func New() *Server {
	s := Server{
		Data: []Dict{},
	}
	return &s
}

func (s *Server) PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dict Dict
	if err := json.NewDecoder(r.Body).Decode(&dict); err != nil {
		http.Error(w, "Error reading your dict.", http.StatusBadRequest)
		return
	}

	s.Data = append(s.Data, dict)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dict); err != nil {
		http.Error(w, "Error encoding JSON.", http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")

	if word == "" {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.Data); err != nil {
			http.Error(w, "Error encoding JSON.", http.StatusInternalServerError)
			return
		}
		return
	}

	var foundItem Dict
	for _, item := range s.Data {
		if item.Word == word {
			foundItem = item
			break
		}
	}

	if foundItem.Word == "" {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(foundItem); err != nil {
		http.Error(w, "Error encoding JSON.", http.StatusInternalServerError)
		return
	}
}

func (s *Server) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing 'word' parameter", http.StatusBadRequest)
		return
	}

	var updatedDict Dict
	if err := json.NewDecoder(r.Body).Decode(&updatedDict); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	var foundIndex = -1
	for i, item := range s.Data {
		if item.Word == word {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	s.Data[foundIndex] = updatedDict

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.Data[foundIndex]); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing 'word' parameter", http.StatusBadRequest)
		return
	}

	var foundIndex = -1
	for i, item := range s.Data {
		if item.Word == word {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	s.Data = append(s.Data[:foundIndex], s.Data[foundIndex+1:]...)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Item deleted successfully",
	}); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
