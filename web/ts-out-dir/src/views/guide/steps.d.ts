declare const steps: ({
    element: string;
    popover: {
        title: string;
        description: string;
        position: string;
    };
    padding?: undefined;
} | {
    element: string;
    popover: {
        title: string;
        description: string;
        position: string;
    };
    padding: number;
})[];
export default steps;
