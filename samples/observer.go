package main

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/lmika/vectytools/vduxwidgets"
	"github.com/lmika/vedux"
)

func main() {
	vecty.SetTitle("Markdown Demo")

	store := vedux.New()
	store.On("set", func(actCtx vedux.ActionContext) error {
		val := actCtx.Arg(0).(string)

		actCtx.Put("value", val)
		return nil
	})

	vecty.RenderBody(&pageWidget{
		store: store,
		valueObserver: vduxwidgets.Observe(store, "value", func(val interface{}) vecty.ComponentOrHTML {
			if val == nil {
				return vecty.Text("Nothing here")
			} else {
				return vecty.Text(fmt.Sprint(val))
			}
		}),
	})
}


type pageWidget struct {
	vecty.Core

	store			*vedux.Store
	valueObserver	*vduxwidgets.ObserverWidget
}

func (pw *pageWidget) onTextChange(e *vecty.Event) {
	val := e.Target.Get("value").String()
	pw.store.Dispatch("set", val)
}

func (pw *pageWidget) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(vecty.Text("Observer test")),

		elem.Paragraph(
			pw.valueObserver,
		),

		elem.Paragraph(
			elem.TextArea(
				vecty.Markup(event.Input(pw.onTextChange).PreventDefault()),
			),
		),
	)
}


