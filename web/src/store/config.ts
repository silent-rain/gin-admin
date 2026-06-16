import { defineStore } from 'pinia'
import { langTitle } from '@/hooks/use-common'
import { i18n } from '@/lang'
import settings from '@/settings'
import { toggleHtmlClass } from '@/theme/utils'

export const useConfigStore = defineStore('config', {
  state: () => {
    return {
      language: settings.defaultLanguage,
      theme: settings.defaultTheme,
      size: settings.defaultSize,
    }
  },
  persist: {
    storage: localStorage,
    pick: ['language', 'theme', 'size'],
  },
  actions: {
    setTheme(data: string) {
      this.theme = data
      toggleHtmlClass(data)
    },
    setSize(data: string) {
      this.size = data
    },
    setLanguage(lang: string, title) {
      const { locale }: any = i18n.global
      this.language = lang
      locale.value = lang
      document.title = langTitle(title) // i18 page title
    },
  },
})
