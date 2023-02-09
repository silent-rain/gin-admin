// Generated by 'unplugin-auto-import'
export {}
declare global {
  const EffectScope: typeof import('vue')['EffectScope']
  const addMenu: typeof import('../api/system/menu')['addMenu']
  const addRole: typeof import('../api/system/role')['addRole']
  const addUser: typeof import('../api/system/user')['addUser']
  const aoaToSheetXlsx: typeof import('../utils/excel')['aoaToSheetXlsx']
  const asyncRoutesByMenus: typeof import('../hooks/use-permission')['asyncRoutesByMenus']
  const axiosReq2: typeof import('../utils/axios-req')['axiosReq2']
  const axiosReq: typeof import('../utils/axios-req')['default']
  const batchDeleteMenu: typeof import('../api/system/menu')['batchDeleteMenu']
  const batchDeleteRole: typeof import('../api/system/role')['batchDeleteRole']
  const batchDeleteUser: typeof import('../api/system/user')['batchDeleteUser']
  const bus: typeof import('../utils/bus')['default']
  const buttonCodes: typeof import('../directives/button-codes')['default']
  const captchaVerify: typeof import('../api/system/login')['captchaVerify']
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
  const deleteMenu: typeof import('../api/system/menu')['deleteMenu']
  const deleteRole: typeof import('../api/system/role')['deleteRole']
  const deleteUser: typeof import('../api/system/user')['deleteUser']
  const directives: typeof import('../directives/index')['default']
  const effectScope: typeof import('vue')['effectScope']
  const elConfirm: typeof import('../hooks/use-element')['elConfirm']
  const elConfirmNoCancelBtn: typeof import('../hooks/use-element')['elConfirmNoCancelBtn']
  const elLoading: typeof import('../hooks/use-element')['elLoading']
  const elMessage: typeof import('../hooks/use-element')['elMessage']
  const elNotify: typeof import('../hooks/use-element')['elNotify']
  const filterAsyncRouter: typeof import('../hooks/use-permission')['filterAsyncRouter']
  const filterAsyncRouterByCodes: typeof import('../hooks/use-permission-bak')['filterAsyncRouterByCodes']
  const filterAsyncRoutesByMenuList: typeof import('../hooks/use-permission-bak')['filterAsyncRoutesByMenuList']
  const filterAsyncRoutesByRoles: typeof import('../hooks/use-permission-bak')['filterAsyncRoutesByRoles']
  const freshRouter: typeof import('../hooks/use-permission')['freshRouter']
  const getAllMenuTree: typeof import('../api/system/menu')['getAllMenuTree']
  const getAllRole: typeof import('../api/system/role')['getAllRole']
  const getAllUser: typeof import('../api/system/user')['getAllUser']
  const getCaptcha: typeof import('../api/system/login')['getCaptcha']
  const getCurrentInstance: typeof import('vue')['getCurrentInstance']
  const getCurrentScope: typeof import('vue')['getCurrentScope']
  const getLangInstance: typeof import('../hooks/use-common')['getLangInstance']
  const getMenuTree: typeof import('../api/system/menu')['getMenuTree']
  const getQueryParam: typeof import('../hooks/use-self-router')['getQueryParam']
  const getRoleList: typeof import('../api/system/role')['getRoleList']
  const getRoleMenuRelList: typeof import('../api/system/role-menu-rel')['getRoleMenuRelList']
  const getToken: typeof import('../utils/auth')['getToken']
  const getUserInfo: typeof import('../api/system/user')['getUserInfo']
  const getUserList: typeof import('../api/system/user')['getUserList']
  const h: typeof import('vue')['h']
  const importsExcel: typeof import('../utils/excel')['importsExcel']
  const inject: typeof import('vue')['inject']
  const isExternal: typeof import('../hooks/use-layout')['isExternal']
  const isProxy: typeof import('vue')['isProxy']
  const isReactive: typeof import('vue')['isReactive']
  const isReadonly: typeof import('vue')['isReadonly']
  const isRef: typeof import('vue')['isRef']
  const lang: typeof import('../directives/lang')['default']
  const langTitle: typeof import('../hooks/use-common')['langTitle']
  const login: typeof import('../api/system/login')['login']
  const logout: typeof import('../api/system/login')['logout']
  const longpress: typeof import('../directives/example/longpress.js')['default']
  const markRaw: typeof import('vue')['markRaw']
  const md5Encode: typeof import('../utils/md5')['md5Encode']
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
  const progressClose: typeof import('../hooks/use-basic')['progressClose']
  const progressStart: typeof import('../hooks/use-basic')['progressStart']
  const provide: typeof import('vue')['provide']
  const reactive: typeof import('vue')['reactive']
  const readonly: typeof import('vue')['readonly']
  const ref: typeof import('vue')['ref']
  const register: typeof import('../api/system/login')['register']
  const removeToken: typeof import('../utils/auth')['removeToken']
  const resetRouter: typeof import('../hooks/use-permission')['resetRouter']
  const resetState: typeof import('../hooks/use-permission')['resetState']
  const resetUserPwd: typeof import('../api/system/user')['resetUserPwd']
  const resizeHandler: typeof import('../hooks/use-layout')['resizeHandler']
  const resolveComponent: typeof import('vue')['resolveComponent']
  const resolveDirective: typeof import('vue')['resolveDirective']
  const rolesPermission: typeof import('../directives/roles-permission')['default']
  const routeInfo: typeof import('../hooks/use-self-router')['routeInfo']
  const routerBack: typeof import('../hooks/use-self-router')['routerBack']
  const routerPush: typeof import('../hooks/use-self-router')['routerPush']
  const routerReplace: typeof import('../hooks/use-self-router')['routerReplace']
  const searchUser: typeof import('../api/remote-search')['searchUser']
  const secret: typeof import('../utils/constant')['secret']
  const setToken: typeof import('../utils/auth')['setToken']
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
  const updateEmail: typeof import('../api/system/user')['updateEmail']
  const updateMenu: typeof import('../api/system/menu')['updateMenu']
  const updateMenuStatus: typeof import('../api/system/menu')['updateMenuStatus']
  const updatePhone: typeof import('../api/system/user')['updatePhone']
  const updateRole: typeof import('../api/system/role')['updateRole']
  const updateRoleMenuRel: typeof import('../api/system/role-menu-rel')['updateRoleMenuRel']
  const updateRoleStatus: typeof import('../api/system/role')['updateRoleStatus']
  const updateUser: typeof import('../api/system/user')['updateUser']
  const updateUserPwd: typeof import('../api/system/user')['updateUserPwd']
  const updateUserStatus: typeof import('../api/system/user')['updateUserStatus']
  const uploadAvatar: typeof import('../api/system/upload')['uploadAvatar']
  const useAttrs: typeof import('vue')['useAttrs']
  const useBasicStore: typeof import('../store/basic')['useBasicStore']
  const useConfigStore: typeof import('../store/config')['useConfigStore']
  const useCssModule: typeof import('vue')['useCssModule']
  const useCssVars: typeof import('vue')['useCssVars']
  const useElement: typeof import('../hooks/use-element')['useElement']
  const useErrorLog: typeof import('../hooks/use-error-log')['useErrorLog']
  const useLink: typeof import('vue-router')['useLink']
  const usePermissionStore: typeof import('../store/permission')['usePermissionStore']
  const useRoute: typeof import('vue-router')['useRoute']
  const useRouter: typeof import('vue-router')['useRouter']
  const useSlots: typeof import('vue')['useSlots']
  const useTable: typeof import('../hooks/use-table')['useTable']
  const useTagsViewStore: typeof import('../store/tags-view')['useTagsViewStore']
  const useUserStore: typeof import('../store/user')['useUserStore']
  const watch: typeof import('vue')['watch']
  const watchEffect: typeof import('vue')['watchEffect']
  const watchPostEffect: typeof import('vue')['watchPostEffect']
  const watchSyncEffect: typeof import('vue')['watchSyncEffect']
  const watermark: typeof import('../directives/example/watermark.js')['default']
  const waves: typeof import('../directives/example/waves.js')['default']
}
