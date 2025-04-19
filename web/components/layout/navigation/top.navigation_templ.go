// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package navigation

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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<nav x-data=\"{ open: false, mobileOpen: false }\" class=\"bg-slate-50/45 text-slate-800 shadow-sm sticky top-0 backdrop-blur-md z-50\"><!-- Desktop Navbar --><div class=\"container mx-auto px-4\"><div class=\"flex justify-between items-center h-16\"><!-- Logo/Brand --><div class=\"flex-shrink-0\"><a href=\"/\" class=\"text-xl font-bold ease-in-out hover:text-slate-700\">Spending Log</a></div><!-- Desktop Menu Items --><div class=\"hidden md:flex space-x-4\"><a href=\"/\" hx-get=\"/\" hx-target=\"body\" hx-swap=\"outerHTML\" class=\"inline-flex items-center justify-center h-10 px-4 py-2 text-sm font-medium transition-colors rounded-md hover:bg-neutral-100\">Home</a> <a href=\"/countries\" hx-get=\"/countries\" hx-target=\"body\" hx-swap=\"outerHTML\" class=\"inline-flex items-center justify-center h-10 px-4 py-2 text-sm font-medium transition-colors border-neutral-200 rounded-md hover:bg-neutral-100\">Countries</a><!-- Dropdown Menu --><div x-data=\"{ dropdownOpen: false }\" class=\"relative\"><button @click=\"dropdownOpen = !dropdownOpen\" class=\"inline-flex items-center justify-center h-10 px-4 py-2 text-sm font-medium transition-colors rounded-md hover:bg-neutral-100\">More <svg class=\"w-4 h-4 ml-1\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 9l-7 7-7-7\"></path></svg></button><div x-show=\"dropdownOpen\" @click.away=\"dropdownOpen = false\" class=\"absolute right-0 mt-2 w-48 bg-white text-slate-800 rounded-md shadow-lg py-1 z-50\"><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100\">Settings</a> <a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100\">Profile</a></div></div></div><!-- Mobile menu button --><div class=\"md:hidden\"><button @click=\"mobileOpen = !mobileOpen\" class=\"p-2 rounded-md hover:bg-neutral-200 focus:outline-none focus:ring-2 focus:ring-neutral-300\"><svg class=\"w-6 h-6\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path x-show=\"!mobileOpen\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h16\"></path> <path x-show=\"mobileOpen\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></div></div></div><!-- Mobile Menu --><div x-show=\"mobileOpen\" @click.away=\"mobileOpen = false\" class=\"md:hidden bg-white border-t border-neutral-200\" x-transition:enter=\"transition ease-out duration-100\" x-transition:enter-start=\"opacity-0 scale-95\" x-transition:enter-end=\"opacity-100 scale-100\" x-transition:leave=\"transition ease-in duration-75\" x-transition:leave-start=\"opacity-100 scale-100\" x-transition:leave-end=\"opacity-0 scale-95\"><div class=\"px-4 py-3 space-y-1\"><a href=\"/\" hx-get=\"/\" hx-target=\"body\" hx-swap=\"outerHTML\" class=\"block px-4 py-2 text-sm font-medium text-slate-700 hover:bg-neutral-100 rounded-md\">Home</a> <a href=\"/countries\" hx-get=\"/countries\" hx-target=\"body\" hx-swap=\"outerHTML\" class=\"block px-4 py-2 text-sm font-medium text-slate-700 hover:bg-neutral-100 rounded-md\">Countries</a> <a href=\"#\" class=\"block px-4 py-2 text-sm font-medium text-slate-700 hover:bg-neutral-100 rounded-md\">Settings</a></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
