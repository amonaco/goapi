// api v0.0.1 268831ec79ad41c6f6fff6d70ae0833650569b51
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/golang
// Do not edit by hand. Update your webrpc schema and re-generate.
package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.0.1"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "268831ec79ad41c6f6fff6d70ae0833650569b51"
}

//
// Types
//

type Status struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	ID      uint32 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type User struct {
	ID        *uint32  `json:"id"`
	CompanyID uint32   `json:"company_id"`
	UUID      *string  `json:"uuid"`
	SectorID  *uint32  `json:"sector_id"`
	Email     string   `json:"email"`
	Address   string   `json:"address"`
	Phone     string   `json:"phone"`
	Name      string   `json:"name"`
	Lastname  string   `json:"lastname"`
	Roles     []string `json:"roles"`
	Status    *string  `json:"status"`
}

type UserService interface {
	Get(ctx context.Context, id uint32) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	CreateEditorUser(ctx context.Context, user *User) (*Status, error)
	Update(ctx context.Context, user *User) (*Status, error)
	GetCurrent(ctx context.Context) (*User, error)
	GetByUUID(ctx context.Context, uuid string) (*User, error)
}

var WebRPCServices = map[string][]string{
	"UserService": {
		"Get",
		"GetAll",
		"CreateEditorUser",
		"Update",
		"GetCurrent",
		"GetByUUID",
	},
}

//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type userServiceServer struct {
	UserService
}

func NewUserServiceServer(svc UserService) WebRPCServer {
	return &userServiceServer{
		UserService: svc,
	}
}

func (s *userServiceServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "UserService")

	if r.Method != "POST" {
		err := Errorf(ErrBadRoute, "unsupported method %q (only POST is allowed)", r.Method)
		RespondWithError(w, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/UserService/Get":
		s.serveGet(ctx, w, r)
		return
	case "/rpc/UserService/GetAll":
		s.serveGetAll(ctx, w, r)
		return
	case "/rpc/UserService/CreateEditorUser":
		s.serveCreateEditorUser(ctx, w, r)
		return
	case "/rpc/UserService/Update":
		s.serveUpdate(ctx, w, r)
		return
	case "/rpc/UserService/GetCurrent":
		s.serveGetCurrent(ctx, w, r)
		return
	case "/rpc/UserService/GetByUUID":
		s.serveGetByUUID(ctx, w, r)
		return
	default:
		err := Errorf(ErrBadRoute, "no handler for path %q", r.URL.Path)
		RespondWithError(w, err)
		return
	}
}

