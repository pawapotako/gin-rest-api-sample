package model

// SignInRequest godoc
// @Summary SignIn request structure
// @Description Detailed description for SignIn request
// @Model SignInRequest
type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
