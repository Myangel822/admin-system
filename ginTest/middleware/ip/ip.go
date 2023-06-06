package ip

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetIp(c *gin.Context) string {

	ip := c.ClientIP()
	fmt.Println(ip)

	return ip
}
