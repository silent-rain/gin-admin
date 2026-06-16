import { createI18n } from 'vue-i18n';
import en from './en';
import zh from './zh';
import settings from '@/settings';
const messages = { en, zh };
const localeData = {
    globalInjection: true,
    legacy: false,
    locale: settings.defaultLanguage,
    messages,
};
export const i18n = createI18n(localeData);
export const setupI18n = {
    install(app) {
        app.use(i18n);
    },
};
