package context

import (
	"源码/beego/session"
)

type BeegoInput struct {
	Context     *Context
	CruSession  session.Store
	pnames      []string
	pvalues		[]string
	data		map[interface{}]interface{} //store some values in this context when calling context in filter or controller.
	RequestBody   []byte
	RunMethod     string
	RunController reflect.Type
}
