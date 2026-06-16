declare const directive: {
    path: string;
    component: import("vue").DefineComponent<{}, {}, any>;
    meta: {
        title: string;
        icon: string;
    };
    alwaysShow: boolean;
    children: {
        path: string;
        component: () => Promise<typeof import("*.vue")>;
        name: string;
        meta: {
            title: string;
        };
    }[];
};
export default directive;
