/*
	返回消息结构
	创建者: Constant 2019/12/31
	版本: v1.0
*/
package Common

type Resp struct {
	Status string `json:"status" cmnt:"返回状态"`
	Msg string `json:"message" cmnt:"提示信息"`
	Data interface{} `json:"data" cmnt:"返回内容"`
}