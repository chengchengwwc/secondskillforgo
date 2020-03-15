package common

import (
	"net/http"
)

//声明一个新的数据类型（函数类型）
type FilterHandle func(rw http.ResponseWriter, req *http.Request) error

type WebHandle func(rw http.ResponseWriter, req *http.Request)

type Filter struct {
	filterMap map[string]FilterHandle
}

//初始化
func NewFilter() *Filter {
	return &Filter{filterMap: make(map[string]FilterHandle)}
}

func (f *Filter) RegisterFilterUri(uri string, handler FilterHandle) {
	f.filterMap[uri] = handler
}

func (f *Filter) GetFilterHandle(uri string) FilterHandle {
	return f.filterMap[uri]
}

func (f *Filter) Handle(webHandle WebHandle) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		for path, handle := range f.filterMap {
			if path == req.RequestURI {
				err := handle(rw, req)
				if err != nil {
					_, _ = rw.Write([]byte(err.Error()))
					return
				}
				break
			}
		}
		//执行正常注册的函数
		webHandle(rw, req)

	}

}
