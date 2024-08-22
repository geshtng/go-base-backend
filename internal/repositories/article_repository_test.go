package repositories

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/models"
)

func TestNewArticleRepository(t *testing.T) {
	dbConn, _, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)
	articleRepo := NewArticleRepository(&ArticleRepoConfig{DB: db})

	assert.NotNil(t, articleRepo)
}

func Test_articleRepository_GetAll(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	tests := []struct {
		name    string
		a       *articleRepository
		mock    func()
		want    []*models.Article
		wantErr bool
	}{
		{
			name: "Success",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
						AddRow(1, "title", "description"))
			},
			want: []*models.Article{
				{
					ID:          1,
					Title:       "title",
					Description: "description",
				},
			},
			wantErr: false,
		},
		{
			name: "Error DB Find",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WillReturnError(errors.New("something error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.a.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleRepository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleRepository_GetByID(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		a       *articleRepository
		mock    func()
		args    args
		want    *models.Article
		wantErr bool
	}{
		{
			name: "Success",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
						AddRow(1, "title", "description"))
			},
			args: args{
				id: 1,
			},
			want: &models.Article{
				ID:          1,
				Title:       "title",
				Description: "description",
			},
			wantErr: false,
		},
		{
			name: "Error DB First",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(1).
					WillReturnError(errors.New("something error"))
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.a.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleRepository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleRepository_Insert(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	type args struct {
		article models.Article
	}
	tests := []struct {
		name    string
		a       *articleRepository
		mock    func()
		args    args
		want    *models.Article
		wantErr bool
	}{
		{
			name: "Success",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "articles" (.+) VALUES (.+)`).
					WithArgs("title", "description").
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
						AddRow(1, "title", "description"))
				mock.ExpectCommit()
			},
			args: args{
				article: models.Article{
					Title:       "title",
					Description: "description",
				},
			},
			want: &models.Article{
				ID:          1,
				Title:       "title",
				Description: "description",
			},
			wantErr: false,
		},
		{
			name: "Error DB Create",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "articles" (.+) VALUES (.+)`).
					WithArgs("title", "description").
					WillReturnError(errors.New("something error"))
				mock.ExpectCommit()
			},
			args: args{
				article: models.Article{
					Title:       "title",
					Description: "description",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.a.Insert(tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleRepository.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleRepository_Update(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	type args struct {
		id      int64
		article models.Article
	}
	tests := []struct {
		name    string
		a       *articleRepository
		mock    func()
		args    args
		want    *models.Article
		wantErr bool
	}{
		{
			name: "Success",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`UPDATE "articles" SET (.+) WHERE id = (.+)`).
					WithArgs("title update", "description update", 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
						AddRow(1, "title update", "description update"))
				mock.ExpectCommit()
			},
			args: args{
				id: 1,
				article: models.Article{
					Title:       "title update",
					Description: "description update",
				},
			},
			want: &models.Article{
				ID:          1,
				Title:       "title update",
				Description: "description update",
			},
			wantErr: false,
		},
		{
			name: "Error DB Update",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`UPDATE "articles" SET (.+) WHERE id = (.+)`).
					WithArgs("title update", "description update", 1).
					WillReturnError(errors.New("something error"))
				mock.ExpectCommit()
			},
			args: args{
				id: 1,
				article: models.Article{
					Title:       "title update",
					Description: "description update",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Rows Affected Zero",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "articles" SET (.+) WHERE id = (.+)`).
					WithArgs("title update", "description update", 1).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectCommit()
			},
			args: args{
				id: 1,
				article: models.Article{
					Title:       "title update",
					Description: "description update",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.a.Update(tt.args.id, tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleRepository_Delete(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		a       *articleRepository
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE")).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "Error DB Delete",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE")).
					WithArgs(1).
					WillReturnError(errors.New("something error"))
				mock.ExpectCommit()
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
		{
			name: "Rows Affected Zero",
			a: &articleRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE")).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectCommit()
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			if err := tt.a.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
