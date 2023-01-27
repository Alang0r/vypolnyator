package api

// import (
// 	"fmt"

// 	"github.com/Alang0r/vypolnyator/pkg/error"
// 	"github.com/Alang0r/vypolnyator/pkg/service"
// 	"github.com/gin-gonic/gin"
// )

// func init() {
// 	service.RegisterRequest("/", (*ReqTestJson)(nil))
// }

// type ReqTestJson struct {
// 	Name string
// 	Age uint64
// }

// func (r *ReqTestJson) Run(c *gin.Context) (service.Reply, error.Error) {
// 	req := ReqTestJson{}
// 	c.Bind(&req)

// 	fmt.Println(req)

// 	return nil, *error.New().SetCode(error.ErrCodeNone)
// }
