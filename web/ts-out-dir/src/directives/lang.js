import { watch } from 'vue';
import { storeToRefs } from 'pinia';
import { langTitle } from '@/hooks/use-common';
import { useConfigStore } from '@/store/config';
const componentToProps = {
    ElInput: 'placeholder',
    ElTableColumn: 'label',
};
function checkPermission(el, { value }) {
    let saveOriginTitle = '';
    const { language } = storeToRefs(useConfigStore());
    const name = el.__vueParentComponent?.type?.name;
    const nameTitle = el.__vueParentComponent?.props[componentToProps[name]];
    saveOriginTitle = nameTitle || el.innerText;
    watch(() => language.value, () => {
        if (name?.startsWith('EL')) {
            if (Object.keys(componentToProps).includes(name)) {
                const { props } = el.__vueParentComponent;
                props[componentToProps[name]] = langTitle(saveOriginTitle);
            }
            else {
                el.innerText = langTitle(saveOriginTitle);
            }
        }
        else {
            if (el.__vnode?.type) {
                el.innerText = langTitle(saveOriginTitle);
            }
        }
    }, { immediate: true });
}
export default {
    mounted(el, binding) {
        checkPermission(el, binding);
    },
};
