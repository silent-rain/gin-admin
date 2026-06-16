declare const chartsRouter: {
    path: string;
    component: import("vue").DefineComponent<{}, {}, any>;
    redirect: string;
    name: string;
    meta: {
        title: string;
        icon: string;
    };
    children: {
        path: string;
        component: () => Promise<typeof import("*.vue")>;
        name: string;
        meta: {
            title: string;
            noCache: boolean;
        };
    }[];
};
export default chartsRouter;
