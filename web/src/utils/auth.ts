const TokenKey = 'basic-token';

// 获取 token
export function getToken() {
  return localStorage.getItem(TokenKey);
}

// 设置 token
export function setToken(token: string) {
  return localStorage.setItem(TokenKey, token);
}

// 移除 token
export function removeToken() {
  return localStorage.removeItem(TokenKey);
}
