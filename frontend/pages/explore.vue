<template>
  <div class="flex max-w-screen-xl w-full space-x-12">
    <div class="w-0 h-0">
      <div class="text-pink-100" />
      <div class="text-green-100" />
      <div class="text-lime-100" />
      <div class="text-violet-100" />
    </div>
    <section class="flex flex-col space-y-6 w-2/3">
      <BaseSectionHeading>
        发现宠物
        <template #icon>
          <Icon name="fluent:animal-cat-24-regular" />
        </template>
      </BaseSectionHeading>
      <ul class="flex flex-wrap gap-x-8 gap-y-8">
        <li
          v-for="(pet, idx) in pets"
          :key="idx"
          class="transition flex flex-col relative overflow-hidden shadow-md p-6 rounded-2xl w-[16em] space-y-4 hover:shadow-lg hover:-translate-y-1 hover:cursor-pointer"
          @click="onViewPetDetail(idx)"
        >
          <h3 class="text-xl font-medium whitespace-nowrap">
            {{ pet['Nickname'] }}
          </h3>
          <p class="overflow-clip line-clamp-3 h-[4.5em]">
            {{ pet['Description'] }}
          </p>
          <div class="text-slate-400 space-x-3">
            <span>#{{ pet['Species'] }}</span>
          </div>
          <Icon
            name="mdi:paw"
            class="right-0 -bottom-3 -rotate-[24deg] absolute text-7xl"
            :class="getRandomColorStyle('text', '100')"
          />
        </li>
      </ul>
    </section>
    <section class="flex flex-col space-y-6">
      <BaseSectionHeading>
        领养申请状态
        <template #icon>
          <Icon name="fluent:apps-list-detail-24-regular" />
        </template>
      </BaseSectionHeading>
      <ul class="flex flex-col gap-y-8">
        <li
          v-for="adoption in adoptionViewList"
          class="w-[15em] flex flex-col space-y-4 rounded-2xl border p-4"
        >
          <h3 class="text-xl font-medium whitespace-nowrap">
            {{ adoption.nickname }}
          </h3>
          <div class="flex justify-between">
            <div class="text-slate-400 space-x-3">
              <span> #{{ adoption.species }} </span>
              <span> #ID: {{ adoption.petId }} </span>
            </div>
            <ElTag :type="getTagType(adoption.status)" effect="plain" round>
              {{ adoption.status }}
            </ElTag>
          </div>
        </li>
      </ul>
    </section>
    <ElDialog
      append-to-body
      title="宠物档案"
      width="50%"
      v-model="isDialogVisible"
    >
      <div
        class="grid auto-rows-auto grid-cols-3 grid-flow-row gap-x-6 gap-y-4"
      >
        <BaseWithLabel>
          <template #label>
            <span>昵称</span>
          </template>
          <span class="text-lg ml-2">{{ selectedPet.nickname }}</span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span>物种</span>
          </template>
          <span class="text-lg ml-2">{{ selectedPet.species }}</span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span>颜色</span>
          </template>
          <span class="text-lg ml-2">{{ selectedPet.color }}</span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span>性别</span>
          </template>
          <span class="text-lg ml-2">{{ selectedPet.gender }}</span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span>出生日期</span>
          </template>
          <span class="text-lg ml-2">{{ selectedPet.birthday }}</span>
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span>发布日期</span>
          </template>
          <span class="text-lg ml-2">{{ selectedPet.publishDate }}</span>
        </BaseWithLabel>
        <BaseWithLabel class="col-span-full">
          <template #label>
            <span>描述</span>
          </template>
          <p class="text-lg ml-2">{{ selectedPet.description }}</p>
        </BaseWithLabel>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="onDialogCancel()">关闭</el-button>
          <el-button type="primary" @click="onAdoptPet()" :disable="isAdopting">
            申请领养
          </el-button>
        </span>
      </template>
    </ElDialog>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { storeToRefs } from 'pinia'
import { useLoginStore } from '~/utils/store/loginStore'
import { usePetStore } from '~/utils/store/petStore'
import { useUserAdoptionStore } from '~/utils/store/userAdoptionStore'

const loginStore = useLoginStore()
const { isLogin, uid } = storeToRefs(loginStore)

if (!isLogin.value) {
  navigateTo('/')
}

const colors = ['green', 'pink', 'violet', 'lime']
const getRandomColorStyle = (prefix: string, suffix: string) => {
  const index = Math.floor(Math.random() * colors.length)
  return `${prefix}-${colors[index]}-${suffix}`
}

const petStore = usePetStore()
const { pets } = storeToRefs(petStore)

const adoptionStore = useUserAdoptionStore()
const { adoptions } = storeToRefs(adoptionStore)

onMounted(() => {
  petStore.updatePets(true)
  adoptionStore.updateAdoptions()
})

const isAdopting = ref(false)
const isDialogVisible = ref(false)
const onDialogCancel = () => {
  isDialogVisible.value = false
}
const onAdoptPet = async () => {
  isAdopting.value = true
  const req = {
    Uid: uid.value,
    PetID: pets.value[petIndex.value]['PetID'],
  }
  const res = await fetchR('/adoption/new', req)
  isAdopting.value = false
  if (res) {
    onDialogCancel()
    adoptionStore.updateAdoptions()
    petStore.updatePets(true)
  }
}

const petIndex = ref(0)
const onViewPetDetail = (index: number) => {
  petIndex.value = index
  isDialogVisible.value = true
}
const genderOptions = [
  {
    label: '雌',
    value: 'female',
  },
  {
    label: '雄',
    value: 'male',
  },
  {
    label: '其他',
    value: 'other',
  },
]
const selectedPet = computed(() => {
  const pet = pets.value[petIndex.value]
  return {
    nickname: pet['Nickname'],
    species: pet['Species'],
    color: pet['Color'] ? pet['Color'] : '未知',
    gender: pet['Gender']
      ? genderOptions.filter((option) => option.value === pet['Gender'])[0]
          .label
      : '未知',
    birthday: pet['Birthday']
      ? dayjs(pet['Birthday']).format('YYYY/MM/DD')
      : '未知',
    publishDate: dayjs(pet['PublishDate']).format('YYYY/MM/DD'),
    description: pet['Description'],
  }
})

const adoptionViewList = computed(() => {
  return adoptions.value.map((adoption) => {
    const status = adoption['IsReviewed']
      ? adoption['IsApproved']
        ? '已通过'
        : '未通过'
      : '待审核'
    return {
      petId: adoption['PetID'],
      nickname: adoption['Nickname'],
      species: adoption['Species'],
      status: status,
    }
  })
})

const getTagType = (status: string) => {
  switch (status) {
    case '已通过':
      return 'success'
    case '未通过':
      return 'danger'
    case '待审核':
      return 'info'
  }
}
</script>
