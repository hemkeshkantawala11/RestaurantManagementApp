package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userID string) (err error) {
	userType := c.GetString("User_type")
	uid := c.GetString("User_id")
	fmt.Println("UserType : ", userType)

	err = nil

	if userType == "USER" && uid != userID {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("User_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}
