import { defineStore, storeToRefs } from 'pinia'
import { useLoginStore } from './loginStore'

export const useUserAdoptionStore = defineStore('user-adoption', {
  state: () => ({
    adoptions: [],
  }),
  actions: {
    async updateAdoptions() {
      const loginStore = useLoginStore()
      const { uid } = storeToRefs(loginStore)

      const req = {
        Uid: uid.value,
      }
      const res = await fetchR('/adoption/list-user', req)
      if (res) {
        this.adoptions = res
      }
    },
  },
})
