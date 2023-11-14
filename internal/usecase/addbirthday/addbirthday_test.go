package addbirthday

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/popooq/rodnoolee_birthday/internal/domain"
	"github.com/popooq/rodnoolee_birthday/internal/repository"
	mockrepo "github.com/popooq/rodnoolee_birthday/internal/repository/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_addBirthday_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		repo *mockrepo.MockUserRepo
	}
	type args struct {
		message domain.TgMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "positive test",
			fields: fields{
				repo: mockrepo.NewMockUserRepo(ctrl),
			},
			args: args{
				message: domain.TgMessage{
					MessageText: "2001-03-19",
					Username:    "Popooq",
				},
			},
			wantErr: false,
		},
		{name: "negative test",
			fields: fields{
				repo: mockrepo.NewMockUserRepo(ctrl),
			},
			args: args{
				message: domain.TgMessage{
					MessageText: "03-19-2001",
					Username:    "Popooq",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.repo.EXPECT().InsertBirthday(repository.Birthday{
				Birthday:  timeparser("2001-03-19"),
				Rodnoolya: "Popooq",
			}).Return(nil).AnyTimes()
			a := &addBirthday{
				repo: tt.fields.repo,
			}
			if err := a.Handle(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("addBirthday.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func timeparser(date string) primitive.DateTime {
	dt, _ := time.Parse(shortForm, date)
	mdt := primitive.NewDateTimeFromTime(dt)
	return mdt
}
