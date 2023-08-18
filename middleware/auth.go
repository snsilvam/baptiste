package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"baptiste.com/helpers"
	"github.com/gin-gonic/gin"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

const (
	missingJWTErrorMessage       = "Requires authentication"
	invalidJWTErrorMessage       = "Bad credentials"
	permissionDeniedErrorMessage = "Permission denied"
)

type CustomClaims struct {
	Permissions []string `json:"permissions"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (c CustomClaims) HasPermissions(expectedClaims []string) bool {
	if len(expectedClaims) == 0 {
		return false
	}
	for _, scope := range expectedClaims {
		if !helpers.Contains(c.Permissions, scope) {
			return false
		}
	}
	return true
}

func ValidatePermissions(expectedClaims []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if !ok {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				map[string]string{"message": "Failed to get validated JWT claims."},
			)
			return
		}

		claims := token.CustomClaims.(*CustomClaims)
		if !claims.HasPermissions(expectedClaims) {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				map[string]string{"message": permissionDeniedErrorMessage},
			)
			return
		}

		c.Next()
	}
}

func ValidateJWT(audience, domain string) gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(ctx *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.Request = r
			ctx.Next()
		}

		middleware.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

		if encounteredError {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "JWT is invalid."},
			)
		}
	}
}
