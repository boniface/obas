package misc

import (
	"net/http"
	"obas/config"
	"strings"
	"time"
)

const (
	YYYYMMDD_FORMAT                    = "2006-01-02"
	YYYMMDDTIME_FORMAT                 = "2006-01-02 15:04:05"
	DangerAlertStyle                   = "alert-danger"
	SuccessAlertStyle                  = "alert-success"
	NonMatricProgressBarIncrementValue = 4.1666666667
	MatricProgressBarIncrementVale     = 5.882353
)

type PageToast struct {
	AlertType string
	AlertInfo string
}

/**
Check if session has alert message
*/
func CheckForSessionAlert(app *config.Env, r *http.Request) PageToast {
	message := app.Session.GetString(r.Context(), "message")
	messageType := app.Session.GetString(r.Context(), "message-type")
	var alert PageToast
	if message != "" && messageType != "" {
		alert = PageToast{messageType, message}
		app.Session.Remove(r.Context(), "message")
		app.Session.Remove(r.Context(), "message-type")
	}
	return alert
}

/**
Set session alert message
*/
func SetSessionMessage(app *config.Env, r *http.Request, messageType string, message string) {
	app.Session.Put(r.Context(), "message-type", messageType)
	app.Session.Put(r.Context(), "message", message)
}

/**
Get status badge
*/
func GetBadge(status string) string {
	status = strings.ToLower(strings.TrimSpace(status))
	statusBadge := "badge-info"
	if status == "approved" {
		statusBadge = "badge-success"
	} else if status == "rejected" {
		statusBadge = "badge-danger"
	} else if status == "unknown" {
		statusBadge = "badge-warning"
	}
	return statusBadge
}

/**
Format date in yyyy-MM-dd HH:mm:ss
 */
func FormatDateTime(date time.Time) string {
	return date.Format(YYYMMDDTIME_FORMAT)
}

/**
format date in yyyy-MM-dd
*/
func FormatDate(date time.Time) string {
	return date.Format(YYYYMMDD_FORMAT)
}
