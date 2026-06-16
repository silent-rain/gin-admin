import { getWebSiteConfigList } from '@/api/data-center/config';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
NProgress.configure({ showSpinner: false });
export const progressStart = () => {
    NProgress.start();
};
export const progressClose = () => {
    NProgress.done();
};
export const fecthWebSiteConfigList = async () => {
    const basicStore = useBasicStore();
    const configHash = {};
    try {
        const resp = (await getWebSiteConfigList()).data;
        for (const item of resp.data_list) {
            configHash[item.key] = item;
        }
        basicStore.setWebSiteConfig(configHash);
    }
    catch (error) {
        console.log(error);
    }
};
