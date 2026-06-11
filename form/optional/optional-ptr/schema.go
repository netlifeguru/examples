package main

import (
	"time"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/optional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeMinAge         = form.Code("min_age")
	CodeMaxAttempts    = form.Code("max_attempts")
	CodePriorityRange  = form.Code("priority_range")
	CodeAfterStart     = form.Code("after_start")
	CodeBeforeDeadline = form.Code("before_deadline")
	CodeBetweenWindow  = form.Code("between_window")
	CodeProfileTitle   = form.Code("profile_title_required")
)

type OptionalPtrRequest struct {
	Age      *int `json:"age"`
	Score    *int `json:"score"`
	Attempts *int `json:"attempts"`
	Limit    *int `json:"limit"`
	Rating   *int `json:"rating"`
	Priority *int `json:"priority"`

	StartAt         *time.Time `json:"start_at"`
	AvailableFrom   *time.Time `json:"available_from"`
	Deadline        *time.Time `json:"deadline"`
	ExpiresAt       *time.Time `json:"expires_at"`
	EventDate       *time.Time `json:"event_date"`
	ReservationDate *time.Time `json:"reservation_date"`

	Nickname     string   `json:"nickname"`
	Tags         []string `json:"tags"`
	ProfileBio   *string  `json:"profile_bio"`
	ProfileTitle string   `json:"profile_title"`
}

