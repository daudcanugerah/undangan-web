package usecase

import (
	"context"
	"errors"
	"time"

	"basic-service/domain"
	"basic-service/interface/sql"

	"braces.dev/errtrace"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	UserManager *sql.UserRepository
	JWTSecret   string
}

func NewAuth(userManager *sql.UserRepository, jwtSecret string) *Auth {
	return &Auth{
		UserManager: userManager,
		JWTSecret:   jwtSecret,
	}
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotActive      = errors.New("user is not active")
)

// Claims represents the JWT claims
type Claims struct {
	UserID string          `json:"user_id"`
	Email  string          `json:"email"`
	Role   domain.RoleType `json:"role"`
	jwt.RegisteredClaims
}

func GetClaimFromContext(ctx context.Context) (*Claims, error) {
	claims, ok := ctx.Value("claims").(*Claims)
	if !ok {
		return nil, errtrace.Wrap(errors.New("no claims in context"))
	}
	return claims, nil
}

func (a *Auth) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), errtrace.Wrap(err)
}

func (a *Auth) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *Auth) generateJWT(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 1 day expiration

	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return errtrace.Wrap2(token.SignedString([]byte(a.JWTSecret)))
}

func (a *Auth) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.UserManager.GetEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.UserNotFoundErr) {
			return "", errtrace.Wrap(ErrInvalidCredentials)
		}
		return "", errtrace.Wrap(err)
	}

	if !user.IsActive {
		return "", errtrace.Wrap(ErrUserNotActive)
	}

	if !a.checkPasswordHash(password, user.Password) {
		return "", errtrace.Wrap(ErrInvalidCredentials)
	}

	token, err := a.generateJWT(user)
	if err != nil {
		return "", errtrace.Wrap(err)
	}

	return token, nil
}

func (a *Auth) Register(ctx context.Context, user domain.User) error {
	// Hash the password before storing
	hashedPassword, err := a.hashPassword(user.Password)
	if err != nil {
		return errtrace.Wrap(err)
	}
	user.Password = hashedPassword

	// Set default role if not specified
	if user.Role == 0 {
		user.Role = domain.RoleUser
	}

	// Set default active status if not specified
	if !user.IsActive && user.Role == domain.RoleUser {
		user.IsActive = true
	}

	return errtrace.Wrap(a.UserManager.Create(ctx, &user))
}

func (a *Auth) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.JWTSecret), nil
	})
	if err != nil {
		return nil, errtrace.Wrap(err)
	}

	if !token.Valid {
		return nil, errtrace.Wrap(errors.New("invalid token"))
	}

	return claims, nil
}

func (a *Auth) Me(ctx context.Context) (domain.SafeUser, error) {
	claims, err := GetClaimFromContext(ctx)
	if err != nil {
		return domain.SafeUser{}, errtrace.Wrap(errors.New("invalid token claims"))
	}

	user, err := a.UserManager.GetUserByID(ctx, claims.UserID)
	if err != nil {
		return domain.SafeUser{}, errtrace.Wrap(err)
	}

	return domain.SafeUser{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Profile:   user.Profile,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
