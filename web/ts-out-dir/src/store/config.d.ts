export declare const useConfigStore: import("pinia").StoreDefinition<"config", {
    language: string;
    theme: string;
    size: string;
}, {}, {
    setTheme(data: string): void;
    setSize(data: string): void;
    setLanguage(lang: string, title: any): void;
}>;
