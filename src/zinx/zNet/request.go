package zNet

import ziface "zinxServer/src/zinx/zIface"

type Request struct {
	//已经和客户端建立好的链接
	conn ziface.IConnection
	//客户端请求的数据
	data []byte
}

// 获取当前链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//得到请求的消息数据
func (r *Request) GetData() []byte {
	return r.data
}
