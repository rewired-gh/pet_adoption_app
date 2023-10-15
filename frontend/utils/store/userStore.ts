import { defineStore, storeToRefs } from 'pinia'
import { useLoginStore } from './loginStore'

export const useUserStore = defineStore('role', () => {
  const roles = ref([''])
  const loginStore = useLoginStore()
  const { uid } = storeToRefs(loginStore)

  const createData = async () => {
    const req = {
      Uid: uid.value,
    }
    const res = await fetchR('/user/get-roles', req)
    if (res) {
      roles.value = res['Roles'] as string[]
    }
  }

  onMounted(async () => {
    await createData()
  })

  return { roles }
})
