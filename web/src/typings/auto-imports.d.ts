// Generated by 'unplugin-auto-import'
export {}
declare global {
  const EffectScope: typeof import('vue')['EffectScope']
  const axiosReq: typeof import('../utils/axios-req')['default']
  const bus: typeof import('../utils/bus')['default']
  const buttonCodes: typeof import('../directives/button-codes')['default']
  const casHandleChange: typeof import('../hooks/use-element')['casHandleChange']
  const clickoutside: typeof import('../directives/example/clickoutside.js')['default']
  const cloneDeep: typeof import('../hooks/use-common')['cloneDeep']
  const closeElLoading: typeof import('../hooks/use-element')['closeElLoading']
  const codesPermission: typeof import('../directives/codes-permission')['default']
  const commonUtil: typeof import('../utils/common-util')['default']
  const computed: typeof import('vue')['computed']
  const copy: typeof import('../directives/example/copy.js')['default']
  const copyValueToClipboard: typeof import('../hooks/use-common')['copyValueToClipboard']
  const createApp: typeof import('vue')['createApp']
  const customRef: typeof import('vue')['customRef']
  const debounce: typeof import('../directives/example/debounce.js')['default']
  const defineAsyncComponent: typeof import('vue')['defineAsyncComponent']
  const defineComponent: typeof import('vue')['defineComponent']
  const directives: typeof import('../directives/index')['default']
  const effectScope: typeof import('vue')['effectScope']
  const elConfirm: typeof import('../hooks/use-element')['elConfirm']
  const elConfirmNoCancelBtn: typeof import('../hooks/use-element')['elConfirmNoCancelBtn']
  const elLoading: typeof import('../hooks/use-element')['elLoading']
  const elMessage: typeof import('../hooks/use-element')['elMessage']
  const elNotify: typeof import('../hooks/use-element')['elNotify']
  const filterAsyncRouter: typeof import('../hooks/use-permission')['filterAsyncRouter']
  const filterAsyncRouterByCodes: typeof import('../hooks/use-permission')['filterAsyncRouterByCodes']
  const filterAsyncRoutesByMenuList: typeof import('../hooks/use-permission')['filterAsyncRoutesByMenuList']
  const filterAsyncRoutesByRoles: typeof import('../hooks/use-permission')['filterAsyncRoutesByRoles']
  const freshRouter: typeof import('../hooks/use-permission')['freshRouter']
  const getCurrentInstance: typeof import('vue')['getCurrentInstance']
  const getCurrentScope: typeof import('vue')['getCurrentScope']
  const getLangInstance: typeof import('../hooks/use-common')['getLangInstance']
  const getQueryParam: typeof import('../hooks/use-self-router')['getQueryParam']
  const h: typeof import('vue')['h']
  const inject: typeof import('vue')['inject']
  const isExternal: typeof import('../hooks/use-layout')['isExternal']
  const isProxy: typeof import('vue')['isProxy']
  const isReactive: typeof import('vue')['isReactive']
  const isReadonly: typeof import('vue')['isReadonly']
  const isRef: typeof import('vue')['isRef']
  const lang: typeof import('../directives/lang')['default']
  const langTitle: typeof import('../hooks/use-common')['langTitle']
  const loginOutReq: typeof import('../api/user')['loginOutReq']
  const loginReq: typeof import('../api/user')['loginReq']
  const longpress: typeof import('../directives/example/longpress.js')['default']
  const markRaw: typeof import('vue')['markRaw']
  const mockAxiosReq: typeof import('../utils/mock-axios-req')['default']
  const nextTick: typeof import('vue')['nextTick']
  const onActivated: typeof import('vue')['onActivated']
  const onBeforeMount: typeof import('vue')['onBeforeMount']
  const onBeforeRouteLeave: typeof import('vue-router')['onBeforeRouteLeave']
  const onBeforeRouteUpdate: typeof import('vue-router')['onBeforeRouteUpdate']
  const onBeforeUnmount: typeof import('vue')['onBeforeUnmount']
  const onBeforeUpdate: typeof import('vue')['onBeforeUpdate']
  const onDeactivated: typeof import('vue')['onDeactivated']
  const onErrorCaptured: typeof import('vue')['onErrorCaptured']
  const onMounted: typeof import('vue')['onMounted']
  const onRenderTracked: typeof import('vue')['onRenderTracked']
  const onRenderTriggered: typeof import('vue')['onRenderTriggered']
  const onScopeDispose: typeof import('vue')['onScopeDispose']
  const onServerPrefetch: typeof import('vue')['onServerPrefetch']
  const onUnmounted: typeof import('vue')['onUnmounted']
  const onUpdated: typeof import('vue')['onUpdated']
  const progressClose: typeof import('../hooks/use-permission')['progressClose']
  const progressStart: typeof import('../hooks/use-permission')['progressStart']
  const provide: typeof import('vue')['provide']
  const reactive: typeof import('vue')['reactive']
  const readonly: typeof import('vue')['readonly']
  const ref: typeof import('vue')['ref']
  const resetRouter: typeof import('../hooks/use-permission')['resetRouter']
  const resetState: typeof import('../hooks/use-permission')['resetState']
  const resizeHandler: typeof import('../hooks/use-layout')['resizeHandler']
  const resolveComponent: typeof import('vue')['resolveComponent']
  const resolveDirective: typeof import('vue')['resolveDirective']
  const rolesPermission: typeof import('../directives/roles-permission')['default']
  const routeInfo: typeof import('../hooks/use-self-router')['routeInfo']
  const routerBack: typeof import('../hooks/use-self-router')['routerBack']
  const routerPush: typeof import('../hooks/use-self-router')['routerPush']
  const routerReplace: typeof import('../hooks/use-self-router')['routerReplace']
  const searchUser: typeof import('../api/remote-search')['searchUser']
  const shallowReactive: typeof import('vue')['shallowReactive']
  const shallowReadonly: typeof import('vue')['shallowReadonly']
  const shallowRef: typeof import('vue')['shallowRef']
  const sleepTimeout: typeof import('../hooks/use-common')['sleepTimeout']
  const storeToRefs: typeof import('pinia/dist/pinia')['storeToRefs']
  const toRaw: typeof import('vue')['toRaw']
  const toRef: typeof import('vue')['toRef']
  const toRefs: typeof import('vue')['toRefs']
  const transactionList: typeof import('../api/remote-search')['transactionList']
  const triggerRef: typeof import('vue')['triggerRef']
  const unref: typeof import('vue')['unref']
  const useAttrs: typeof import('vue')['useAttrs']
  const useBasicStore: typeof import('../store/basic')['useBasicStore']
  const useConfigStore: typeof import('../store/config')['useConfigStore']
  const useCssModule: typeof import('vue')['useCssModule']
  const useCssVars: typeof import('vue')['useCssVars']
  const useElement: typeof import('../hooks/use-element')['useElement']
  const useErrorLog: typeof import('../hooks/use-error-log')['useErrorLog']
  const useLink: typeof import('vue-router')['useLink']
  const useRoute: typeof import('vue-router')['useRoute']
  const useRouter: typeof import('vue-router')['useRouter']
  const useSlots: typeof import('vue')['useSlots']
  const useTable: typeof import('../hooks/use-table')['useTable']
  const useTagsViewStore: typeof import('../store/tags-view')['useTagsViewStore']
  const userInfoReq: typeof import('../api/user')['userInfoReq']
  const watch: typeof import('vue')['watch']
  const watchEffect: typeof import('vue')['watchEffect']
  const watchPostEffect: typeof import('vue')['watchPostEffect']
  const watchSyncEffect: typeof import('vue')['watchSyncEffect']
  const watermark: typeof import('../directives/example/watermark.js')['default']
  const waves: typeof import('../directives/example/waves.js')['default']
}
