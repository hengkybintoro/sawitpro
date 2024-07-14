package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func TestGetHello(t *testing.T) {

}

func TestServer_AddEstate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()

	serviceMock := service.NewMockServiceInterface(ctrl)
	serviceMock.EXPECT().AddEstate(gomock.Any(), gomock.Any()).Return(id, nil).AnyTimes()

	addEstateReq, _ := http.NewRequest("POST", "/estate", nil)
	w := httptest.NewRecorder()

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		h       *Server
		args    args
		wantErr bool
	}{
		{
			name: "1. normal case",
			h: &Server{
				Service: serviceMock,
			},
			args: args{
				c: echo.New().NewContext(addEstateReq, w),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.AddEstate(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.AddEstate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_AddTree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()

	serviceMock := service.NewMockServiceInterface(ctrl)
	serviceMock.EXPECT().AddTree(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(id, nil).AnyTimes()

	addTreeReq, _ := http.NewRequest("POST", "/tree", nil)
	w := httptest.NewRecorder()

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		h       *Server
		args    args
		wantErr bool
	}{
		{
			name: "1. normal case",
			h: &Server{
				Service: serviceMock,
			},
			args: args{
				c: echo.New().NewContext(addTreeReq, w),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.args.c.SetParamNames("id")
		tt.args.c.SetParamValues(uuid.New().String())

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.AddTree(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.AddTree() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_GetEstateStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	count := 10
	maxHeight := 20
	minHeight := 5
	medianHeight := 12

	serviceMock := service.NewMockServiceInterface(ctrl)
	serviceMock.EXPECT().GetEstateStats(gomock.Any()).Return(count, maxHeight, minHeight, medianHeight, nil)

	getStatsReq, _ := http.NewRequest("POST", "/stats", nil)
	w := httptest.NewRecorder()

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		h       *Server
		args    args
		wantErr bool
	}{
		{
			name: "1. normal case",
			h: &Server{
				Service: serviceMock,
			},
			args: args{
				c: echo.New().NewContext(getStatsReq, w),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.args.c.SetParamNames("id")
		tt.args.c.SetParamValues(uuid.New().String())

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetEstateStats(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.GetEstateStats() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_GetDronePlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	distance := 100
	coordinate := model.Coordinate{
		X: 7,
		Y: 2,
	}

	serviceMock := service.NewMockServiceInterface(ctrl)
	serviceMock.EXPECT().GetDronePlanDistance(gomock.Any()).Return(distance, nil)
	serviceMock.EXPECT().GetDronePlanMaxDistance(gomock.Any(), gomock.Any()).Return(distance, coordinate, nil)

	getDronePlanReq, _ := http.NewRequest("GET", "/droneplan", nil)
	getDronePlanWithMaxDistanceReq, _ := http.NewRequest("GET", "/droneplan?max_distance=100", nil)
	w := httptest.NewRecorder()

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		h       *Server
		args    args
		wantErr bool
	}{
		{
			name: "1. normal case",
			h: &Server{
				Service: serviceMock,
			},
			args: args{
				c: echo.New().NewContext(getDronePlanReq, w),
			},
			wantErr: false,
		},
		{
			name: "2. normal case with max_distance",
			h: &Server{
				Service: serviceMock,
			},
			args: args{
				c: echo.New().NewContext(getDronePlanWithMaxDistanceReq, w),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.args.c.SetParamNames("id")
		tt.args.c.SetParamValues(uuid.New().String())

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetDronePlan(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.GetDronePlan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
