import { defineStore, storeToRefs } from 'pinia'
import { useLoginStore } from './loginStore'

export const useRoleStore = defineStore('role', {
  state: () => ({
    roles: [] as string[],
  }),
  actions: {
    async updateRole() {
      const loginStore = useLoginStore()
      const { uid } = storeToRefs(loginStore)
      const req = {
        Uid: uid.value,
      }
      const res = await fetchR('/user/get-roles', req)
      if (res) {
        this.roles = res['Roles'] as string[]
      }
    },
  },
  persist: true,
})
