package jwt

import (
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/controller"
	"AdminPro/dao/model/adminDao"
	adminServerDao "AdminPro/dao/service/admin"
	"AdminPro/server/tonke"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte(viper.GetString("jwt.jwt_key"))
var tokenTimeOut = viper.GetInt32("jwt.token_timeOut")
var authHeader = viper.GetString("jwt.jwt_auth")

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
	expirationTime := time.Now().Add(time.Duration(tokenTimeOut) * time.Minute)

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
		//使用自訂的 myContext
		//ctx := &myContext.MyContext{Context: c, Trace: c.GetHeader(viper.GetString("http.headerTrace"))}

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Authorization header is missing"})
			c.Abort()
			return
		}
		tokenString := c.GetHeader(authHeader)
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
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
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		nickname := claims["nickname"].(string)
		adminId := claims["adminId"].(string)
		channelID := claims["channelID"].(string)

		//驗證token是否有效 Redis
		if !tonke.VerifyOnlineToken(c, adminId, tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "token is Expired or Invalid"})
			c.Abort()
			return
		}

		// 将用户信息保存到上下文中，供后续处理函数使用
		c.Set("username", username)
		c.Set("nickname", nickname)
		c.Set("adminId", adminId)
		c.Set("channelID", channelID)

		c.Next()
	}
}

// GetJwtTokenData 解析Token Data
func GetJwtTokenData(tokenString string) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 驗證簽名算法是否為HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)

	result := Claims{
		Username:  claims["username"].(string),
		AdminName: claims["adminName"].(string),
		Nickname:  claims["nickname"].(string),
		AdminId:   claims["adminId"].(string),
		ChannelID: claims["channelID"].(string),
	}
	return &result, nil
}

// 刷新Token
func RefreshToken(c *gin.Context, token string) (tokenStr string, globalError *utils.GlobalError) {
	adminId := controller.GetCurrentAdminId(c)
	if !tonke.VerifyOnlineToken(c, adminId, token) {
		tokenError := tool.TokenError
		return "", utils.NewGlobalError(errors.New(tokenError.Msg), &tokenError)
	}
	adminDAO, err := adminServerDao.GetAdminById(adminId)
	if err != nil {
		notFindAdmin := tool.NotFindAdmin
		return "", utils.NewGlobalError(err, &notFindAdmin)
	}
	str, err := LoginHandler(adminDAO)
	if err != nil {
		createTokenError := tool.CreateTokenError
		return "", utils.NewGlobalError(err, &createTokenError)
	}
	err = tonke.SetTokenToRides(c, adminId, str)
	if err != nil {
		setRedisKeyError := tool.SetRedisKeyError
		return "", utils.NewGlobalError(err, &setRedisKeyError)
	}
	return str, nil
}
