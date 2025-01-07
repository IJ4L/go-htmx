// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package landing

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AllContact() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section id=\"contact\" class=\"mt-20\"><div class=\"flex flex-col xl:flex-row space-y-8 md:space-y-8 xl:space-y-0 xl:space-x-16 items-center justify-center\"><div class=\"flex flex-col justify-center items-center space-y-4 border-2 border-gray-200 p-6 h-[288px] w-[370px] flex-shrink-0\"><img src=\"https://krustycraft.my.id/icons/div.con-info-icon.svg\" alt=\"\" class=\"h-12 w-12\"><p>Lokasi Kami</p><p class=\"text-xs\">Bakung Garden - Gowa</p></div><div class=\"flex flex-col justify-center items-center space-y-4 border-2 border-gray-200 p-6 h-[288px] w-[370px] flex-shrink-0\"><img src=\"https://krustycraft.my.id/icons/div.con-info-icon-1.svg\" alt=\"\" class=\"h-12 w-12\"><p>Hubungi kami Kapan Saja</p><p class=\"text-xs\">Wa<br>0895321935303</p></div><div class=\"flex flex-col items-center justify-center space-y-4 border-2 border-gray-200 p-6 h-[288px] w-[370px] flex-shrink-0\"><img src=\"https://krustycraft.my.id/icons/div.con-info-icon-2.svg\" alt=\"\" class=\"h-12 w-12\"><p>Dukungan Secara Keseluruhan</p><p class=\"text-xs\">krustycraft@gmail.com</p></div></div><div><iframe class=\"w-full h-[450px] mt-16 md:px-32 md:py-20\" src=\"https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3973.3204409498458!2d119.48757549999999!3d-5.212245899999999!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dbee37be46edd1f%3A0x117c7c6ec4b4eb4a!2sKrusty%20Craft!5e0!3m2!1sen!2sid!4v1733968626166!5m2!1sen!2sid\" width=\"600\" height=\"450\" style=\"border:0;\" allowfullscreen=\"\" loading=\"lazy\" referrerpolicy=\"no-referrer-when-downgrade\"></iframe></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
