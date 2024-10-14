//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

const (
	BasicAuthScopes  = "basicAuth.Scopes"
	CookieAuthScopes = "cookieAuth.Scopes"
)

// Defines values for BotStatus.
const (
	IDLE   BotStatus = "IDLE"
	MINING BotStatus = "MINING"
	MOVING BotStatus = "MOVING"
)

// Bot defines model for Bot.
type Bot struct {
	Coordinates Coordinates `json:"coordinates"`
	Identifier  string      `json:"identifier"`
	Inventory   int         `json:"inventory"`
	Name        string      `json:"name"`
	Status      BotStatus   `json:"status"`
}

// BotStatus defines model for Bot.Status.
type BotStatus string

// Coordinates defines model for Coordinates.
type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// PostBotsBotIdNewBotJSONBody defines parameters for PostBotsBotIdNewBot.
type PostBotsBotIdNewBotJSONBody struct {
	NewBotName string `json:"NewBotName"`
}

// PostBotsBotIdMoveJSONRequestBody defines body for PostBotsBotIdMove for application/json ContentType.
type PostBotsBotIdMoveJSONRequestBody = Coordinates

// PostBotsBotIdNewBotJSONRequestBody defines body for PostBotsBotIdNewBot for application/json ContentType.
type PostBotsBotIdNewBotJSONRequestBody PostBotsBotIdNewBotJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /bots)
	GetBots(w http.ResponseWriter, r *http.Request)

	// (GET /bots/{botId})
	GetBotsBotId(w http.ResponseWriter, r *http.Request, botId string)

	// (POST /bots/{botId}/mine)
	PostBotsBotIdMine(w http.ResponseWriter, r *http.Request, botId string)

	// (POST /bots/{botId}/move)
	PostBotsBotIdMove(w http.ResponseWriter, r *http.Request, botId string)

	// (POST /bots/{botId}/newBot)
	PostBotsBotIdNewBot(w http.ResponseWriter, r *http.Request, botId string)

	// (POST /init)
	PostInit(w http.ResponseWriter, r *http.Request)

	// (POST /login)
	PostLogin(w http.ResponseWriter, r *http.Request)

	// (GET /mines)
	GetMines(w http.ResponseWriter, r *http.Request)

	// (POST /newUser)
	PostNewUser(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetBots operation middleware
func (siw *ServerInterfaceWrapper) GetBots(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, CookieAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBots(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetBotsBotId operation middleware
func (siw *ServerInterfaceWrapper) GetBotsBotId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "botId" -------------
	var botId string

	err = runtime.BindStyledParameterWithOptions("simple", "botId", r.PathValue("botId"), &botId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "botId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, CookieAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBotsBotId(w, r, botId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostBotsBotIdMine operation middleware
func (siw *ServerInterfaceWrapper) PostBotsBotIdMine(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "botId" -------------
	var botId string

	err = runtime.BindStyledParameterWithOptions("simple", "botId", r.PathValue("botId"), &botId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "botId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, CookieAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostBotsBotIdMine(w, r, botId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostBotsBotIdMove operation middleware
func (siw *ServerInterfaceWrapper) PostBotsBotIdMove(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "botId" -------------
	var botId string

	err = runtime.BindStyledParameterWithOptions("simple", "botId", r.PathValue("botId"), &botId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "botId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, CookieAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostBotsBotIdMove(w, r, botId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostBotsBotIdNewBot operation middleware
func (siw *ServerInterfaceWrapper) PostBotsBotIdNewBot(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "botId" -------------
	var botId string

	err = runtime.BindStyledParameterWithOptions("simple", "botId", r.PathValue("botId"), &botId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "botId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, CookieAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostBotsBotIdNewBot(w, r, botId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostInit operation middleware
func (siw *ServerInterfaceWrapper) PostInit(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, CookieAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostInit(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostLogin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMines operation middleware
func (siw *ServerInterfaceWrapper) GetMines(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, CookieAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMines(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostNewUser operation middleware
func (siw *ServerInterfaceWrapper) PostNewUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNewUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/bots", wrapper.GetBots)
	m.HandleFunc("GET "+options.BaseURL+"/bots/{botId}", wrapper.GetBotsBotId)
	m.HandleFunc("POST "+options.BaseURL+"/bots/{botId}/mine", wrapper.PostBotsBotIdMine)
	m.HandleFunc("POST "+options.BaseURL+"/bots/{botId}/move", wrapper.PostBotsBotIdMove)
	m.HandleFunc("POST "+options.BaseURL+"/bots/{botId}/newBot", wrapper.PostBotsBotIdNewBot)
	m.HandleFunc("POST "+options.BaseURL+"/init", wrapper.PostInit)
	m.HandleFunc("POST "+options.BaseURL+"/login", wrapper.PostLogin)
	m.HandleFunc("GET "+options.BaseURL+"/mines", wrapper.GetMines)
	m.HandleFunc("POST "+options.BaseURL+"/newUser", wrapper.PostNewUser)

	return m
}

type GetBotsRequestObject struct {
}

type GetBotsResponseObject interface {
	VisitGetBotsResponse(w http.ResponseWriter) error
}

type GetBots200JSONResponse []Bot

func (response GetBots200JSONResponse) VisitGetBotsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetBotsBotIdRequestObject struct {
	BotId string `json:"botId"`
}

type GetBotsBotIdResponseObject interface {
	VisitGetBotsBotIdResponse(w http.ResponseWriter) error
}

type GetBotsBotId200JSONResponse Bot

func (response GetBotsBotId200JSONResponse) VisitGetBotsBotIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostBotsBotIdMineRequestObject struct {
	BotId string `json:"botId"`
}

type PostBotsBotIdMineResponseObject interface {
	VisitPostBotsBotIdMineResponse(w http.ResponseWriter) error
}

type PostBotsBotIdMine200JSONResponse Bot

func (response PostBotsBotIdMine200JSONResponse) VisitPostBotsBotIdMineResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostBotsBotIdMine422TextResponse string

func (response PostBotsBotIdMine422TextResponse) VisitPostBotsBotIdMineResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(422)

	_, err := w.Write([]byte(response))
	return err
}

type PostBotsBotIdMoveRequestObject struct {
	BotId string `json:"botId"`
	Body  *PostBotsBotIdMoveJSONRequestBody
}

type PostBotsBotIdMoveResponseObject interface {
	VisitPostBotsBotIdMoveResponse(w http.ResponseWriter) error
}

type PostBotsBotIdMove200JSONResponse Bot

func (response PostBotsBotIdMove200JSONResponse) VisitPostBotsBotIdMoveResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostBotsBotIdNewBotRequestObject struct {
	BotId string `json:"botId"`
	Body  *PostBotsBotIdNewBotJSONRequestBody
}

type PostBotsBotIdNewBotResponseObject interface {
	VisitPostBotsBotIdNewBotResponse(w http.ResponseWriter) error
}

type PostBotsBotIdNewBot200JSONResponse Bot

func (response PostBotsBotIdNewBot200JSONResponse) VisitPostBotsBotIdNewBotResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostBotsBotIdNewBot422TextResponse string

func (response PostBotsBotIdNewBot422TextResponse) VisitPostBotsBotIdNewBotResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(422)

	_, err := w.Write([]byte(response))
	return err
}

type PostInitRequestObject struct {
}

type PostInitResponseObject interface {
	VisitPostInitResponse(w http.ResponseWriter) error
}

type PostInit200JSONResponse struct {
	Bots  *[]Bot         `json:"bots,omitempty"`
	Mines *[]Coordinates `json:"mines,omitempty"`
}

func (response PostInit200JSONResponse) VisitPostInitResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostLoginRequestObject struct {
}

type PostLoginResponseObject interface {
	VisitPostLoginResponse(w http.ResponseWriter) error
}

type PostLogin200TextResponse string

func (response PostLogin200TextResponse) VisitPostLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	_, err := w.Write([]byte(response))
	return err
}

type PostLogin401TextResponse string

func (response PostLogin401TextResponse) VisitPostLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(401)

	_, err := w.Write([]byte(response))
	return err
}

type GetMinesRequestObject struct {
}

type GetMinesResponseObject interface {
	VisitGetMinesResponse(w http.ResponseWriter) error
}

type GetMines200JSONResponse []Coordinates

func (response GetMines200JSONResponse) VisitGetMinesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostNewUserRequestObject struct {
}

type PostNewUserResponseObject interface {
	VisitPostNewUserResponse(w http.ResponseWriter) error
}

type PostNewUser200TextResponse string

func (response PostNewUser200TextResponse) VisitPostNewUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	_, err := w.Write([]byte(response))
	return err
}

type PostNewUser401TextResponse string

func (response PostNewUser401TextResponse) VisitPostNewUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(401)

	_, err := w.Write([]byte(response))
	return err
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /bots)
	GetBots(ctx context.Context, request GetBotsRequestObject) (GetBotsResponseObject, error)

	// (GET /bots/{botId})
	GetBotsBotId(ctx context.Context, request GetBotsBotIdRequestObject) (GetBotsBotIdResponseObject, error)

	// (POST /bots/{botId}/mine)
	PostBotsBotIdMine(ctx context.Context, request PostBotsBotIdMineRequestObject) (PostBotsBotIdMineResponseObject, error)

	// (POST /bots/{botId}/move)
	PostBotsBotIdMove(ctx context.Context, request PostBotsBotIdMoveRequestObject) (PostBotsBotIdMoveResponseObject, error)

	// (POST /bots/{botId}/newBot)
	PostBotsBotIdNewBot(ctx context.Context, request PostBotsBotIdNewBotRequestObject) (PostBotsBotIdNewBotResponseObject, error)

	// (POST /init)
	PostInit(ctx context.Context, request PostInitRequestObject) (PostInitResponseObject, error)

	// (POST /login)
	PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error)

	// (GET /mines)
	GetMines(ctx context.Context, request GetMinesRequestObject) (GetMinesResponseObject, error)

	// (POST /newUser)
	PostNewUser(ctx context.Context, request PostNewUserRequestObject) (PostNewUserResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetBots operation middleware
func (sh *strictHandler) GetBots(w http.ResponseWriter, r *http.Request) {
	var request GetBotsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetBots(ctx, request.(GetBotsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetBots")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetBotsResponseObject); ok {
		if err := validResponse.VisitGetBotsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetBotsBotId operation middleware
func (sh *strictHandler) GetBotsBotId(w http.ResponseWriter, r *http.Request, botId string) {
	var request GetBotsBotIdRequestObject

	request.BotId = botId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetBotsBotId(ctx, request.(GetBotsBotIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetBotsBotId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetBotsBotIdResponseObject); ok {
		if err := validResponse.VisitGetBotsBotIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostBotsBotIdMine operation middleware
func (sh *strictHandler) PostBotsBotIdMine(w http.ResponseWriter, r *http.Request, botId string) {
	var request PostBotsBotIdMineRequestObject

	request.BotId = botId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostBotsBotIdMine(ctx, request.(PostBotsBotIdMineRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostBotsBotIdMine")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostBotsBotIdMineResponseObject); ok {
		if err := validResponse.VisitPostBotsBotIdMineResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostBotsBotIdMove operation middleware
func (sh *strictHandler) PostBotsBotIdMove(w http.ResponseWriter, r *http.Request, botId string) {
	var request PostBotsBotIdMoveRequestObject

	request.BotId = botId

	var body PostBotsBotIdMoveJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostBotsBotIdMove(ctx, request.(PostBotsBotIdMoveRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostBotsBotIdMove")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostBotsBotIdMoveResponseObject); ok {
		if err := validResponse.VisitPostBotsBotIdMoveResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostBotsBotIdNewBot operation middleware
func (sh *strictHandler) PostBotsBotIdNewBot(w http.ResponseWriter, r *http.Request, botId string) {
	var request PostBotsBotIdNewBotRequestObject

	request.BotId = botId

	var body PostBotsBotIdNewBotJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostBotsBotIdNewBot(ctx, request.(PostBotsBotIdNewBotRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostBotsBotIdNewBot")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostBotsBotIdNewBotResponseObject); ok {
		if err := validResponse.VisitPostBotsBotIdNewBotResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostInit operation middleware
func (sh *strictHandler) PostInit(w http.ResponseWriter, r *http.Request) {
	var request PostInitRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostInit(ctx, request.(PostInitRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostInit")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostInitResponseObject); ok {
		if err := validResponse.VisitPostInitResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostLogin operation middleware
func (sh *strictHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	var request PostLoginRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostLogin(ctx, request.(PostLoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostLogin")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostLoginResponseObject); ok {
		if err := validResponse.VisitPostLoginResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetMines operation middleware
func (sh *strictHandler) GetMines(w http.ResponseWriter, r *http.Request) {
	var request GetMinesRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetMines(ctx, request.(GetMinesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMines")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetMinesResponseObject); ok {
		if err := validResponse.VisitGetMinesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostNewUser operation middleware
func (sh *strictHandler) PostNewUser(w http.ResponseWriter, r *http.Request) {
	var request PostNewUserRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostNewUser(ctx, request.(PostNewUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostNewUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostNewUserResponseObject); ok {
		if err := validResponse.VisitPostNewUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
