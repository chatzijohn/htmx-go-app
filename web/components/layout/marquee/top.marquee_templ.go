// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package marquee

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Top() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div x-data=\"{\n      messages: [\n        &#39;🌟 New feature released! Check out version 2.0&#39;,\n        &#39;🎉 Special offer: 30% discount this week only&#39;,\n        &#39;📢 Join our community of 10,000+ developers&#39;,\n        &#39;🚀 Performance improvements live now&#39;,\n        &#39;💡 Pro tip: Try our new keyboard shortcuts&#39;\n      ]\n    }\" class=\"relative overflow-hidden bg-gray-50 py-2\"><!-- Left fade --><div class=\"absolute inset-y-0 left-0 z-10 w-20 bg-gradient-to-r from-gray-50 to-transparent\"></div><!-- Right fade --><div class=\"absolute inset-y-0 right-0 z-10 w-20 bg-gradient-to-l from-gray-50 to-transparent\"></div><!-- Marquee track --><div class=\"marquee-track hover:pause-animation\" x-ref=\"marqueeContainer\"><div class=\"flex gap-x-20 animate-scroll whitespace-nowrap\"><!-- Original messages --><template x-for=\"(message, index) in messages\" :key=\"index\"><div class=\"flex items-center px-2\"><span x-text=\"message\" class=\"text-md font-medium text-gray-700 dark:text-gray-300\"></span></div></template><!-- Duplicate messages for seamless scroll --><template x-for=\"(message, index) in messages\" :key=\"&#39;dup-&#39; + index\"><div class=\"flex items-center px-2\"><span x-text=\"message\" class=\"text-sm font-medium text-gray-700 dark:text-gray-300\"></span></div></template></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
