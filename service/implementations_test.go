package service

import (
	"reflect"
	"testing"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestService_AddEstate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()

	repoMock := repository.NewMockRepositoryInterface(ctrl)
	repoMock.EXPECT().AddEstate(gomock.Any(), gomock.Any()).Return(id, nil)

	type args struct {
		width  int
		length int
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "1. normal case",
			s: &Service{
				repo: repoMock,
			},
			args: args{
				width:  10,
				length: 10,
			},
			want:    id,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AddEstate(tt.args.width, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddEstate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddTree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()

	repoMock := repository.NewMockRepositoryInterface(ctrl)
	repoMock.EXPECT().AddTree(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(id, nil)

	type args struct {
		estateID uuid.UUID
		x        int
		y        int
		height   int
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "1. normal case",
			s: &Service{
				repo: repoMock,
			},
			args: args{
				estateID: uuid.New(),
				x:        10,
				y:        10,
				height:   30,
			},
			want:    id,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AddTree(tt.args.estateID, tt.args.x, tt.args.y, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetEstateStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	estateID := uuid.New()

	trees := []model.Tree{
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        10,
			Y:        10,
			Height:   30,
		},
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        15,
			Y:        15,
			Height:   20,
		},
	}

	repoMock := repository.NewMockRepositoryInterface(ctrl)
	repoMock.EXPECT().GetTreesByEstate(gomock.Any()).Return(trees, nil)

	type args struct {
		estateID uuid.UUID
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    int
		want1   int
		want2   int
		want3   int
		wantErr bool
	}{
		{
			name: "1. normal case",
			s: &Service{
				repo: repoMock,
			},
			args: args{
				estateID: estateID,
			},
			want:    2,
			want1:   30,
			want2:   20,
			want3:   25,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := tt.s.GetEstateStats(tt.args.estateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetEstateStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetEstateStats() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Service.GetEstateStats() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Service.GetEstateStats() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("Service.GetEstateStats() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func TestService_GetDronePlanDistance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	estateID := uuid.New()

	estate := model.Estate{
		ID:     estateID,
		Width:  20,
		Length: 20,
	}

	trees := []model.Tree{
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        10,
			Y:        10,
			Height:   30,
		},
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        15,
			Y:        15,
			Height:   20,
		},
	}

	repoMock := repository.NewMockRepositoryInterface(ctrl)
	repoMock.EXPECT().GetEstate(gomock.Any()).Return(&estate, nil)
	repoMock.EXPECT().GetTreesByEstate(gomock.Any()).Return(trees, nil)

	type args struct {
		estateID uuid.UUID
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "1. normal case",
			s: &Service{
				repo: repoMock,
			},
			args: args{
				estateID: estateID,
			},
			want:    428,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetDronePlanDistance(tt.args.estateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetDronePlanDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetDronePlanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetDronePlanMaxDistance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	estateID := uuid.New()

	estate := model.Estate{
		ID:     estateID,
		Width:  20,
		Length: 20,
	}

	trees := []model.Tree{
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        10,
			Y:        10,
			Height:   30,
		},
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        15,
			Y:        15,
			Height:   20,
		},
	}

	repoMock := repository.NewMockRepositoryInterface(ctrl)
	repoMock.EXPECT().GetEstate(gomock.Any()).Return(&estate, nil)
	repoMock.EXPECT().GetTreesByEstate(gomock.Any()).Return(trees, nil)

	type args struct {
		estateID    uuid.UUID
		maxDistance int
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    int
		want1   model.Coordinate
		wantErr bool
	}{
		{
			name: "1. normal case",
			s: &Service{
				repo: repoMock,
			},
			args: args{
				estateID:    estateID,
				maxDistance: 100,
			},
			want: 98,
			want1: model.Coordinate{
				X: 9,
				Y: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.GetDronePlanMaxDistance(tt.args.estateID, tt.args.maxDistance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetDronePlanMaxDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetDronePlanMaxDistance() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Service.GetDronePlanMaxDistance() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_mapTreeToPlot(t *testing.T) {
	estateID := uuid.New()
	trees := []model.Tree{
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        10,
			Y:        10,
			Height:   30,
		},
		{
			ID:       uuid.New(),
			EstateID: estateID,
			X:        15,
			Y:        15,
			Height:   20,
		},
	}

	type args struct {
		trees []model.Tree
	}
	tests := []struct {
		name string
		args args
		want map[Plot]int
	}{
		{
			name: "1. normal case",
			args: args{
				trees: trees,
			},
			want: map[Plot]int{
				{
					X: 10,
					Y: 10,
				}: 30,
				{
					X: 15,
					Y: 15,
				}: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapTreeToPlot(tt.args.trees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapTreeToPlot() = %v, want %v", got, tt.want)
			}
		})
	}
}
