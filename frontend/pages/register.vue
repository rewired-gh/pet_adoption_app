<template>
  <div
    class="m-auto first-letter:flex flex-col space-y-8 max-w-screen-md pt-10"
  >
    <BaseSectionHeading>
      让我们开始注册吧
      <template #icon>
        <Icon name="fluent:signature-16-regular" />
      </template>
    </BaseSectionHeading>
    <section
      class="grid auto-rows-auto grid-cols-2 grid-flow-row gap-x-6 gap-y-6"
    >
      <BaseWithLabel>
        <template #label>
          <span
            class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
            :class="username ? 'after:text-transparent' : ''"
          >
            用户名
          </span>
        </template>
        <ElInput v-model="username" size="large" />
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span
            class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
            :class="district ? 'after:text-transparent' : ''"
          >
            地区
          </span>
        </template>
        <div class="flex space-x-1">
          <ElSelectV2
            v-model="province"
            :options="provinceOptions"
            placeholder="省级"
            size="large"
          />
          <ElSelectV2
            v-model="city"
            :options="cityOptions"
            placeholder="市级"
            size="large"
          />
          <ElSelectV2
            v-model="district"
            :options="districtOptions"
            placeholder="区级"
            size="large"
          />
        </div>
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span
            class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
            :class="pword ? 'after:text-transparent' : ''"
          >
            密码
          </span>
        </template>
        <ElInput v-model="pword" type="password" size="large" />
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span
            class="after:transition after:content-['两次输入的密码不一致'] after:text-red-300 after:ml-2"
            :class="getPwordMismatchStyle"
          >
            再次输入密码
          </span>
        </template>
        <ElInput v-model="pwordAgain" type="password" size="large" />
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span
            class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
            :class="
              contacts[0].content && contacts[0].kind
                ? 'after:text-transparent'
                : ''
            "
          >
            联系方式 1
          </span>
        </template>
        <div class="flex space-x-1">
          <ElSelectV2
            v-model="contacts[0].kind"
            :options="contactKindOptions"
            placeholder="类型"
            size="large"
            class="w-28"
          />
          <ElInput v-model="contacts[0].content" size="large" />
        </div>
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span> 联系方式 2 </span>
        </template>
        <div class="flex space-x-1">
          <ElSelectV2
            v-model="contacts[1].kind"
            :options="contactKindOptions"
            placeholder="类型"
            size="large"
            class="w-28"
          />
          <ElInput v-model="contacts[1].content" size="large" />
        </div>
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span
            class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
            :class="birthday ? 'after:text-transparent' : ''"
          >
            出生日期
          </span>
        </template>
        <ElDatePicker
          v-model="birthday"
          type="dates"
          placeholder="请选择"
          size="large"
        />
      </BaseWithLabel>
      <div class="flex space-x-6 items-end">
        <BaseWithLabel>
          <template #label>
            <span
              class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
              :class="gender ? 'after:text-transparent' : ''"
            >
              性别
            </span>
          </template>
          <ElSelectV2
            v-model="gender"
            placeholder="请选择"
            :options="genderOptions"
            size="large"
            class="max-w-[7em]"
          />
        </BaseWithLabel>
        <BaseWithLabel class="grow">
          <template #label></template>
          <ElButton
            type="primary"
            plain
            class="w-full"
            size="large"
            round
            :disabled="!isInfoValid || isRegistering"
            :loading="isRegistering"
            @click="onRegister()"
            >开始注册
          </ElButton>
        </BaseWithLabel>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { conf } from '~/utils/config'
import { useLoginStore } from '~/utils/store/loginStore'

const dummyLocations = [
  {
    locationId: 1,
    province: '北京市',
    city: '',
    district: '朝阳区',
  },
  {
    locationId: 2,
    province: '广东省',
    city: '深圳市',
    district: '南山区',
  },
]

const username = ref('')
const pword = ref('')
const pwordAgain = ref('')

const getPwordMismatchStyle = computed(() => {
  return pword.value !== pwordAgain.value ? '' : 'after:text-transparent'
})

const valToOption = (val: string) => {
  const option = { value: val, label: val }
  return option
}

const province = ref()
const provinceOptions = computed(() => {
  return dummyLocations.map((location) => location.province).map(valToOption)
})

const city = ref()
const cityOptions = computed(() => {
  return dummyLocations
    .filter((location) => location.province === province.value && location.city)
    .map((location) => location.city)
    .map(valToOption)
})

const district = ref()
const districtOptions = computed(() => {
  return dummyLocations
    .filter(
      (location) =>
        location.province === province.value &&
        (location.city === city.value || !location.city)
    )
    .map((location) => location.district)
    .map(valToOption)
})

const locationId = computed(() => {
  const location = dummyLocations.filter((location) => {
    return (
      province.value === location.province &&
      district.value === location.district &&
      (city.value === location.city || (!city.value && !location.city))
    )
  })
  if (location.length > 0) {
    return location[0].locationId
  }
  return 0
})

watch(province, () => {
  city.value = null
  district.value = null
})
watch(city, () => {
  district.value = null
})

const gender = ref()
const genderOptions = [
  {
    label: '女',
    value: 'female',
  },
  {
    label: '男',
    value: 'male',
  },
  {
    label: '不透露或其他',
    value: 'other',
  },
]

const birthday = ref()

const isInfoValid = computed(() => {
  return (
    username.value &&
    pword.value &&
    pword.value === pwordAgain.value &&
    district &&
    gender.value &&
    birthday.value &&
    contacts[0].content &&
    contacts[0].kind
  )
})

const isRegistering = ref(false)
const store = useLoginStore()

const onRegister = async () => {
  isRegistering.value = true
  let isErr = false
  const req = {
    LocationID: locationId.value,
    Username: username.value,
    Pword: pword.value,
    Gender: gender.value,
    Birthday: (birthday.value[0] as Date).toISOString(),
  }
  const res = await fetchR('/user/new', req)
  if (res) {
    const uid: number = res['Uid']
    const contactReq = {
      Uid: uid,
      Contacts: contacts
        .filter((contact) => contact.kind && contact.content)
        .map((contact) => ({
          Kind: contact.kind,
          Content: contact.content,
        })),
    }
    const contactRes = await fetchR('/user/add-contacts', contactReq)
    if (contactRes) {
      store.updateLogin(true)
      store.updateUid(uid)
      store.updateUsername(req.Username)
      navigateTo('/explore')
    } else {
      isErr = true
    }
  } else {
    isErr = true
  }
  isRegistering.value = false
}

const contactKindOptions = [
  {
    label: '邮箱',
    value: 'email',
  },
  {
    label: '电话',
    value: 'phone',
  },
]
type Contact = {
  kind: string | null
  content: string
}
const contacts: Contact[] = reactive(
  Array(2)
    .fill(null)
    .map(() => ({
      kind: null,
      content: '',
    }))
)
</script>
