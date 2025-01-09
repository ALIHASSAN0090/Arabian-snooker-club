package models

import "time"

type MatchesPlayed struct {
	ID          int64      `db:"id" json:"id"`
	UserID      int64      `db:"user_id" json:"user_id"`
	PriceID     int64      `db:"price_id" json:"price_id"`
	Status      int16      `db:"status" json:"status"`
	PlayedAt    time.Time  `db:"played_at" json:"played_at"`
	PaidAt      *time.Time `db:"paid_at" json:"paid_at,omitempty"`
	PaidBy      *int64     `db:"paid_by" json:"paid_by,omitempty"`
	NoOfMatches int        `db:"no_of_matches" json:"no_of_matches"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type Users struct {
	ID       int     `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Email    *string `db:"email" json:"email,omitempty"`
	Password string  `db:"password" json:"password"`
	Role     string  `db:"role" json:"role"`
	Active   bool    `db:"active" json:"active"`
}

type Matches struct {
	ID            int64     `db:"id" json:"id"`
	MatchName     string    `db:"match_name" json:"match_name"`
	PerMatchPrice float64   `db:"per_match_price" json:"per_match_price"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy     int64     `db:"created_by" json:"created_by"`
	UpdatedBy     int64     `db:"updated_by" json:"updated_by"`
}

type Centuries struct {
	ID             int64      `db:"id" json:"id"`
	PricePerMinute float64    `db:"price_per_minute" json:"price_per_minute"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type CenturiesPlayed struct {
	ID            int64   `db:"id" json:"id"`
	UserID        int64   `db:"user_id" json:"user_id"`
	MinutesPlayed int     `db:"minutes_played" json:"minutes_played"`
	TotalPrice    float64 `db:"total_price" json:"total_price"`
	Status        int16   `db:"status" json:"status"`
	CreatedBy     int64   `db:"created_by" json:"created_by"`
	UpdatedBy     *int64  `db:"updated_by" json:"updated_by,omitempty"`
	DeletedBy     *int64  `db:"deleted_by" json:"deleted_by,omitempty"`
}
