import router from '@/router'

export function getQueryParam() {
  const route: any = router.currentRoute
  if (route.value?.query.params) {
    return JSON.parse(route.value.query.params)
  }
}
// vue router
export function routerPush(name, params) {
  let data = {}
  if (params) {
    data = {
      params: JSON.stringify(params),
    }
  }
  else {
    data = {}
  }
  router.push({
    name,
    query: data,
  })
}
export function routerReplace(name, params) {
  let data = {}
  if (params) {
    data = {
      params: JSON.stringify(params),
    }
  }
  else {
    data = {}
  }
  router.replace({
    name,
    query: data,
  })
}

export function routeInfo() {
  return router.currentRoute
}
export function routerBack() {
  router.go(-1)
}
