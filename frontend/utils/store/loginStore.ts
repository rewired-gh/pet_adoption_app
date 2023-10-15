import { defineStore } from 'pinia'

export const useLoginStore = defineStore('login', {
  state: () => ({
    isLogin: false,
    uid: 0,
    username: '',
  }),
  actions: {
    updateLogin(isLogin: boolean) {
      this.isLogin = isLogin
    },
    updateUid(uid: number) {
      this.uid = uid
    },
    updateUsername(username: string) {
      this.username = username
    },
  },
  persist: true,
})
