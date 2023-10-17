import { defineStore, storeToRefs } from 'pinia'

export const useUserInfoStore = defineStore('user-info', {
  state: () => ({
    username: '',
    birthday: '',
    gender: '',
    province: '',
    city: '',
    district: '',
    contacts: [],
  }),
  actions: {
    async updateUserInfo(uid: number) {
      let req = {
        Uid: uid,
      }
      const res = await fetchR('/user/get-info', req)
      if (res) {
        this.username = res['Username']
        this.birthday = res['Birthday']
        this.gender = res['Gender']
        this.province = res['Province']
        this.city = res['City']
        this.district = res['District']
        this.contacts = res['Contacts']
      }
    },
  },
})
