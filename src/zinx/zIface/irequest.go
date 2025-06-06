package ziface

/*
IRequest 接口
实际上是把客户端请求的链接信息和数据包装到一个Request中，然后交给Server去处理
*/
type IRequest interface {
	// 获取当前链接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
}
