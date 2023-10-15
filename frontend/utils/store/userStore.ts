// import { defineStore } from 'pinia'

// export const useUserStore = defineStore('user', () => {
//   const username = ref('')
//   const gender = ref('')
//   const birthday = ref('')
//   const province = ref('')
//   const city = ref('')
//   const district = ref('')

//   const createData = async () => {

//     let res = null;
//     try {
//       res = await fetch(`${conf.baseApi}/user`, {
//         method: 'GET',
//         body: JSON.stringify(req),
//       })
//     }
//   }

//   onMounted(async () => {
//     await createData()
//   })

//   return {
//     rows,
//     loading,
//     rawData,
//   }
// })
