export const Local = {
  // 设置永久缓存
  set(key: string, value: any) {
      window.localStorage.setItem(key, JSON.stringify(value))
  },
  // 获取永久缓存
  get(key: string) {
      return JSON.parse(window.localStorage.getItem(key) || '{}')
  },
  // 删除永久缓存
  remove(key: string) {
      window.localStorage.removeItem(key)
  },
  // 清空永久缓存
  clear() {
      window.localStorage.clear()
  },
};

export const Session = {
  // 设置临时缓存
  set(key: string, value: any) {
      window.sessionStorage.setItem(key, JSON.stringify(value))
  },
  // 获取临时缓存
  get(key: string) {
      return JSON.parse(window.sessionStorage.getItem(key) || 'null')
  },
  // 删除临时缓存
  remove(key: string) {
      window.sessionStorage.removeItem(key)
  },
  // 清空临时缓存
  clear() {
      window.sessionStorage.clear();
  },
};