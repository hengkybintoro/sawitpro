package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
// func (s *Server) GetHello(ctx echo.Context, params generated.GetHelloParams) error {
// 	var resp generated.HelloResponse
// 	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
// 	return ctx.JSON(http.StatusOK, resp)
// }

// Add estate endpoint
func (h *Server) AddEstate(c echo.Context) error {
	start := time.Now()
	req := model.RequestEstate{}
	resp := model.Response{}

	if err := c.Bind(&req); err != nil {
		resp.Message = "invalid request"
		return c.JSON(http.StatusBadRequest, resp)
	}
	if req.Width <= 0 || req.Length <= 0 || req.Width > 50000 || req.Length > 50000 {
		resp.Message = "invalid dimensions"
		return c.JSON(http.StatusBadRequest, resp)
	}

	id, err := h.Service.AddEstate(req.Width, req.Length)
	if err != nil {
		resp.Message = "failed to create estate"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	respData := model.ResponseWithID{
		ID: id.String(),
	}
	resp.Data = respData
	resp.Message = "success"

	resp.Header.ProcessTime = time.Since(start).Seconds() * 1000
	return c.JSON(http.StatusCreated, resp)
}

// Add tree endpoint
func (h *Server) AddTree(c echo.Context) error {
	start := time.Now()
	req := model.RequestTree{}
	resp := model.Response{}

	estateID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.Message = "invalid estate ID"
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Bind(&req); err != nil {
		resp.Message = "invalid request"
		return c.JSON(http.StatusBadRequest, resp)
	}
	if req.X <= 0 || req.Y <= 0 || req.Height <= 0 || req.Height > 30 {
		resp.Message = "invalid tree parameters"
		return c.JSON(http.StatusBadRequest, resp)
	}

	id, err := h.Service.AddTree(estateID, req.X, req.Y, req.Height)
	if err != nil {
		resp.Message = "failed to add tree"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	respData := model.ResponseWithID{
		ID: id.String(),
	}
	resp.Data = respData
	resp.Message = "success"

	resp.Header.ProcessTime = time.Since(start).Seconds() * 1000
	return c.JSON(http.StatusCreated, resp)
}

// Get stats endpoint
func (h *Server) GetEstateStats(c echo.Context) error {
	start := time.Now()
	resp := model.Response{}
	estateID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.Message = "invalid estate ID"
		return c.JSON(http.StatusBadRequest, resp)
	}

	count, maxHeight, minHeight, medianHeight, err := h.Service.GetEstateStats(estateID)
	if err != nil {
		resp.Message = "failed to get estate stats"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	respData := model.ResponseWithStats{
		TreeCount:    count,
		MaxHeight:    maxHeight,
		MinHeight:    minHeight,
		MedianHeight: medianHeight,
	}
	resp.Data = respData
	resp.Message = "success"

	resp.Header.ProcessTime = time.Since(start).Seconds() * 1000
	return c.JSON(http.StatusOK, resp)
}

// Get drone-plan endpoint
func (h *Server) GetDronePlan(c echo.Context) error {
	start := time.Now()
	resp := model.Response{}

	estateID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.Message = "invalid estate ID"
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Get drone-plan with max_distance
	maxDistanceStr := c.QueryParam("max_distance")
	if maxDistanceStr != "" {
		maxDistance, err := strconv.Atoi(maxDistanceStr)
		if err != nil || maxDistance <= 0 {
			resp.Message = "invalid max_distance"
			return c.JSON(http.StatusBadRequest, resp)
		}

		distance, rest, err := h.Service.GetDronePlanMaxDistance(estateID, maxDistance)
		if err != nil {
			if err.Error() == "estate not found" {
				resp.Message = "estate not found"
				return c.JSON(http.StatusNotFound, resp)
			}
			resp.Message = "failed to calculate drone plan distance"
			return c.JSON(http.StatusInternalServerError, resp)
		}

		respData := model.ResponseWithDistanceAndCoordinate{
			ResponseWithDistance: model.ResponseWithDistance{
				Distance: distance,
			},
			Rest: rest,
		}
		resp.Data = respData
		resp.Message = "success"

		return c.JSON(http.StatusOK, resp)
	}

	distance, err := h.Service.GetDronePlanDistance(estateID)
	if err != nil {
		if err.Error() == "estate not found" {
			resp.Message = "estate not found"
			return c.JSON(http.StatusNotFound, resp)
		}

		resp.Message = "failed to calculate drone plan distance"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	respData := model.ResponseWithDistance{
		Distance: distance,
	}
	resp.Data = respData
	resp.Message = "success"

	resp.Header.ProcessTime = time.Since(start).Seconds() * 1000
	return c.JSON(http.StatusOK, resp)
}
