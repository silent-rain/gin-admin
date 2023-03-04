/* settings */
export interface SettingsConfig {
  title: string;
  fixedHeader: boolean;
  sidebarLogo: boolean;
  showLeftMenu: boolean;
  ShowDropDown: boolean;
  showHamburger: boolean;
  isNeedLogin: boolean;
  isNeedNprogress: boolean;
  showTagsView: boolean;
  showFooter: boolean;
  tagsViewNum: number;
  openProdMock: boolean;
  errorLog: string | Array<string>;
  permissionMode: string;
  delWindowHeight: string;
  tmpToken: string;
  showNavbarTitle: boolean;
  showTopNavbar: boolean;
  mainNeedAnimation: boolean;
  viteBasePath: string;
  defaultLanguage: string;
  defaultSize: string;
  defaultTheme: string;
  plateFormId: number;
  defaultPassword: string;
}
