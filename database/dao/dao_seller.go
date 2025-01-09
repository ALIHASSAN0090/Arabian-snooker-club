package Dao

import (
	"arabian-snooker/models"
	"errors"
	"strconv"
)

func (s *DaoImp) CreateRates(req models.CreateMatch) (models.Matches, error) {

	var data models.Matches
	query := `INSERT INTO matches(match_name, per_match_price, created_by, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id, match_name, per_match_price, created_by, created_at`
	err := s.db.QueryRow(query, req.MatchName, req.PerMatchPrice, req.CreatedBy).Scan(&data.ID, &data.MatchName, &data.PerMatchPrice, &data.CreatedBy, &data.CreatedAt)
	if err != nil {
		return models.Matches{}, err
	}

	return data, nil
}

func (s *DaoImp) UpdateRate(req models.UpdateMatch) (models.Matches, error) {
	var data models.Matches
	if req.MatchName == "" && req.PerMatchPrice == 0 {
		return models.Matches{}, errors.New("either match name or per match price must be supplied")
	}

	query := `UPDATE matches SET `
	args := []interface{}{}
	argCounter := 1

	if req.MatchName != "" {
		query += `match_name = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.MatchName)
		argCounter++
	}

	if req.PerMatchPrice != 0 {
		query += `per_match_price = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.PerMatchPrice)
		argCounter++
	}

	query += `updated_by = $` + strconv.Itoa(argCounter) + `, updated_at = NOW() WHERE id = $` + strconv.Itoa(argCounter+1) + ` RETURNING id, match_name, per_match_price, created_by, created_at, updated_at, updated_by`
	args = append(args, req.UpdatedBy, req.ID)

	err := s.db.QueryRow(query, args...).Scan(&data.ID, &data.MatchName, &data.PerMatchPrice, &data.CreatedBy, &data.CreatedAt, &data.UpdatedAt, &data.UpdatedBy)
	if err != nil {
		return models.Matches{}, errors.New("error in updating rates: " + err.Error())
	}
	return data, nil
}
