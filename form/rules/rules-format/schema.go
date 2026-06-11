package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeWebsiteURL  = form.Code("website_url")
	CodeServerIP    = form.Code("server_ip")
	CodeResourceID  = form.Code("resource_id")
	CodePayloadJSON = form.Code("payload_json")
	CodeUserZone    = form.Code("user_timezone")
)

type FormatRulesRequest struct {
	Website           string `json:"website"`
	CallbackURL       string `json:"callback_url"`
	ServerIP          string `json:"server_ip"`
	BackupIP          string `json:"backup_ip"`
	ResourceID        string `json:"resource_id"`
	SessionID         string `json:"session_id"`
	Payload           string `json:"payload"`
	Metadata          string `json:"metadata"`
	Timezone          string `json:"timezone"`
	PreferredTimezone string `json:"preferred_timezone"`
}

func FormatRulesSchema() form.Schema[FormatRulesRequest] {
	FormatForm := struct {
		Website           form.StringField[FormatRulesRequest]
		CallbackURL       form.StringField[FormatRulesRequest]
		ServerIP          form.StringField[FormatRulesRequest]
		BackupIP          form.StringField[FormatRulesRequest]
		ResourceID        form.StringField[FormatRulesRequest]
		SessionID         form.StringField[FormatRulesRequest]
		Payload           form.StringField[FormatRulesRequest]
		Metadata          form.StringField[FormatRulesRequest]
		Timezone          form.StringField[FormatRulesRequest]
		PreferredTimezone form.StringField[FormatRulesRequest]
	}{
		Website: form.Str[FormatRulesRequest]("website", func(r *FormatRulesRequest) string {
			return r.Website
		}),
		CallbackURL: form.Str[FormatRulesRequest]("callback_url", func(r *FormatRulesRequest) string {
			return r.CallbackURL
		}),
		ServerIP: form.Str[FormatRulesRequest]("server_ip", func(r *FormatRulesRequest) string {
			return r.ServerIP
		}),
		BackupIP: form.Str[FormatRulesRequest]("backup_ip", func(r *FormatRulesRequest) string {
			return r.BackupIP
		}),
		ResourceID: form.Str[FormatRulesRequest]("resource_id", func(r *FormatRulesRequest) string {
			return r.ResourceID
		}),
		SessionID: form.Str[FormatRulesRequest]("session_id", func(r *FormatRulesRequest) string {
			return r.SessionID
		}),
		Payload: form.Str[FormatRulesRequest]("payload", func(r *FormatRulesRequest) string {
			return r.Payload
		}),
		Metadata: form.Str[FormatRulesRequest]("metadata", func(r *FormatRulesRequest) string {
			return r.Metadata
		}),
		Timezone: form.Str[FormatRulesRequest]("timezone", func(r *FormatRulesRequest) string {
			return r.Timezone
		}),
		PreferredTimezone: form.Str[FormatRulesRequest]("preferred_timezone", func(r *FormatRulesRequest) string {
			return r.PreferredTimezone
		}),
	}

	return form.Schema[FormatRulesRequest]{
		rules.URL(FormatForm.Website),
		rules.URLWithCode(FormatForm.CallbackURL, CodeWebsiteURL),

		rules.IP(FormatForm.ServerIP),
		rules.IPWithCode(FormatForm.BackupIP, CodeServerIP),

		rules.UUID(FormatForm.ResourceID),
		rules.UUIDWithCode(FormatForm.SessionID, CodeResourceID),

		rules.JSON(FormatForm.Payload),
		rules.JSONWithCode(FormatForm.Metadata, CodePayloadJSON),

		rules.Timezone(FormatForm.Timezone),
		rules.TimezoneWithCode(FormatForm.PreferredTimezone, CodeUserZone),
	}
}
