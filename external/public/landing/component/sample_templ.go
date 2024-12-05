// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package landing

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Sample() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section id=\"sample\" class=\"w-full py-12 md:py-20 bg-[#F8F8F8] flex flex-col space-y-10\"><div class=\"px-4 flex flex-col md:flex-row justify-center items-center md:gap-16 gap-6\"><div class=\"\"><h1 class=\"dancing-script text-pink-600\">Krusty Craft</h1><p class=\"josefin-slab mt-2 text-3xl font-semibold\">Bucket</p><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor<br>incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.</p></div><img class=\"h-64 md:h-max\" src=\"https://krustycraft.my.id/images/buket.svg\" alt=\"buket.svg\"></div><div class=\"px-4 flex flex-col md:flex-row justify-center items-center md:gap-16 gap-6\"><img class=\"h-64 md:h-max\" src=\"https://krustycraft.my.id/images/hampers.svg\" alt=\"hampers.svg\"><div class=\"\"><h1 class=\"dancing-script text-pink-600\">Krusty Craft</h1><p class=\"josefin-slab mt-2 text-3xl font-semibold\">Bucket</p><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor<br>incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.</p></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