func (s *userServiceServer) serveGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveGetJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Get")
	reqContent := struct {
		Arg0 uint32 `json:"id"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to read request data")
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = WrapError(ErrInvalidArgument, err, "failed to unmarshal request data")
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.Get(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *User `json:"result"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *userServiceServer) serveGetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetAllJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveGetAllJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetAll")

	// Call service method
	var ret0 []*User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.GetAll(ctx)
	}()
	respContent := struct {
		Ret0 []*User `json:"result"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *userServiceServer) serveCreateEditorUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateEditorUserJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveCreateEditorUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "CreateEditorUser")
	reqContent := struct {
		Arg0 *User `json:"user"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to read request data")
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = WrapError(ErrInvalidArgument, err, "failed to unmarshal request data")
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *Status
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.CreateEditorUser(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *Status `json:"status"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *userServiceServer) serveUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUpdateJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveUpdateJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Update")
	reqContent := struct {
		Arg0 *User `json:"user"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to read request data")
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = WrapError(ErrInvalidArgument, err, "failed to unmarshal request data")
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *Status
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.Update(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *Status `json:"status"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *userServiceServer) serveGetCurrent(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetCurrentJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveGetCurrentJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetCurrent")

	// Call service method
	var ret0 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.GetCurrent(ctx)
	}()
	respContent := struct {
		Ret0 *User `json:"result"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *userServiceServer) serveGetByUUID(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetByUUIDJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveGetByUUIDJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetByUUID")
	reqContent := struct {
		Arg0 string `json:"uuid"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to read request data")
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = WrapError(ErrInvalidArgument, err, "failed to unmarshal request data")
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.GetByUUID(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *User `json:"result"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(Error)
	if !ok {
		rpcErr = WrapError(ErrInternal, err, "webrpc error")
	}

	statusCode := HTTPStatusFromErrorCode(rpcErr.Code())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	respBody, _ := json.Marshal(rpcErr.Payload())
	w.Write(respBody)
}

//
// Helpers
//

type ErrorPayload struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Cause  string `json:"cause,omitempty"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type Error interface {
	// Code is of the valid error codes
	Code() ErrorCode

	// Msg returns a human-readable, unstructured messages describing the error
	Msg() string

	// Cause is reason for the error
	Cause() error

	// Error returns a string of the form "webrpc error <Code>: <Msg>"
	Error() string

	// Error response payload
	Payload() ErrorPayload
}

func Errorf(code ErrorCode, msgf string, args ...interface{}) Error {
	msg := fmt.Sprintf(msgf, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code)}
}

func WrapError(code ErrorCode, cause error, format string, args ...interface{}) Error {
	msg := fmt.Sprintf(format, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg, cause: cause}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code), cause: cause}
}

func Failf(format string, args ...interface{}) Error {
	return Errorf(ErrFail, format, args...)
}

func WrapFailf(cause error, format string, args ...interface{}) Error {
	return WrapError(ErrFail, cause, format, args...)
}

func ErrorNotFound(format string, args ...interface{}) Error {
	return Errorf(ErrNotFound, format, args...)
}

func ErrorInvalidArgument(argument string, validationMsg string) Error {
	return Errorf(ErrInvalidArgument, argument+" "+validationMsg)
}

func ErrorRequiredArgument(argument string) Error {
	return ErrorInvalidArgument(argument, "is required")
}

func ErrorInternal(format string, args ...interface{}) Error {
	return Errorf(ErrInternal, format, args...)
}

type ErrorCode string

const (
	// Unknown error. For example when handling errors raised by APIs that do not
	// return enough error information.
	ErrUnknown ErrorCode = "unknown"

	// Fail error. General failure error type.
	ErrFail ErrorCode = "fail"

	// Canceled indicates the operation was cancelled (typically by the caller).
	ErrCanceled ErrorCode = "canceled"

	// InvalidArgument indicates client specified an invalid argument. It
	// indicates arguments that are problematic regardless of the state of the
	// system (i.e. a malformed file name, required argument, number out of range,
	// etc.).
	ErrInvalidArgument ErrorCode = "invalid argument"

	// DeadlineExceeded means operation expired before completion. For operations
	// that change the state of the system, this error may be returned even if the
	// operation has completed successfully (timeout).
	ErrDeadlineExceeded ErrorCode = "deadline exceeded"

	// NotFound means some requested entity was not found.
	ErrNotFound ErrorCode = "not found"

	// BadRoute means that the requested URL path wasn't routable to a webrpc
	// service and method. This is returned by the generated server, and usually
	// shouldn't be returned by applications. Instead, applications should use
	// NotFound or Unimplemented.
	ErrBadRoute ErrorCode = "bad route"

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	ErrAlreadyExists ErrorCode = "already exists"

	// PermissionDenied indicates the caller does not have permission to execute
	// the specified operation. It must not be used if the caller cannot be
	// identified (Unauthenticated).
	ErrPermissionDenied ErrorCode = "permission denied"

	// Unauthenticated indicates the request does not have valid authentication
	// credentials for the operation.
	ErrUnauthenticated ErrorCode = "unauthenticated"

	// ResourceExhausted indicates some resource has been exhausted, perhaps a
	// per-user quota, or perhaps the entire file system is out of space.
	ErrResourceExhausted ErrorCode = "resource exhausted"

	// FailedPrecondition indicates operation was rejected because the system is
	// not in a state required for the operation's execution. For example, doing
	// an rmdir operation on a directory that is non-empty, or on a non-directory
	// object, or when having conflicting read-modify-write on the same resource.
	ErrFailedPrecondition ErrorCode = "failed precondition"

	// Aborted indicates the operation was aborted, typically due to a concurrency
	// issue like sequencer check failures, transaction aborts, etc.
	ErrAborted ErrorCode = "aborted"

	// OutOfRange means operation was attempted past the valid range. For example,
	// seeking or reading past end of a paginated collection.
	//
	// Unlike InvalidArgument, this error indicates a problem that may be fixed if
	// the system state changes (i.e. adding more items to the collection).
	//
	// There is a fair bit of overlap between FailedPrecondition and OutOfRange.
	// We recommend using OutOfRange (the more specific error) when it applies so
	// that callers who are iterating through a space can easily look for an
	// OutOfRange error to detect when they are done.
	ErrOutOfRange ErrorCode = "out of range"

	// Unimplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	ErrUnimplemented ErrorCode = "unimplemented"

	// Internal errors. When some invariants expected by the underlying system
	// have been broken. In other words, something bad happened in the library or
	// backend service. Do not confuse with HTTP Internal Server Error; an
	// Internal error could also happen on the client code, i.e. when parsing a
	// server response.
	ErrInternal ErrorCode = "internal"

	// Unavailable indicates the service is currently unavailable. This is a most
	// likely a transient condition and may be corrected by retrying with a
	// backoff.
	ErrUnavailable ErrorCode = "unavailable"

	// DataLoss indicates unrecoverable data loss or corruption.
	ErrDataLoss ErrorCode = "data loss"

	// ErrNone is the zero-value, is considered an empty error and should not be
	// used.
	ErrNone ErrorCode = ""
)

func HTTPStatusFromErrorCode(code ErrorCode) int {
	switch code {
	case ErrCanceled:
		return 408 // RequestTimeout
	case ErrUnknown:
		return 400 // Bad Request
	case ErrFail:
		return 422 // Unprocessable Entity
	case ErrInvalidArgument:
		return 400 // BadRequest
	case ErrDeadlineExceeded:
		return 408 // RequestTimeout
	case ErrNotFound:
		return 404 // Not Found
	case ErrBadRoute:
		return 404 // Not Found
	case ErrAlreadyExists:
		return 409 // Conflict
	case ErrPermissionDenied:
		return 403 // Forbidden
	case ErrUnauthenticated:
		return 401 // Unauthorized
	case ErrResourceExhausted:
		return 403 // Forbidden
	case ErrFailedPrecondition:
		return 412 // Precondition Failed
	case ErrAborted:
		return 409 // Conflict
	case ErrOutOfRange:
		return 400 // Bad Request
	case ErrUnimplemented:
		return 501 // Not Implemented
	case ErrInternal:
		return 500 // Internal Server Error
	case ErrUnavailable:
		return 503 // Service Unavailable
	case ErrDataLoss:
		return 500 // Internal Server Error
	case ErrNone:
		return 200 // OK
	default:
		return 0 // Invalid!
	}
}

func IsErrorCode(err error, code ErrorCode) bool {
	if rpcErr, ok := err.(Error); ok {
		if rpcErr.Code() == code {
			return true
		}
	}
	return false
}

func IsValidErrorCode(code ErrorCode) bool {
	return HTTPStatusFromErrorCode(code) != 0
}

type rpcErr struct {
	code  ErrorCode
	msg   string
	cause error
}

func (e *rpcErr) Code() ErrorCode {
	return e.code
}

func (e *rpcErr) Msg() string {
	return e.msg
}

func (e *rpcErr) Cause() error {
	return e.cause
}

func (e *rpcErr) Error() string {
	if e.cause != nil && e.cause.Error() != "" {
		if e.msg != "" {
			return fmt.Sprintf("webrpc %s error: %s -- %s", e.code, e.cause.Error(), e.msg)
		} else {
			return fmt.Sprintf("webrpc %s error: %s", e.code, e.cause.Error())
		}
	} else {
		return fmt.Sprintf("webrpc %s error: %s", e.code, e.msg)
	}
}

func (e *rpcErr) Payload() ErrorPayload {
	statusCode := HTTPStatusFromErrorCode(e.Code())
	errPayload := ErrorPayload{
		Status: statusCode,
		Code:   string(e.Code()),
		Msg:    e.Msg(),
		Error:  e.Error(),
	}
	if e.Cause() != nil {
		errPayload.Cause = e.Cause().Error()
	}
	return errPayload
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	// For Client
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}

	// For Server
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)
