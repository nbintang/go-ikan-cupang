package authentication

import (
	"ikan-cupang/config"
	"ikan-cupang/daos"
	"ikan-cupang/helper"
	"ikan-cupang/lib"
	"ikan-cupang/models"
	"ikan-cupang/schemas"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	db := config.DB
	otp, _ := helper.GenerateOTP()
	body := c.Locals("validatedBody").(*schemas.LoginSchema)
	defer (func() {
		if r := recover(); r != nil {
			helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid input")
		}
	})()
	hashedOTP := helper.HashToken(otp)

	user, err := daos.FindUserByEmail(body.Email, "id", "name", "email", "is_verified")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			nameGenerated, _, _ := strings.Cut(body.Email, "@")
			user = &models.User{
				Name:       nameGenerated,
				Email:      body.Email,
				IsVerified: false,
			}
			db.Create(&user)

		} else {
			return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get user")
		}
	}

	token := models.VerificationToken{
		Token:     hashedOTP,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		UserID:    user.ID,
	}
	db.Create(&token)

	err = lib.SendEmail(user.Email, "OTP Verification", "Your OTP is: "+otp)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to send email")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "OTP sent successfully, Please check your email",
	})
}

func VerifyOTP(c *fiber.Ctx) error {
	body := c.Locals("validatedBody").(*schemas.OTPSchema)
	db := config.DB

	existedUser, err := daos.FindUserByEmail(body.Email, "id", "name", "email", "is_verified", "role")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.ErrorHandler(c, fiber.StatusNotFound, "User not found")
		}
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get user")
	}

	existedToken, err := daos.FindExistedTokenByUserID(uint(existedUser.ID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.ErrorHandler(c, fiber.StatusNotFound, "Token not found")
		}
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get token")
	}

	if existedToken.ExpiresAt.Before(time.Now()) {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "OTP expired")
	}

	if !helper.CompareToken(body.OTP, existedToken.Token) {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid OTP")
	}

	existedUser.IsVerified = true
	db.Save(&existedUser)

	_, err = daos.DeleteAllTokenByUserID(uint(existedUser.ID))
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to delete token")
	}

	accessToken, refreshToken, err := lib.GenerateTokens(uint(existedUser.ID), string(existedUser.Role), existedUser.IsVerified)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to generate access token")
	}

	cookieOptions := fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   os.Getenv("NODE_ENV") != "development",
		SameSite: func() string {
			if os.Getenv("NODE_ENV") == "development" {
				return "lax"
			}
			return "none"
		}(),
		MaxAge: 60 * 60 * 24 * 1, // 1 days
	}

	if refreshToken != "" && accessToken != "" {
		c.Cookie(&cookieOptions)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":      true,
		"message":      "User verified successfully",
		"access_token": accessToken,
	})
}

func ResendOTP(c *fiber.Ctx) error {

	db := config.DB
	body := c.Locals("validatedBody").(*schemas.LoginSchema)
	otp, _ := helper.GenerateOTP()

	existedUser, err := daos.FindUserByEmail(body.Email, "id", "name", "email", "is_verified")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.ErrorHandler(c, fiber.StatusNotFound, "User not found")
		}
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get user")
	}

	_, err = daos.DeleteAllTokenByUserID(uint(existedUser.ID))
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to delete token")
	}

	hashedOTP := helper.HashToken(otp)

	token := models.VerificationToken{
		Token:     hashedOTP,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		UserID:    existedUser.ID,
	}

	err = db.Create(&token).Error
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to create token")
	}

	err = lib.SendEmail(existedUser.Email, "OTP Verification", "Your OTP is "+otp)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to send email")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "OTP sent successfully, Please check your email",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	cookies := c.Cookies("refresh_token")
	if cookies == "" {
		return helper.ErrorHandler(c, fiber.StatusUnauthorized, "Refresh token not found")
	}

	decodedPayload, err := lib.VerifyToken(cookies)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusUnauthorized, "Invalid refresh token")
	}

	claims, ok := decodedPayload.Claims.(jwt.MapClaims)

	if !ok {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Invalid token claims")
	}
	id, ok := claims["id"].(float64)

	if !ok {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Invalid ID claim")
	}

	existedUser, err := daos.FindUserByID(uint(id), "id", "name", "email", "is_verified", "role")
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get user")
	}
	accessToken, err := lib.GenerateAccessToken(uint(id), string(existedUser.Role), existedUser.IsVerified)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to generate access token")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":      true,
		"message":      "Token refreshed successfully",
		"access_token": accessToken,
	})
}