func OptionalPtrSchema() form.Schema[OptionalPtrRequest] {
	OptionalPtrForm := struct {
		Age      form.OptIntField[OptionalPtrRequest]
		Score    form.OptIntField[OptionalPtrRequest]
		Attempts form.OptIntField[OptionalPtrRequest]
		Limit    form.OptIntField[OptionalPtrRequest]
		Rating   form.OptIntField[OptionalPtrRequest]
		Priority form.OptIntField[OptionalPtrRequest]

		StartAt         form.OptTimeField[OptionalPtrRequest]
		AvailableFrom   form.OptTimeField[OptionalPtrRequest]
		Deadline        form.OptTimeField[OptionalPtrRequest]
		ExpiresAt       form.OptTimeField[OptionalPtrRequest]
		EventDate       form.OptTimeField[OptionalPtrRequest]
		ReservationDate form.OptTimeField[OptionalPtrRequest]

		Nickname     form.StringField[OptionalPtrRequest]
		Tags         form.SliceStringField[OptionalPtrRequest]
		ProfileBio   form.OptStringField[OptionalPtrRequest]
		ProfileTitle form.StringField[OptionalPtrRequest]
	}{
		Age: form.OptInt[OptionalPtrRequest]("age", func(r *OptionalPtrRequest) *int {
			return r.Age
		}),
		Score: form.OptInt[OptionalPtrRequest]("score", func(r *OptionalPtrRequest) *int {
			return r.Score
		}),
		Attempts: form.OptInt[OptionalPtrRequest]("attempts", func(r *OptionalPtrRequest) *int {
			return r.Attempts
		}),
		Limit: form.OptInt[OptionalPtrRequest]("limit", func(r *OptionalPtrRequest) *int {
			return r.Limit
		}),
		Rating: form.OptInt[OptionalPtrRequest]("rating", func(r *OptionalPtrRequest) *int {
			return r.Rating
		}),
		Priority: form.OptInt[OptionalPtrRequest]("priority", func(r *OptionalPtrRequest) *int {
			return r.Priority
		}),

		StartAt: form.OptTime[OptionalPtrRequest]("start_at", func(r *OptionalPtrRequest) *time.Time {
			return r.StartAt
		}),
		AvailableFrom: form.OptTime[OptionalPtrRequest]("available_from", func(r *OptionalPtrRequest) *time.Time {
			return r.AvailableFrom
		}),
		Deadline: form.OptTime[OptionalPtrRequest]("deadline", func(r *OptionalPtrRequest) *time.Time {
			return r.Deadline
		}),
		ExpiresAt: form.OptTime[OptionalPtrRequest]("expires_at", func(r *OptionalPtrRequest) *time.Time {
			return r.ExpiresAt
		}),
		EventDate: form.OptTime[OptionalPtrRequest]("event_date", func(r *OptionalPtrRequest) *time.Time {
			return r.EventDate
		}),
		ReservationDate: form.OptTime[OptionalPtrRequest]("reservation_date", func(r *OptionalPtrRequest) *time.Time {
			return r.ReservationDate
		}),

		Nickname: form.Str[OptionalPtrRequest]("nickname", func(r *OptionalPtrRequest) string {
			return r.Nickname
		}),
		Tags: form.SliceStr[OptionalPtrRequest]("tags", func(r *OptionalPtrRequest) []string {
			return r.Tags
		}),
		ProfileBio: form.OptStr[OptionalPtrRequest]("profile_bio", func(r *OptionalPtrRequest) *string {
			return r.ProfileBio
		}),
		ProfileTitle: form.Str[OptionalPtrRequest]("profile_title", func(r *OptionalPtrRequest) string {
			return r.ProfileTitle
		}),
	}

	startBoundary := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	deadlineBoundary := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

	windowStart := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	windowEnd := time.Date(2026, 6, 30, 23, 59, 59, 0, time.UTC)

	return form.Schema[OptionalPtrRequest]{
		// Optional int rules skip validation when the pointer is nil.
		optional.MinOpt(OptionalPtrForm.Age, 18),
		optional.MinOptWithCode(OptionalPtrForm.Score, 50, CodeMinAge),

		optional.MaxOpt(OptionalPtrForm.Attempts, 3),
		optional.MaxOptWithCode(OptionalPtrForm.Limit, 100, CodeMaxAttempts),

		optional.BetweenIntOpt(OptionalPtrForm.Rating, 1, 5),
		optional.BetweenIntOptWithCode(OptionalPtrForm.Priority, 1, 10, CodePriorityRange),

		// OptionalInt applies nested rules only when the pointer is non-nil.
		optional.OptionalInt(
			OptionalPtrForm.Age,
			optional.MinOpt(OptionalPtrForm.Age, 18),
		),

		// Optional time rules skip validation when the pointer is nil or zero.
		optional.AfterOpt(OptionalPtrForm.StartAt, startBoundary),
		optional.AfterOptWithCode(OptionalPtrForm.AvailableFrom, startBoundary, CodeAfterStart),

		optional.BeforeOpt(OptionalPtrForm.Deadline, deadlineBoundary),
		optional.BeforeOptWithCode(OptionalPtrForm.ExpiresAt, deadlineBoundary, CodeBeforeDeadline),

		optional.BetweenTimeOpt(OptionalPtrForm.EventDate, windowStart, windowEnd),
		optional.BetweenTimeOptWithCode(OptionalPtrForm.ReservationDate, windowStart, windowEnd, CodeBetweenWindow),

		// OptionalTime applies nested rules only when the pointer is non-nil.
		optional.OptionalTime(
			OptionalPtrForm.StartAt,
			optional.AfterOpt(OptionalPtrForm.StartAt, startBoundary),
		),

		// OptionalString applies nested rules only when the string is non-blank.
		optional.OptionalString(
			OptionalPtrForm.Nickname,
			rules.MinLen(OptionalPtrForm.Nickname, 3),
		),

		// OptionalSlice applies nested rules only when the slice has at least one item.
		optional.OptionalSlice[OptionalPtrRequest, string](
			OptionalPtrForm.Tags.Field,
			rules.MinItemsStr(OptionalPtrForm.Tags, 2),
			rules.DistinctStr(OptionalPtrForm.Tags),
		),

		// OptionalPtr applies nested rules only when profile_bio is non-nil.
		// In this example, when profile_bio is sent, profile_title becomes required.
		optional.OptionalPtr[OptionalPtrRequest, string](
			OptionalPtrForm.ProfileBio.Field,
			rules.RequiredWithCode(OptionalPtrForm.ProfileTitle, CodeProfileTitle),
		),
	}
}
