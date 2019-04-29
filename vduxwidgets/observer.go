package vduxwidgets

import (
	"github.com/gopherjs/vecty"
	"github.com/lmika/vedux"
)

type WidgetBuilder func(val interface{}) vecty.ComponentOrHTML

func Observe(store *vedux.Store, key string, builder WidgetBuilder) *ObserverWidget {
	initVal := store.Get(key)

	comp := &ObserverWidget{
		Core: vecty.Core{},
		builder: builder,
		val: initVal,
		widget: builder(initVal),
	}

	go func(subChan <-chan interface{}) {
		for v := range subChan {
			comp.setVal(v)
		}
	}(store.Observe(key).Chan())

	return comp
}


type ObserverWidget struct {
	vecty.Core
	builder WidgetBuilder
	val     interface{}
	widget  vecty.ComponentOrHTML
}

func (ow *ObserverWidget) setVal(newVal interface{}) {
	if ow.val == newVal {
		return
	}

	ow.val = newVal
	ow.widget = ow.builder(ow.val)

	vecty.Rerender(ow)
}

func (ow *ObserverWidget) Key() interface{} {
	return ow.val
}

func (ow *ObserverWidget) Render() vecty.ComponentOrHTML {
	return ow.widget
}
