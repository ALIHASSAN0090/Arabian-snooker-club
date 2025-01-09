package Dao

// // CreateMatch creates a new match record in the database
// func (d *DaoImp) CreateMatch(match *models.Matches) (*models.Matches, error) {
// 	query := `INSERT INTO matches (match_name, per_match_price, created_at, updated_at, created_by, updated_by) VALUES ($1, $2, NOW(), NOW(), $3, $4) RETURNING id`
// 	err := d.db.QueryRow(query, match.MatchName, match.PerMatchPrice, match.CreatedBy, match.UpdatedBy).Scan(&match.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return match, nil
// }

// // GetMatchByID retrieves a match by its ID
// func (d *DaoImp) GetMatchByID(id int64) (*models.Matches, error) {
// 	var match models.Matches
// 	query := `SELECT id, match_name, per_match_price, created_at, updated_at, created_by, updated_by FROM matches WHERE id = $1`
// 	err := d.db.QueryRow(query, id).Scan(&match.ID, &match.MatchName, &match.PerMatchPrice, &match.CreatedAt, &match.UpdatedAt, &match.CreatedBy, &match.UpdatedBy)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.New("match not found")
// 		}
// 		return nil, err
// 	}
// 	return &match, nil
// }

// // UpdateMatch updates an existing match record
// func (d *DaoImp) UpdateMatch(match *models.Matches) error {
// 	query := `UPDATE matches SET match_name = $1, per_match_price = $2, updated_at = NOW(), updated_by = $3 WHERE id = $4`
// 	_, err := d.db.Exec(query, match.MatchName, match.PerMatchPrice, match.UpdatedBy, match.ID)
// 	return err
// }

// // DeleteMatch removes a match record from the database
// func (d *DaoImp) DeleteMatch(id int64) error {
// 	query := `DELETE FROM matches WHERE id = $1`
// 	_, err := d.db.Exec(query, id)
// 	return err
// }

// // CreateMatchesPlayed creates a new matches_played record
// func (d *DaoImp) CreateMatchesPlayed(mp *models.MatchesPlayed) (*models.MatchesPlayed, error) {
// 	query := `INSERT INTO matches_played (user_id, price_id, status, played_at, paid_at, paid_by, no_of_matches, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING id`
// 	err := d.db.QueryRow(query, mp.UserID, mp.PriceID, mp.Status, mp.PlayedAt, mp.PaidAt, mp.PaidBy, mp.NoOfMatches).Scan(&mp.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return mp, nil
// }

// // GetMatchesPlayedByID retrieves a matches_played record by ID
// func (d *DaoImp) GetMatchesPlayedByID(id int64) (*models.MatchesPlayed, error) {
// 	var mp models.MatchesPlayed
// 	query := `SELECT id, user_id, price_id, status, played_at, paid_at, paid_by, no_of_matches, created_at, updated_at FROM matches_played WHERE id = $1`
// 	err := d.db.QueryRow(query, id).Scan(&mp.ID, &mp.UserID, &mp.PriceID, &mp.Status, &mp.PlayedAt, &mp.PaidAt, &mp.PaidBy, &mp.NoOfMatches, &mp.CreatedAt, &mp.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &mp, nil
// }

// // UpdateMatchesPlayed updates an existing matches_played record
// func (d *DaoImp) UpdateMatchesPlayed(mp *models.MatchesPlayed) error {
// 	query := `UPDATE matches_played SET user_id = $1, price_id = $2, status = $3, played_at = $4, paid_at = $5, paid_by = $6, no_of_matches = $7, updated_at = NOW() WHERE id = $8`
// 	_, err := d.db.Exec(query, mp.UserID, mp.PriceID, mp.Status, mp.PlayedAt, mp.PaidAt, mp.PaidBy, mp.NoOfMatches, mp.ID)
// 	return err
// }

// // DeleteMatchesPlayed removes a matches_played record
// func (d *DaoImp) DeleteMatchesPlayed(id int64) error {
// 	query := `DELETE FROM matches_played WHERE id = $1`
// 	_, err := d.db.Exec(query, id)
// 	return err
// }

// // CreateCenturies creates a new centuries record
// func (d *DaoImp) CreateCenturies(c *models.Centuries) (*models.Centuries, error) {
// 	query := `INSERT INTO centuries (price_per_minute, created_at, updated_at) VALUES ($1, NOW(), NOW()) RETURNING id`
// 	err := d.db.QueryRow(query, c.PricePerMinute).Scan(&c.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return c, nil
// }

// // GetCenturiesByID retrieves a centuries record by ID
// func (d *DaoImp) GetCenturiesByID(id int64) (*models.Centuries, error) {
// 	var c models.Centuries
// 	query := `SELECT id, price_per_minute, created_at, updated_at FROM centuries WHERE id = $1`
// 	err := d.db.QueryRow(query, id).Scan(&c.ID, &c.PricePerMinute, &c.CreatedAt, &c.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &c, nil
// }

// // UpdateCenturies updates an existing centuries record
// func (d *DaoImp) UpdateCenturies(c *models.Centuries) error {
// 	query := `UPDATE centuries SET price_per_minute = $1, updated_at = NOW() WHERE id = $2`
// 	_, err := d.db.Exec(query, c.PricePerMinute, c.ID)
// 	return err
// }

// // DeleteCenturies removes a centuries record
// func (d *DaoImp) DeleteCenturies(id int64) error {
// 	query := `DELETE FROM centuries WHERE id = $1`
// 	_, err := d.db.Exec(query, id)
// 	return err
// }

// // CreateCenturiesPlayed creates a new centuries_played record
// func (d *DaoImp) CreateCenturiesPlayed(cp *models.CenturiesPlayed) (*models.CenturiesPlayed, error) {
// 	query := `INSERT INTO centuries_played (user_id, minutes_played, total_price, status, created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id`
// 	err := d.db.QueryRow(query, cp.UserID, cp.MinutesPlayed, cp.TotalPrice, cp.Status, cp.CreatedBy, cp.UpdatedBy).Scan(&cp.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return cp, nil
// }

// // GetCenturiesPlayedByID retrieves a centuries_played record by ID
// func (d *DaoImp) GetCenturiesPlayedByID(id int64) (*models.CenturiesPlayed, error) {
// 	var cp models.CenturiesPlayed
// 	query := `SELECT id, user_id, minutes_played, total_price, status, created_by, updated_by FROM centuries_played WHERE id = $1`
// 	err := d.db.QueryRow(query, id).Scan(&cp.ID, &cp.UserID, &cp.MinutesPlayed, &cp.TotalPrice, &cp.Status, &cp.CreatedBy, &cp.UpdatedBy)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &cp, nil
// }

// // UpdateCenturiesPlayed updates an existing centuries_played record
// func (d *DaoImp) UpdateCenturiesPlayed(cp *models.CenturiesPlayed) error {
// 	query := `UPDATE centuries_played SET user_id = $1, minutes_played = $2, total_price = $3, status = $4, updated_by = $5, updated_at = NOW() WHERE id = $6`
// 	_, err := d.db.Exec(query, cp.UserID, cp.MinutesPlayed, cp.TotalPrice, cp.Status, cp.UpdatedBy, cp.ID)
// 	return err
// }

// // DeleteCenturiesPlayed removes a centuries_played record
// func (d *DaoImp) DeleteCenturiesPlayed(id int64) error {
// 	query := `DELETE FROM centuries_played WHERE id = $1`
// 	_, err := d.db.Exec(query, id)
// 	return err
// }
