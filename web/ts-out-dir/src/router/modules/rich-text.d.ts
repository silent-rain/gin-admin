declare const richText: {
    path: string;
    component: import("vue").DefineComponent<{}, {}, any>;
    meta: {
        title: string;
        icon: string;
    };
    alwaysShow: boolean;
    children: {
        path: string;
        name: string;
        component: () => Promise<typeof import("*.vue")>;
        meta: {
            title: string;
            icon: string;
        };
    }[];
};
export default richText;
