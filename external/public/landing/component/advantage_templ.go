// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package landing

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Advantage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section id=\"advantage\" class=\"w-full md:py-16 md:px-44 lg:px-64 bg-white\"><ul class=\"md:flex md:justify-center space-y-10 md:space-x-20 md:mt-0 mt-16 md:space-y-0\"><li class=\"flex flex-col md:flex-row md:gap-6 items-center justify-center\"><img src=\"https://krustycraft.my.id/images/Container.svg\" alt=\"Container.svg\" class=\"w-16 h-16 md:w-20 md:h-20\"><div class=\"flex flex-col items-center md:items-start mt-4 md:mt-0\"><h1 class=\"text-xl font-bold\">Free Delivery</h1><p class=\"text-gray-500 text-center text-xs md:text-left\">Free shipping around the<br>world for all orders over $120</p></div></li><li class=\"flex flex-col md:flex-row md:gap-6 items-center justify-center\"><img src=\"https://krustycraft.my.id/images/Freshness.svg\" alt=\"Freshness.svg\" class=\"w-16 h-16 md:w-20 md:h-20\"><div class=\"flex flex-col items-center md:items-start mt-4 md:mt-0\"><h1 class=\"text-xl font-bold\">Freshness</h1><p class=\"text-gray-500 text-center text-xs md:text-left\">You have freshness flowers<br>every single order</p></div></li><li class=\"flex flex-col md:flex-row md:gap-6 items-center justify-center\"><img src=\"https://krustycraft.my.id/images/Made-By-Artists.svg\" alt=\"Made-By-Artists.svg\" class=\"w-16 h-16 md:w-20 md:h-20\"><div class=\"flex flex-col items-center md:items-start mt-4 md:mt-0\"><h1 class=\"text-xl font-bold\">Made by Artist</h1><p class=\"text-gray-500 text-center text-xs md:text-left\">World for all made by artists<br>orders over now</p></div></li><li class=\"flex flex-col md:flex-row md:gap-6 items-center justify-center\"><img src=\"https://krustycraft.my.id/images/Online-Order.svg\" alt=\"Online-Order.svg\" class=\"w-16 h-16 md:w-20 md:h-20\"><div class=\"flex flex-col items-center md:items-start mt-4 md:mt-0\"><h1 class=\"text-xl font-bold\">Online Order</h1><p class=\"text-gray-500 text-center text-xs md:text-left\">Don’t worry you can order<br>Online by our Site</p></div></li></ul></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
