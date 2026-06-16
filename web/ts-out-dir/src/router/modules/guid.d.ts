declare const guid: {
    path: string;
    component: import("vue").DefineComponent<{}, {}, any>;
    children: {
        path: string;
        component: () => Promise<typeof import("*.vue")>;
        name: string;
        meta: {
            title: string;
            icon: string;
        };
    }[];
};
export default guid;
