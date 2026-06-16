import buttonCodes from './button-codes'
import codesPermission from './codes-permission'
import clickoutside from './example/clickoutside'
import copy from './example/copy'

import debounce from './example/debounce'
import longpress from './example/longpress'
import watermark from './example/watermark'
import waves from './example/waves.js'
import lang from './lang'
import rolesPermission from './roles-permission'

export default function (app) {
  app.directive('ButtonCodes', buttonCodes)
  app.directive('CodesPermission', codesPermission)
  app.directive('RolesPermission', rolesPermission)
  app.directive('lang', lang)

  // example
  app.directive('copy', copy)
  app.directive('longpress', longpress)
  app.directive('debounce', debounce)
  app.directive('watermark', watermark)
  app.directive('waves', waves)
  app.directive('clickoutside', clickoutside)
}
