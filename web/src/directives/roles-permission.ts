import { useUserStore } from '@/store/user'

function checkPermission(el, { value }) {
  if (value && Array.isArray(value)) {
    if (value.length) {
      const permissionRoles = value
      const hasPermission = useUserStore().roles?.some(role =>
        permissionRoles.includes(role.name),
      )
      if (!hasPermission)
        el.parentNode && el.parentNode.removeChild(el)
    }
  }
  else {
    throw new Error('need roles! Like v-permission="[\'admin\',\'editor\']"')
  }
}
export default {
  mounted(el, binding) {
    checkPermission(el, binding)
  },
  componentUpdated(el, binding) {
    checkPermission(el, binding)
  },
}
