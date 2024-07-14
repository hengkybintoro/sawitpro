package repository

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func TestRepository_AddEstate(t *testing.T) {
	conn, sqlm, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' when connecting to database", err)
	}
	defer conn.Close()

	sqlxDB := sqlx.NewDb(conn, "sqlmock")

	type args struct {
		width  int
		length int
	}

	tests := []struct {
		name    string
		r       *Repository
		args    args
		wantErr bool
	}{
		{
			name: "1. normal case",
			r: &Repository{
				Db: sqlxDB,
			},
			args: args{
				width:  10,
				length: 10,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlm.ExpectExec("INSERT INTO estates").
				WithArgs(sqlmock.AnyArg(), tt.args.width, tt.args.length).
				WillReturnResult(sqlmock.NewResult(1, 1))

			_, err := tt.r.AddEstate(tt.args.width, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.AddEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepository_GetEstate(t *testing.T) {
	conn, sqlm, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' when connecting to database", err)
	}
	defer conn.Close()

	sqlxDB := sqlx.NewDb(conn, "sqlmock")

	id := uuid.New()

	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		r       *Repository
		args    args
		want    *model.Estate
		wantErr bool
	}{
		{
			name: "1. normal test",
			r: &Repository{
				Db: sqlxDB,
			},
			args: args{
				id: id,
			},
			want: &model.Estate{
				ID:     id,
				Width:  10,
				Length: 10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		rows := sqlmock.NewRows([]string{"id", "width", "length"}).AddRow(id, 10, 10)
		sqlm.ExpectQuery("SELECT").WillReturnRows(rows)

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetEstate(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetEstate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_AddTree(t *testing.T) {
	conn, sqlm, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' when connecting to database", err)
	}
	defer conn.Close()

	sqlxDB := sqlx.NewDb(conn, "sqlmock")

	estateID := uuid.New()

	type args struct {
		estateID uuid.UUID
		x        int
		y        int
		height   int
	}
	tests := []struct {
		name    string
		r       *Repository
		args    args
		wantErr bool
	}{
		{
			name: "1. normal case",
			r: &Repository{
				Db: sqlxDB,
			},
			args: args{
				estateID: estateID,
				x:        10,
				y:        10,
				height:   30,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		sqlm.ExpectExec("INSERT INTO trees").
			WithArgs(sqlmock.AnyArg(), estateID, tt.args.x, tt.args.y, tt.args.height).
			WillReturnResult(sqlmock.NewResult(1, 1))

		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.AddTree(tt.args.estateID, tt.args.x, tt.args.y, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.AddTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepository_GetTreesByEstate(t *testing.T) {
	conn, sqlm, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' when connecting to database", err)
	}
	defer conn.Close()

	sqlxDB := sqlx.NewDb(conn, "sqlmock")

	id := uuid.New()
	estateID := uuid.New()

	type args struct {
		estateID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *Repository
		args    args
		want    []model.Tree
		wantErr bool
	}{
		{
			name: "1. normal case",
			r: &Repository{
				Db: sqlxDB,
			},
			args: args{
				estateID: estateID,
			},
			want: []model.Tree{
				{
					ID:       id,
					EstateID: estateID,
					X:        10,
					Y:        10,
					Height:   30,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		rows := sqlmock.NewRows([]string{"id", "estate_id", "x", "y", "height"}).AddRow(id, estateID, 10, 10, 30)
		sqlm.ExpectQuery("SELECT").WillReturnRows(rows)

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetTreesByEstate(tt.args.estateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetTreesByEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetTreesByEstate() = %v, want %v", got, tt.want)
			}
		})
	}
}
