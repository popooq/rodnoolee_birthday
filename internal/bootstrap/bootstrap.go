package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/popooq/rodnoolee_birthday/internal/config"
	"github.com/popooq/rodnoolee_birthday/internal/repository"
	"github.com/popooq/rodnoolee_birthday/internal/repository/mongorepo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Bootstrap() error {
	ctx := context.Background()
	cfg := config.New()
	log.Println("конфиг подсосался")
	repo, err := mongorepo.New(ctx, cfg.DBAddress)
	if err != nil {
		log.Printf("ошибка при создании репки bootstrap.go: %s", err)
		return err
	}
	log.Println("репка создалась")

	err = repoTesting(repo)
	if err != nil {

		log.Printf("Ошибка при тестировании репки bootstrap.go: %s", err)
		return err
	}
	log.Println("репка затестилась")

	// hndlr := handler.New(repo)

	// rt := chi.NewRouter()
	// rt.Mount("/", Route())

	// return http.ListenAndServe(cfg.Address, rt)
	return nil
}

// func Route() *chi.Mux {
// 	r := chi.NewRouter()

// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.RealIP)
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)

// 	r.Post("/insertbday", addbirthday.InsertBirthday)
// 	r.Get("/getbday", allbirthdays.AllBirthdays)
// 	r.Post("/updatebday", updatebirthday.UpdateBirthday)
// 	return r
// }

func repoTesting(repo repository.UserRepo) error {

	err := repo.DeleteAllVeryDangerous()
	if err != nil {
		return fmt.Errorf("не удалось удалить все записи из базы %e", err)
	}

	var shortForm = time.DateOnly
	birthdays := make([]repository.Birthday, 5)

	for k, v := range birthdays {
		v.ID = fmt.Sprintf("%d", k)
		dt, _ := time.Parse(shortForm, fmt.Sprintf("%d, %d, %d", k, k, k))
		mdt := primitive.NewDateTimeFromTime(dt)
		v.Birthday = mdt
		v.Rodnoolya = fmt.Sprintf("rodnoolya №%d", k)
		err := repo.InsertBirthday(v)
		if err != nil {
			return err
		}
	}

	birthdays[1] = repository.Birthday{
		ID:        "2",
		Rodnoolya: "edited rodnoolya",
	}
	err = repo.UpdateBirthday(birthdays[1])
	if err != nil {
		return err
	}

	allbirthdays, err := repo.GetAllBirthdays()
	if err != nil {
		return err
	}

	fmt.Printf("all birthdays: %v", allbirthdays)

	err = repo.DeleteAllVeryDangerous()
	if err != nil {
		return fmt.Errorf("не удалось удалить все записи из базы %e", err)
	}

	return nil
}
