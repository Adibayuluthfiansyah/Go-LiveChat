package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user_id", "Adibayu")
		c.Next()
		// authHeader := c.GetHeader("Authorization")
		// if authHeader == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error:": " Access denied token not found"})
		// 	c.Abort()
		// 	return
		// }
		// parts := strings.Split(authHeader, " ")
		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied token not found"})
		// 	c.Abort()
		// 	return
		// }
		// tokenString := parts[1]
		// clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
		// claims, err := clerkjwt.Verify(c.Request.Context(), &clerkjwt.VerifyParams{
		// 	Token: tokenString,
		// })
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Access not valid or expired"})
		// 	c.Abort()
		// 	return
		// }

		// c.Set("user_id", claims.Subject)
		// c.Next()
	}
}
