import { defineStore } from 'pinia';
import { langTitle } from '@/hooks/use-common';
import settings from '@/settings';
import { toggleHtmlClass } from '@/theme/utils';
import { i18n } from '@/lang';
export const useConfigStore = defineStore('config', {
    state: () => {
        return {
            language: settings.defaultLanguage,
            theme: settings.defaultTheme,
            size: settings.defaultSize,
        };
    },
    persist: {
        storage: localStorage,
        pick: ['language', 'theme', 'size'],
    },
    actions: {
        setTheme(data) {
            this.theme = data;
            toggleHtmlClass(data);
        },
        setSize(data) {
            this.size = data;
        },
        setLanguage(lang, title) {
            const { locale } = i18n.global;
            this.language = lang;
            locale.value = lang;
            document.title = langTitle(title);
        },
    },
});
