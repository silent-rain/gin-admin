import { UserTry } from '~/store/user';
export declare const useUserStore: import("pinia").StoreDefinition<"user", UserTry, {}, {
    userInfo(): Promise<{
        userInfo: any;
        roles: any;
        menus: any;
        permissions: any;
        codes: never[];
    }>;
    setToken(data: string): void;
    setUserInfo(data: any): void;
    resetState(): void;
    resetStateAndToLogin(): void;
}>;
