<template>
  <div class="flex flex-col max-w-screen-sm w-full space-y-12">
    <BaseSectionHeading>
      用户档案
      <template #icon>
        <Icon name="fluent:notepad-person-24-regular" />
      </template>
    </BaseSectionHeading>
    <div class="flex flex-col gap-y-8">
      <div class="flex flex-wrap gap-x-28 gap-y-8">
        <BaseWithLabel>
          <template #label>
            <span> UID </span>
          </template>
          <span class="text-lg ml-2"> {{ otherUid }} </span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span> 用户名 </span>
          </template>
          <span class="text-lg ml-2"> {{ username }} </span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span> 年龄 </span>
          </template>
          <span class="text-lg ml-2"> {{ age }} </span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span> 性别 </span>
          </template>
          <span class="text-lg ml-2"> {{ genderStr }} </span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span> 地址 </span>
          </template>
          <span class="text-lg ml-2"> {{ locationStr }} </span>
        </BaseWithLabel>
      </div>
      <BaseWithLabel>
        <template #label>
          <span> 联系方式 </span>
        </template>
        <ul class="text-lg ml-2">
          <li v-for="contact in parsedContacts">
            {{ contact.kind }}: {{ contact.content }}
          </li>
        </ul>
      </BaseWithLabel>
    </div>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { storeToRefs } from 'pinia'
import { useUserInfoStore } from '~/utils/store/userInfoStore'

const route = useRoute()
const otherUid = parseInt(route.params.id as string)

const userInfoStore = useUserInfoStore()
onMounted(() => {
  userInfoStore.updateUserInfo(otherUid)
})
const { username, birthday, gender, province, city, district, contacts } =
  storeToRefs(userInfoStore)

const age = computed(() => {
  return dayjs().diff(birthday.value, 'year')
})
const genderMap: { [key: string]: string } = {
  male: '男',
  female: '女',
  other: '不透露或其他',
}
const contactKindMap: { [key: string]: string } = {
  email: '邮箱',
  phone: '电话',
}
const genderStr = computed(() => {
  return genderMap[gender.value]
})
const locationStr = computed(() => {
  return `${province.value} ${city.value} ${district.value}`
})
const parsedContacts = computed(() => {
  return contacts.value.map((contact) => {
    return {
      kind: contactKindMap[contact['Kind']],
      content: contact['Content'],
    }
  })
})
</script>
