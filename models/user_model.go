package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MatchDate        string             `json:"matchDate,omitempty" validate:"required"`
	MatchTime        string             `json:"matchTime,omitempty" validate:"required"`
	Cid              string             `json:"cId,omitempty" validate:"required"`
	MatchID          string             `json:"matchId,omitempty" validate:"required"`
	LeagueID         string             `json:"leagueId,omitempty" validate:"required"`
	HomeTeam         string             `json:"homeTeam,omitempty" validate:"required"`
	AwayTeam         string             `json:"awayTeam,omitempty" validate:"required"`
	HalfTimeScore    string             `json:"halfTimeScore,omitempty" validate:"required"`
	FullTimeScore    string             `json:"fullTimeScore,omitempty" validate:"required"`
	TotalGoals       int32              `json:"totalGoals,omitempty" validate:"required"`
	FirstHalfValue   string             `json:"firstHalfValue,omitempty" validate:"required"`
	SecondHalfValue  string             `json:"secondHalfValue,omitempty" validate:"required"`
	HomeWinOdd       float64            `json:"homeWinOdd,omitempty" validate:"required"`
	DrawOdd          float64            `json:"drawOdd,omitempty" validate:"required"`
	AwayWinOdd       float64            `json:"awayWinOdd,omitempty" validate:"required"`
	Over_2_5_odd     float64            `json:"over_2_5_odd,omitempty" validate:"required"`
	Under_2_5_odd    float64            `json:"under_2_5_odd,omitempty" validate:"required"`
	Over_1_5_odd     float64            `json:"over_1_5_odd,omitempty" validate:"required"`
	Under_1_5_odd    float64            `json:"under_1_5_odd,omitempty" validate:"required"`
	Over_2_5_status  int32              `json:"over_2_5_status,omitempty" validate:"required"`
	Under_2_5_status int32              `json:"under_2_5_status,omitempty" validate:"required"`
	Over_1_5_status  int32              `json:"over_1_5_status,omitempty" validate:"required"`
	Under_1_5_status int32              `json:"under_1_5_status,omitempty" validate:"required"`
}
