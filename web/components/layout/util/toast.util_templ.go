// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.865
package util

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ToastGlobalState() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_ToastGlobalState_4311`,
		Function: `function __templ_ToastGlobalState_4311(){document.addEventListener('alpine:init', () => {
		Alpine.store('toast', {
			toastVisible: false,
			toastMessage: '',
			toastType: 'danger',
			
			displayToast(message, type) {
				this.toastVisible = true;
				this.toastMessage = message;
				this.toastType = type;
				setTimeout(() => {this.clearToast()}, 3000);
			},
 
			clearToast() {
				this.toastVisible = false;
				this.toastMessage = '';
			}
		});
	});
 
	document.addEventListener('htmx:afterRequest', (event) => {
		const contentType = event.detail.xhr.getResponseHeader("Content-Type");
		if (contentType !== 'application/json') {
			return;
		}
 
	    const responseData = event.detail.xhr.responseText;
		if (responseData === '') {
			return;
		}
 
		let toastType = 'success';
		const isResponseError = event.detail.xhr.status >= 400;
		if (isResponseError) {
			toastType = 'danger';
		}
 
		const parsedResponse = JSON.parse(responseData);
		if (parsedResponse.body === undefined || parsedResponse.body === '') {
			return;
		}
		const toastMessage = parsedResponse.body;
	
		Alpine.store('toast').displayToast(toastMessage, toastType);
	});
}`,
		Call:       templ.SafeScript(`__templ_ToastGlobalState_4311`),
		CallInline: templ.SafeScriptInline(`__templ_ToastGlobalState_4311`),
	}
}

func Toast() templ.Component {
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
		templ_7745c5c3_Err = ToastGlobalState().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"toast\" x-data x-show=\"$store.toast.toastVisible\" x-transition x-cloak un-cloak class=\"fixed bottom-5 right-5 mb-4 flex w-full max-w-xs items-center rounded-lg bg-white p-4 text-slate-500 shadow-xl dark:bg-slate-700 dark:text-gray-300\" role=\"alert\"><div x-show=\"$store.toast.toastType == &#39;danger&#39;\" class=\"inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-red-100 text-red-500 dark:bg-red-800 dark:text-red-200\"><svg class=\"h-5 w-5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\"><path d=\"M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 11.793a1 1 0 1 1-1.414 1.414L10 11.414l-2.293 2.293a1 1 0 0 1-1.414-1.414L8.586 10 6.293 7.707a1 1 0 0 1 1.414-1.414L10 8.586l2.293-2.293a1 1 0 0 1 1.414 1.414L11.414 10l2.293 2.293Z\"></path></svg> <span class=\"sr-only\">Error icon</span></div><div x-show=\"$store.toast.toastType == &#39;success&#39;\" class=\"inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-green-100 text-green-500 dark:bg-green-800 dark:text-green-200\"><svg class=\"h-5 w-5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\"><path d=\"M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 8.207-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414L9 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414Z\"></path></svg> <span class=\"sr-only\">Check icon</span></div><div class=\"toast-message ms-3 text-sm font-normal\" x-text=\"$store.toast.toastMessage\"></div><button @click=\"$store.toast.clearToast()\" type=\"button\" class=\"-mx-1.5 -my-1.5 ms-auto inline-flex h-8 w-8 items-center justify-center rounded-lg bg-white p-1.5 text-gray-400 hover:bg-gray-100 hover:text-gray-900 focus:ring-2 focus:ring-gray-300 dark:bg-slate-700 dark:text-gray-500 dark:hover:bg-slate-600 dark:hover:text-white\" aria-label=\"Close\"><span class=\"sr-only\">Close</span> <svg class=\"h-3 w-3\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 14 14\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6\"></path></svg></button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
