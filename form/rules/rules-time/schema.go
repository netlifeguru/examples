package main

import (
	"time"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeAfterStart     = form.Code("after_start")
	CodeBeforeDeadline = form.Code("before_deadline")
	CodeBetweenWindow  = form.Code("between_window")
)

type TimeRulesRequest struct {
	StartAt         time.Time `json:"start_at"`
	AvailableFrom   time.Time `json:"available_from"`
	Deadline        time.Time `json:"deadline"`
	ExpiresAt       time.Time `json:"expires_at"`
	EventDate       time.Time `json:"event_date"`
	ReservationDate time.Time `json:"reservation_date"`
}

func TimeRulesSchema() form.Schema[TimeRulesRequest] {
	TimeForm := struct {
		StartAt         form.TimeField[TimeRulesRequest]
		AvailableFrom   form.TimeField[TimeRulesRequest]
		Deadline        form.TimeField[TimeRulesRequest]
		ExpiresAt       form.TimeField[TimeRulesRequest]
		EventDate       form.TimeField[TimeRulesRequest]
		ReservationDate form.TimeField[TimeRulesRequest]
	}{
		StartAt: form.Time[TimeRulesRequest]("start_at", func(r *TimeRulesRequest) time.Time {
			return r.StartAt
		}),
		AvailableFrom: form.Time[TimeRulesRequest]("available_from", func(r *TimeRulesRequest) time.Time {
			return r.AvailableFrom
		}),
		Deadline: form.Time[TimeRulesRequest]("deadline", func(r *TimeRulesRequest) time.Time {
			return r.Deadline
		}),
		ExpiresAt: form.Time[TimeRulesRequest]("expires_at", func(r *TimeRulesRequest) time.Time {
			return r.ExpiresAt
		}),
		EventDate: form.Time[TimeRulesRequest]("event_date", func(r *TimeRulesRequest) time.Time {
			return r.EventDate
		}),
		ReservationDate: form.Time[TimeRulesRequest]("reservation_date", func(r *TimeRulesRequest) time.Time {
			return r.ReservationDate
		}),
	}

	startBoundary := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	deadlineBoundary := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

	windowStart := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	windowEnd := time.Date(2026, 6, 30, 23, 59, 59, 0, time.UTC)

	return form.Schema[TimeRulesRequest]{
		// After validates that the time is after the given boundary.
		rules.After(TimeForm.StartAt, startBoundary),

		// AfterWithCode validates that the time is after the given boundary and returns a custom code.
		rules.AfterWithCode(TimeForm.AvailableFrom, startBoundary, CodeAfterStart),

		// Before validates that the time is before the given boundary.
		rules.Before(TimeForm.Deadline, deadlineBoundary),

		// BeforeWithCode validates that the time is before the given boundary and returns a custom code.
		rules.BeforeWithCode(TimeForm.ExpiresAt, deadlineBoundary, CodeBeforeDeadline),

		// BetweenTime validates that the time is inside the min/max range, inclusive.
		rules.BetweenTime(TimeForm.EventDate, windowStart, windowEnd),

		// BetweenTimeWithCode validates that the time is inside the min/max range and returns a custom code.
		rules.BetweenTimeWithCode(TimeForm.ReservationDate, windowStart, windowEnd, CodeBetweenWindow),
	}
}
