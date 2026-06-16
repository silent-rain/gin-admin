import { defineStore } from 'pinia';
import setting from '@/settings';
export const useTagsViewStore = defineStore('tagsView', {
    state: () => {
        return {
            visitedViews: [],
        };
    },
    actions: {
        addVisitedView(view) {
            this.$patch((state) => {
                if (state.visitedViews.some((v) => v.path === view.path))
                    return;
                if (state.visitedViews.length >= setting.tagsViewNum) {
                    state.visitedViews.pop();
                    state.visitedViews.push({
                        ...view,
                        title: view.meta?.title || 'no-name',
                    });
                }
                else {
                    state.visitedViews.push({
                        ...view,
                        title: view.meta?.title || 'no-name',
                    });
                }
            });
        },
        async delVisitedView(view) {
            this.$patch((state) => {
                for (const [i, v] of state.visitedViews.entries()) {
                    if (v.path === view.path) {
                        state.visitedViews.splice(i, 1);
                        break;
                    }
                }
            });
            return [...this.visitedViews];
        },
        async delOthersVisitedViews(view) {
            this.$patch((state) => {
                state.visitedViews = state.visitedViews.filter((v) => {
                    return v.meta.affix || v.path === view.path;
                });
            });
            return [...this.visitedViews];
        },
        async delAllVisitedViews() {
            this.$patch((state) => {
                state.visitedViews = state.visitedViews.filter((tag) => tag.meta?.affix);
            });
            return [...this.visitedViews];
        },
    },
});
