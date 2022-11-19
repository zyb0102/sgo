package ioc

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Container interface {
	// Register 默认注册实例
	Register(ins any) Container
	// RegisterIFace 注册接口实例
	RegisterIFace(iFace any, ins any) Container
	// RegisterKey 通过Key注册实例
	RegisterKey(key string, ins any) Container
	// Bind 绑定值
	Bind(values ...any) Container
	// BindField 绑定字段
	BindField(values ...any) Container
	// Init 容器初始化函数
	Init() error
	// After 容器初始化后执行函数
	After(...func() error)
}
type container struct {
	mux       sync.RWMutex
	insStore  map[string]diIns
	resStore  []bindIns
	afterFunc []func() error
}

type bindKind int

const (
	field bindKind = iota + 1 // 字段绑定
	value                     // 值绑定
)

// 绑定的实例
type bindIns struct {
	k bindKind // 绑定实例类型
	i any      // 绑定的实例
}

type diKind int

const (
	pointer diKind = iota + 1 // 指针类型
	other                     // 其他类型
)

// 注入的实例
type diIns struct {
	k diKind // 注入实例类型
	i any    // 注入的实例
}

func (c *container) Register(ins any) Container {
	c.mux.Lock()
	defer c.mux.Unlock()
	// 获取实例类型名称
	insType := reflect.TypeOf(ins)
	// 判断是否是指针类型
	insVal := reflect.ValueOf(ins)
	if insVal.Kind() == reflect.Pointer {
		// 指针类型
		storeKey := fmt.Sprintf("%s:%s", "default", insType.String()[1:])
		c.insStore[storeKey] = diIns{
			k: pointer,
			i: ins,
		}
	} else {
		// 不是指针类型,取地址
		storeKey := fmt.Sprintf("%s:%s", "default", insType)
		c.insStore[storeKey] = diIns{
			k: other,
			i: ins,
		}
	}
	return c
}

func (c *container) RegisterKey(key string, ins any) Container {
	if key == "default" {
		panic("不能使用default作为key")
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	// 获取实例类型名称
	insType := reflect.TypeOf(ins)
	// 判断是否是指针类型
	insVal := reflect.ValueOf(ins)
	if insVal.Kind() == reflect.Pointer {
		// 指针类型
		storeKey := fmt.Sprintf("%s:%s", key, insType.String()[1:])
		c.insStore[storeKey] = diIns{
			k: pointer,
			i: ins,
		}
	} else {
		// 不是指针类型,取地址
		storeKey := fmt.Sprintf("%s:%s", key, insType)
		c.insStore[storeKey] = diIns{
			k: other,
			i: ins,
		}
	}
	return c
}

func (c *container) RegisterIFace(iFace any, ins any) Container {
	// 获取接口类型
	iFaceType := reflect.TypeOf(iFace)
	insType := reflect.TypeOf(ins)
	if !insType.Implements(iFaceType.Elem()) {
		panic(insType.String() + "未实现接口" + iFaceType.String()[1:])
	}
	storeKey := fmt.Sprintf("%s:%s", "default", iFaceType.String()[1:])

	if reflect.ValueOf(ins).Kind() == reflect.Pointer {
		// 指针类型
		c.insStore[storeKey] = diIns{
			k: pointer,
			i: ins,
		}
	} else {
		// 不是指针类型,取地址
		c.insStore[storeKey] = diIns{
			k: other,
			i: ins,
		}
	}
	return c
}

// Bind 绑定值
func (c *container) Bind(values ...any) Container {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, ins := range values {
		insVal := reflect.ValueOf(ins)
		if insVal.Kind() != reflect.Pointer {
			panic("绑定的实例必须是指针类型")
		}
		c.resStore = append(c.resStore, bindIns{
			k: value,
			i: ins,
		})
	}
	return c
}

// BindField 绑定字段
func (c *container) BindField(values ...any) Container {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, ins := range values {
		insVal := reflect.ValueOf(ins)
		if insVal.Kind() != reflect.Pointer {
			panic("绑定的实例必须是指针类型")
		}
		c.resStore = append(c.resStore, bindIns{
			k: field,
			i: ins,
		})
	}
	return c
}

func (c *container) Init() error {
	for _, bIns := range c.resStore {
		switch bIns.k {
		case field:
			err := c.bindField(bIns.i)
			if err != nil {
				return err
			}
		case value:
			err := c.bindValue(bIns.i)
			if err != nil {
				return err
			}
		}
	}
	for _, f := range c.afterFunc {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *container) bindValue(ins any) error {
	t := reflect.TypeOf(ins)
	storeKey := fmt.Sprintf("%s:%s", "default", t.String()[1:])
	di, ok := c.insStore[storeKey]
	if !ok {
		return errors.New("没有有效的实例进行绑定")
	}
	// 判断注入的实例是否是指针
	if di.k == pointer {
		diInsValue := reflect.ValueOf(di.i)
		// 实例的值进行复制
		insValue := reflect.ValueOf(ins)
		insValue.Elem().Set(diInsValue.Elem())

	} else {
		diInsValue := reflect.ValueOf(di.i)
		// 实例的值进行复制
		insValue := reflect.ValueOf(ins)
		insValue.Elem().Set(diInsValue)
	}
	return nil
}

func (c *container) bindField(ins any) error {
	insValue := reflect.ValueOf(ins)
	insType := reflect.TypeOf(ins)
	for i := 0; i < insValue.Elem().NumField(); i++ {
		fd := insValue.Elem().Field(i)
		tag := insType.Elem().Field(i).Tag.Get("di")
		if tag == "" {
			continue
		}
		// 判断字段类型是否是指针
		if fd.Kind() == reflect.Pointer {
			// 如果是指针类型
			storeKey := fmt.Sprintf("%s:%s", tag, fd.Type().String()[1:])
			di, ok := c.insStore[storeKey]
			if !ok {
				return errors.New("绑定实例失败")
			}
			// 如果是nil 进行初始化
			if fd.IsNil() {
				fd.Set(reflect.New(fd.Type().Elem()))
			}
			if di.k == pointer {
				fd.Set(reflect.ValueOf(di.i))
			} else {
				v := reflect.ValueOf(di.i)
				fd.Elem().Set(v)
			}
		} else if fd.Kind() == reflect.Interface {
			// 接口类型
			storeKey := fmt.Sprintf("%s:%s", tag, fd.Type())
			di, ok := c.insStore[storeKey]
			if !ok {
				return errors.New("绑定实例失败")
			}
			fd.Set(reflect.ValueOf(di.i))
		} else {
			storeKey := fmt.Sprintf("%s:%s", tag, fd.Type())
			di, ok := c.insStore[storeKey]
			if !ok {
				return errors.New("绑定实例失败")
			}
			fd.Set(reflect.ValueOf(di.i))
		}
	}
	return nil
}

func (c *container) After(f ...func() error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.afterFunc = append(c.afterFunc, f...)
}

func NewContainer() Container {
	return &container{
		insStore:  map[string]diIns{},
		resStore:  []bindIns{},
		afterFunc: []func() error{},
	}
}
