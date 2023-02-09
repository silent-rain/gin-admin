import NProgress from 'nprogress'; // 进度条
import 'nprogress/nprogress.css'; // 进度条样式

NProgress.configure({ showSpinner: false });
// 开始进度条
export const progressStart = () => {
  NProgress.start();
};
// 关闭进度条
export const progressClose = () => {
  NProgress.done();
};
