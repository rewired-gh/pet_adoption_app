import { defineStore, storeToRefs } from 'pinia'
import { useLoginStore } from './loginStore'

export const usePetStore = defineStore('pet', {
  state: () => ({
    pets: [],
  }),
  actions: {
    async updatePets(isExplore: boolean = false) {
      const loginStore = useLoginStore()
      const { uid } = storeToRefs(loginStore)
      const req = {
        Uid: uid.value,
      }
      let res
      if (isExplore) {
        res = await fetchR('/pet/list-explore', req)
      } else {
        res = await fetchR('/pet/list', req)
      }
      if (res) {
        this.pets = res
      }
    },
  },
})
