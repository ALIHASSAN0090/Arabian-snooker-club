package Dao

import "arabian-snooker/models"

func (s *DaoImp) CreateRates(req models.CreateMatch) (models.Matches, error) {

	var data models.Matches
	query := `INSERT INTO matches(match_name, per_match_price, created_by, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id, match_name, per_match_price, created_by, created_at`
	err := s.db.QueryRow(query, req.MatchName, req.PerMatchPrice, req.CreatedBy).Scan(&data.ID, &data.MatchName, &data.PerMatchPrice, &data.CreatedBy, &data.CreatedAt)
	if err != nil {
		return models.Matches{}, err
	}

	return data, nil
}

func (s *DaoImp) UpdateRates(req models.UpdateMatch) (models.Matches, error) {
	var data models.Matches
	query := `UPDATE matches SET match_name = $1, per_match_price = $2, updated_by = $3, updated_at = NOW() WHERE id = $4 RETURNING id, match_name, per_match_price, created_by, created_at`
	err := s.db.QueryRow(query, req.MatchName, req.PerMatchPrice, req.UpdatedBy, req.ID).Scan(&data.ID, &data.MatchName, &data.PerMatchPrice, &data.CreatedBy, &data.CreatedAt)
	if err != nil {
		return models.Matches{}, err
	}

	return data, nil
}
