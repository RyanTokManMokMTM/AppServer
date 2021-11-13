package request

type (
	RegisterRequest struct {
		Email string `form:"email" binding:"required" validate:"required,email"`
		Password string `form:"password" binding:"required"`
		ConfirmPassword string `form:"confirmPassword" binding:"required"`
		FirstName string `form:"firstName" binding:"required"`
		LastName string `form:"LastName" binding:"required"`
	}

	LoginRequest struct {
		Email string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
)
