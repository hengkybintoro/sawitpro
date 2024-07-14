package service

import (
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

type Plot struct {
	X, Y int // Coordinates
}

func (s *Service) AddEstate(width, length int) (uuid.UUID, error) {
	if width <= 0 || length <= 0 {
		return uuid.Nil, errors.New("invalid value")
	}

	return s.repo.AddEstate(width, length)
}

func (s *Service) AddTree(estateID uuid.UUID, x, y, height int) (uuid.UUID, error) {
	if x <= 0 || y <= 0 || height < 1 || height > 30 {
		return uuid.Nil, errors.New("invalid value")
	}

	return s.repo.AddTree(estateID, x, y, height)
}

func (s *Service) GetEstateStats(estateID uuid.UUID) (int, int, int, int, error) {
	trees, err := s.repo.GetTreesByEstate(estateID)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	if len(trees) == 0 {
		return 0, 0, 0, 0, nil
	}

	count := len(trees)
	maxHeight := trees[0].Height
	minHeight := trees[0].Height
	heights := make([]int, count)

	// Check tallest and shortest
	for i, tree := range trees {
		if tree.Height > maxHeight {
			maxHeight = tree.Height
		}
		if tree.Height < minHeight {
			minHeight = tree.Height
		}
		heights[i] = tree.Height
	}

	// Calculate median
	sort.Ints(heights)
	medianHeight := heights[count/2]
	if count%2 == 0 {
		medianHeight = (heights[count/2-1] + heights[count/2]) / 2
	}

	return count, maxHeight, minHeight, medianHeight, nil
}

func (s *Service) GetDronePlanDistance(estateID uuid.UUID) (int, error) {
	// Fetch estate dimensions
	estate, err := s.repo.GetEstate(estateID)
	if err != nil {
		return 0, err
	}
	if estate == nil {
		return 0, fmt.Errorf("estate not found")
	}

	// Fetch trees in estate
	trees, err := s.repo.GetTreesByEstate(estateID)
	if err != nil {
		return 0, err
	}

	// Calculate total distance
	totalDistance := 0

	// Determine the dimensions of the estate in plots
	widthPlots := estate.Width
	lengthPlots := estate.Length

	// Calculate horizontal distance (east-west)
	totalDistance += (widthPlots - 1) * 10 // Each plot is 10 meters apart horizontally

	// Calculate vertical distance (north-south)
	totalDistance += (lengthPlots - 1) * 10 // Each plot is 10 meters apart vertically

	// Calculate the additional vertical distance for monitoring trees
	for _, tree := range trees {
		// Calculate the vertical distance from the ground to 1 meter above the tree
		totalDistance += int(math.Abs(float64(tree.Height - 1)))
	}

	return totalDistance, nil
}

// Bonus : calculate drone-plan with max_distance
func (s *Service) GetDronePlanMaxDistance(estateID uuid.UUID, maxDistance int) (int, model.Coordinate, error) {
	coordinate := model.Coordinate{}
	estate, err := s.repo.GetEstate(estateID)
	if err != nil {
		return 0, coordinate, err
	}
	if estate == nil {
		return 0, coordinate, fmt.Errorf("estate not found")
	}

	trees, err := s.repo.GetTreesByEstate(estateID)
	if err != nil {
		return 0, coordinate, err
	}

	treesPlot := mapTreeToPlot(trees)

	height := estate.Length
	width := estate.Width
	totalDistance := 0
	x, y := 1, 1

	for y <= height {
		for direction := 1; direction >= -1; direction -= 2 {
			for x = 1; x <= width; x++ {
				// Calculate the vertical distance to the current plot
				if treeHeight, exists := treesPlot[Plot{X: x, Y: y}]; exists {
					totalDistance += treeHeight + 1 // Add 1 meter above the tree
				} else {
					totalDistance += 1 // Add 1 meter above the ground
				}

				if totalDistance > maxDistance {
					coordinate.X = x
					coordinate.Y = y

					return totalDistance, coordinate, nil
				}

				totalDistance += 1 // Monitoring the plot

				// Move to the next plot
				if x < width {
					if totalDistance+10 > maxDistance {
						coordinate.X = x
						coordinate.Y = y

						return totalDistance, coordinate, nil
					}
					totalDistance += 10
				}
			}

			y++
			if y > height {
				break
			}

			for x = width; x >= 1; x-- {
				// Calculate the vertical distance to the current plot
				if treeHeight, exists := treesPlot[Plot{X: x, Y: y}]; exists {
					totalDistance += treeHeight + 1 // Add 1 meter above the tree
				} else {
					totalDistance += 1 // Add 1 meter above the ground
				}

				if totalDistance > maxDistance {
					coordinate.X = x
					coordinate.Y = y

					return totalDistance, coordinate, nil
				}

				totalDistance += 1 // Monitoring the plot

				// Move to the next plot
				if x > 1 {
					if totalDistance+10 > maxDistance {
						coordinate.X = x
						coordinate.Y = y

						return totalDistance, coordinate, nil
					}
					totalDistance += 10
				}
			}
		}
		y++
	}

	// If the drone finishes the path without reaching the maxDistance, return the last point
	coordinate.X = x
	coordinate.Y = y

	return totalDistance, coordinate, nil
}

// Map trees from database to plot
func mapTreeToPlot(trees []model.Tree) map[Plot]int {
	result := make(map[Plot]int)

	for _, tree := range trees {
		plot := Plot{
			X: tree.X,
			Y: tree.Y,
		}
		result[plot] = tree.Height
	}

	return result
}
