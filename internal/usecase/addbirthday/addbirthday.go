package addbirthday

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/popooq/rodnoolee_birthday/internal/types"
)

func InsertBirthday(w http.ResponseWriter, r *http.Request) {
	var birthday types.Birthday

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error during ReadAll: %s", err)
	}

	err = json.Unmarshal(body, birthday)

	err = repository.InsertBirthday(birthday)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
}
