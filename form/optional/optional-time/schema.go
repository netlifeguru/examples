package main

import (
	"time"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/optional"
)

const (
	CodeAfterStart     = form.Code("after_start")
	CodeBeforeDeadline = form.Code("before_deadline")
	CodeBetweenWindow  = form.Code("between_window")
)

type OptionalTimeRequest struct {
	StartAt         *time.Time `json:"start_at"`
	AvailableFrom   *time.Time `json:"available_from"`
	Deadline        *time.Time `json:"deadline"`
	ExpiresAt       *time.Time `json:"expires_at"`
	EventDate       *time.Time `json:"event_date"`
	ReservationDate *time.Time `json:"reservation_date"`
}

func OptionalTimeSchema() form.Schema[OptionalTimeRequest] {
	OptionalTimeForm := struct {
		StartAt         form.OptTimeField[OptionalTimeRequest]
		AvailableFrom   form.OptTimeField[OptionalTimeRequest]
		Deadline        form.OptTimeField[OptionalTimeRequest]
		ExpiresAt       form.OptTimeField[OptionalTimeRequest]
		EventDate       form.OptTimeField[OptionalTimeRequest]
		ReservationDate form.OptTimeField[OptionalTimeRequest]
	}{
		StartAt: form.OptTime[OptionalTimeRequest]("start_at", func(r *OptionalTimeRequest) *time.Time {
			return r.StartAt
		}),
		AvailableFrom: form.OptTime[OptionalTimeRequest]("available_from", func(r *OptionalTimeRequest) *time.Time {
			return r.AvailableFrom
		}),
		Deadline: form.OptTime[OptionalTimeRequest]("deadline", func(r *OptionalTimeRequest) *time.Time {
			return r.Deadline
		}),
		ExpiresAt: form.OptTime[OptionalTimeRequest]("expires_at", func(r *OptionalTimeRequest) *time.Time {
			return r.ExpiresAt
		}),
		EventDate: form.OptTime[OptionalTimeRequest]("event_date", func(r *OptionalTimeRequest) *time.Time {
			return r.EventDate
		}),
		ReservationDate: form.OptTime[OptionalTimeRequest]("reservation_date", func(r *OptionalTimeRequest) *time.Time {
			return r.ReservationDate
		}),
	}

	startBoundary := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	deadlineBoundary := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

	windowStart := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	windowEnd := time.Date(2026, 6, 30, 23, 59, 59, 0, time.UTC)

	return form.Schema[OptionalTimeRequest]{
		// Optional time rules skip validation when the pointer is nil or zero time.

		optional.AfterOpt(OptionalTimeForm.StartAt, startBoundary),
		optional.AfterOptWithCode(OptionalTimeForm.AvailableFrom, startBoundary, CodeAfterStart),

		optional.BeforeOpt(OptionalTimeForm.Deadline, deadlineBoundary),
		optional.BeforeOptWithCode(OptionalTimeForm.ExpiresAt, deadlineBoundary, CodeBeforeDeadline),

		optional.BetweenTimeOpt(OptionalTimeForm.EventDate, windowStart, windowEnd),
		optional.BetweenTimeOptWithCode(OptionalTimeForm.ReservationDate, windowStart, windowEnd, CodeBetweenWindow),
	}
}
