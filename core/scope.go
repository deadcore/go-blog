package core

import (
	"net/http"
	"time"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	// UserID returns the ID of the user for the current request
	UserID() string
	// SetUserID sets the ID of the currently authenticated user
	SetUserID(id string)
	// RequestID returns the ID of the current request
	RequestID() string
	// Rollback returns a value indicating whether the current database transaction should be rolled back
	Rollback() bool
	// SetRollback sets a value indicating whether the current database transaction should be rolled back
	SetRollback(bool)
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
}

type requestScope struct {
	now       time.Time // the time when the request is being processed
	requestID string    // an ID identifying one or multiple correlated HTTP requests
	userID    string    // an ID identifying the current user
	rollback  bool      // whether to roll back the current transaction
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) SetUserID(id string) {
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Rollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

// newRequestScope creates a new RequestScope with the current request information.
func newRequestScope(now time.Time, request *http.Request) RequestScope {
	requestID := request.Header.Get("X-Request-Id")
	return &requestScope{
		now:       now,
		requestID: requestID,
	}
}