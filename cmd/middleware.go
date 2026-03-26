package main

import "github.com/gin-gonic/gin"

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetSessionString(c, "userId")
		if userId == "" {
			c.Abort()
			return
		}

		_, err := h.users.GetUserById(userId)
		if err != nil {
			ClearSession(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
