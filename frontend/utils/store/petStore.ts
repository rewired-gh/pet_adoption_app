import { defineStore, storeToRefs } from 'pinia'
import { useLoginStore } from './loginStore'

export const usePetStore = defineStore('pet', {
  state: () => ({
    pets: [],
  }),
  actions: {
    async updatePets(isExplore: boolean = false) {
      let req = {}
      let res
      if (isExplore) {
        const loginStore = useLoginStore()
        const { uid } = storeToRefs(loginStore)
        req = {
          Uid: uid.value,
        }
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
