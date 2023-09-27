package allbirthdays

import (
	"fmt"
	"net/http"
)

func AllBirthdays(w http.ResponseWriter, r *http.Request) {
	birthdays, err := repository.GetAllBirthdays()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("bruh"))
	}
	listOfBirthdays := fmt.Sprintf("%+v", birthdays)

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(listOfBirthdays))
}
