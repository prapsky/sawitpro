// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Errors ErrorResponseData `json:"errors"`
}

// ErrorResponseData defines model for ErrorResponseData.
type ErrorResponseData = []struct {
	Message string `json:"message"`
}

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Data *LoginResponseData `json:"data,omitempty"`
}

// LoginResponseData defines model for LoginResponseData.
type LoginResponseData struct {
	Token  string `json:"token"`
	UserID uint64 `json:"userID"`
}

// ProfileResponse defines model for ProfileResponse.
type ProfileResponse struct {
	Data *ProfileResponseData `json:"data,omitempty"`
}

// ProfileResponseData defines model for ProfileResponseData.
type ProfileResponseData struct {
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

// RegisterRequest defines model for RegisterRequest.
type RegisterRequest struct {
	FullName    string `json:"fullName"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

// RegisterResponse defines model for RegisterResponse.
type RegisterResponse struct {
	Data *RegisterResponseData `json:"data,omitempty"`
}

// RegisterResponseData defines model for RegisterResponseData.
type RegisterResponseData struct {
	UserID uint64 `json:"userID"`
}

// UpdateProfileRequest defines model for UpdateProfileRequest.
type UpdateProfileRequest struct {
	FullName    *string `json:"fullName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// UpdateProfileResponse defines model for UpdateProfileResponse.
type UpdateProfileResponse struct {
	Message *string `json:"message,omitempty"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LoginRequest

// UpdateProfileJSONRequestBody defines body for UpdateProfile for application/json ContentType.
type UpdateProfileJSONRequestBody = UpdateProfileRequest

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody = RegisterRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// This endpoint is to login to an account.
	// (POST /login)
	Login(ctx echo.Context) error
	// This endpoint is to view user profile.
	// (GET /profile)
	Profile(ctx echo.Context) error
	// This endpoint is to update user profile.
	// (PATCH /profile)
	UpdateProfile(ctx echo.Context) error
	// This endpoint is to register a new account.
	// (POST /register)
	Register(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// Profile converts echo context to params.
func (w *ServerInterfaceWrapper) Profile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Profile(ctx)
	return err
}

// UpdateProfile converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateProfile(ctx)
	return err
}

// Register converts echo context to params.
func (w *ServerInterfaceWrapper) Register(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Register(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/login", wrapper.Login)
	router.GET(baseURL+"/profile", wrapper.Profile)
	router.PATCH(baseURL+"/profile", wrapper.UpdateProfile)
	router.POST(baseURL+"/register", wrapper.Register)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RWTY/bNhD9K8K0R2HteDfbQLcEadEN2jTYD/Sw8IGmxhZTiWSG5G4MQ/+9IPW1suh1",
	"0awPbW6yNZx5fO/NaHbAVaWVRGkNZDswvMCKhcefiRRdo9FKGvR/aFIayQoMr9G/Dk8/Eq4hgx9mQ6pZ",
	"m2c2SvKeWQZ1nQLhFycIc8juuzzLFOxWI2SgVp+RW6hTmB7OdiAsVmYKp0Jj2CbgbPMYS0JuJvW6wGXd",
	"V2REbOsL/qY2Ql7jF4fGTktoZsyjojxSIwVdKIkfXbVCOo7haXA65I1x0EI6pELesvKcBqMUvQbPV+rY",
	"Hlez6i+UQfyvrNKlP/9qcc5WHNIpI84gXb0fRS/mKawVVcxCBk5Ie3kxnBTS4gZpQlabJ22rx0j6RGot",
	"Svw2mvaSHCYqFjipuHZl+ZFVOGbrg6hE8ivKnMTXGGd7LuoPXi7evL58tTi/eH3505v0iLv60uN8MeKu",
	"cSOMRTro+afXmKI9UUM8vcBzvTGg/xbd97McFj4aOan5wr6PXfxO58xi78N/pd0xeY4VHSiXrizZyt/T",
	"ksP08GAe+uDGcY7GJC7kTHSTFNLjOOoUDHJHwm5vvIBNkRUyQnrrbDH8+qUj/MOft5A2XzafqXk71Cqs",
	"1VD7xEKulT9fCo7d3QJ/8PvVrafAChvg3xmk5AbpQXAP+gHJCCX9NDybn819pNIomRaQwXn4yzvZFgHr",
	"rPSDNsilGtk8XcwKJa9yyJo5DI0X0Nh3Kt/6IK6kRRnimdal4OHE7LNRcvhw/8NPQWOYcOccDSehbYO/",
	"qR2M2OgbEC/m85dG0LonAuFtYhpzrF2ZUB+YwsULohjvNhEU71ieUM9TCsZVFaMtZHBbCJOgzLUS0ibC",
	"JFYlQVH/wGTCOFdO2jNvMLYxvpO9XWDp08w6p2c72GBE/E99J5xMgf0O/u9o0PY9ZPfjjr9f1stjEj0I",
	"fEz8SO2GTUSg0KS8mKoyGn0nas3oTI+wcjcemafs1fjE/z780n6ZjjjGtzS1W8Hhid7tDSdyzv4SF6Gm",
	"R3BKu0zWsf/JdO8ETlgi8fG5AR9MR34dCJ5zVLb7RTablYqzsvD+8OZrT+72gP3R+cYkbKWcDf6DtFtD",
	"Qp16Wf8dAAD//xMVaYi2DwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
