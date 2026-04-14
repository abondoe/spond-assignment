package service

import "errors"

var (
	ErrInvalidMemberTypeID   = errors.New("invalid member type id")
	ErrFormNotFound          = errors.New("form not found")
	ErrInvalidRegistration   = errors.New("invalid registration")
	ErrDuplicateRegistration = errors.New("duplicate registration")
	ErrDatabaseError         = errors.New("database error")
)
