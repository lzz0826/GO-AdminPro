package jwt

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

// TODO 帶移致配置文件
var jwtKey = []byte("your_secret_key")

// Claims 是JWT 的內容，可以自定義需要的欄位

type Claims struct {
	Username  string `json:"username"`
	AdminName string `json:"adminName"`
	Nickname  string `json:"nickname"`
	AdminId   string `json:"adminId"`
	ChannelID string `json:"channelID"`

	jwt.StandardClaims
}

func LoginHandler(adminDao adminDao.AdminDAO) (tokenStr string, err error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username:  adminDao.Username,
		AdminName: adminDao.AdminName,
		Nickname:  adminDao.Nickname,
		AdminId:   adminDao.ID,
		ChannelID: adminDao.ChannelID,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// 驗證簽名算法是否為HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// 驗證必要的參數
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, fmt.Errorf("Invalid token claims")
			}

			username, ok := claims["username"].(string)
			if !ok || username == "" {
				return nil, fmt.Errorf("Missing required parameter: username")
			}
			adminId, ok := claims["adminId"].(string)
			if !ok || adminId == "" {
				return nil, fmt.Errorf("Missing required parameter: adminId")
			}

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		nickname := claims["nickname"].(string)
		adminId := claims["adminId"].(string)
		channelID := claims["channelID"].(string)

		// 将用户信息保存到上下文中，供后续处理函数使用
		c.Set("username", username)
		c.Set("nickname", nickname)
		c.Set("adminId", adminId)
		c.Set("channelID", channelID)

		c.Next()
	}
}

//
//func GetCurrentAdminId(c *gin.Context) string {
//	// 從上下文中獲取管理員ID
//	adminIdInterface, ok := c.Get("adminId")
//	if !ok {
//		return ""
//	}
//	adminId, ok := adminIdInterface.(string)
//	if !ok {
//		return ""
//	}
//	return adminId
//}
