package zNet

import ziface "zinxServer/src/zinx/zIface"

//实现router时，先嵌入BaseRouter基类，然后根据需求对这个基类的方法进行重写
type BaseRouter struct {
}

//这里之所以BaseRouter的方法都为空，是因为有的Router不希望有PreHandle和PostHandle这两个业务，所有的Router继承BaseRouter的好处就是，不需要实现PreHandle和PostHandle方法
//在处理conn业务之前的钩子方法Hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {
}

//在处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {
}

//在处理conn业务之后的钩子方法Hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {
}
