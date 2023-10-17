<template>
  <div class="flex flex-col items-start space-y-4">
    <ElButton size="large" round @click="onOpenDialog()">
      添加宠物信息
    </ElButton>
    <ElTableV2
      class="[&>*]:text-base"
      :width="1000"
      :height="2000"
      :columns="columns"
      :data="petRows"
      :sort-by="sortState"
    >
    </ElTableV2>
    <ElDialog
      append-to-body
      :title="editDialogTitle"
      width="50%"
      v-model="isDialogVisible"
      @closed="clearEditPetData()"
    >
      <div
        class="grid auto-rows-auto grid-cols-3 grid-flow-row gap-x-6 gap-y-6"
      >
        <BaseWithLabel>
          <template #label>
            <span
              class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
              :class="nickname ? 'after:text-transparent' : ''"
            >
              昵称
            </span>
          </template>
          <ElInput v-model="nickname" size="large" />
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span
              class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
              :class="species ? 'after:text-transparent' : ''"
            >
              物种
            </span>
          </template>
          <ElInput v-model="species" size="large" />
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span> 颜色 </span>
          </template>
          <ElInput v-model="color" size="large" />
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span> 性别 </span>
          </template>
          <ElSelectV2
            v-model="gender"
            :options="genderOptions"
            placeholder="未知"
            size="large"
            class="max-w-[7em]"
          />
        </BaseWithLabel>
        <BaseWithLabel>
          <template #label>
            <span> 出生日期 </span>
          </template>
          <ElDatePicker
            v-model="birthday"
            type="dates"
            placeholder="未知"
            size="large"
          />
        </BaseWithLabel>
        <BaseWithLabel class="col-span-full">
          <template #label>
            <span
              class="after:transition after:content-['必填'] after:text-red-300 after:ml-2"
              :class="description ? 'after:text-transparent' : ''"
            >
              描述
            </span>
          </template>
          <ElInput
            type="textarea"
            :rows="8"
            v-model="description"
            size="large"
          />
        </BaseWithLabel>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="onDialogCancel()">取消</el-button>
          <el-button
            type="primary"
            :disabled="!isInfoValid && isEditingPet"
            @click="onEditPet()"
          >
            提交
          </el-button>
        </span>
      </template>
    </ElDialog>
  </div>
</template>

<script setup lang="tsx">
import dayjs from 'dayjs'
import { Column, ElButton, SortBy, TableV2SortOrder } from 'element-plus'
import { storeToRefs } from 'pinia'
import { useLoginStore } from '~/utils/store/loginStore'
import { usePetStore } from '~/utils/store/petStore'

definePageMeta({
  layout: 'console',
})

const columns: Column[] = [
  {
    dataKey: 'nickname',
    title: '昵称',
    width: 200,
    sortable: true,
  },
  {
    dataKey: 'species',
    title: '物种',
    width: 100,
    sortable: true,
  },
  {
    dataKey: 'color',
    title: '颜色',
    width: 100,
    sortable: true,
  },
  {
    dataKey: 'gender',
    title: '性别',
    width: 100,
    sortable: true,
  },
  {
    dataKey: 'description',
    title: '描述',
    width: 300,
  },
  {
    dataKey: 'birthday',
    title: '出生日期',
    width: 200,
    cellRenderer: ({ cellData: birthday }) => (
      <>{dayjs(birthday).format('YYYY/MM/DD')}</>
    ),
    sortable: true,
  },
  {
    dataKey: 'publishDate',
    title: '发布日期',
    width: 200,
    cellRenderer: ({ cellData: publishDate }) => (
      <>{dayjs(publishDate).format('YYYY/MM/DD')}</>
    ),
    sortable: true,
  },
  {
    dataKey: 'isAdopted',
    title: '是否已领养',
    width: 100,
    sortable: true,
  },
  {
    dataKey: 'operation',
    title: '操作',
    width: 200,
    cellRenderer: ({ rowIndex }) => (
      <>
        <ElButton
          type="primary"
          size="small"
          round
          onClick={async () => onEdit(rowIndex)}
        >
          编辑
        </ElButton>
        <ElButton
          type="danger"
          size="small"
          round
          onClick={async () => await onDelete(rowIndex)}
        >
          删除
        </ElButton>
      </>
    ),
  },
]

const sortState = ref<SortBy>({
  key: 'publishDate',
  order: TableV2SortOrder.ASC,
})

const petStore = usePetStore()
onMounted(() => {
  petStore.updatePets()
})
const { pets } = storeToRefs(petStore)
const petRows = computed(() => {
  return pets.value.map((pet, idx) => {
    const genderStr = pet['Gender']
      ? genderOptions.filter((option) => option.value === pet['Gender'])[0]
          .label
      : ''
    return {
      petId: pet['PetID'],
      nickname: pet['Nickname'],
      species: pet['Species'],
      color: pet['Color'],
      gender: genderStr,
      description: pet['Description'],
      publishDate: pet['PublishDate'],
      birthday: pet['Birthday'],
      isAdopted: pet['IsAdopted'] ? '是' : '还没',
    }
  })
})

const isDialogVisible = ref(false)
const isUpdateMode = ref(false)

const onOpenDialog = () => {
  isUpdateMode.value = false
  isDialogVisible.value = true
}
const onDialogCancel = () => {
  isDialogVisible.value = false
}

const nickname = ref()
const species = ref()
const color = ref()
const gender = ref()
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
const birthday = ref()
const description = ref()
let petId = 0

const isInfoValid = computed(() => {
  return nickname.value && species.value && description.value
})

const loginStore = useLoginStore()
const { uid } = storeToRefs(loginStore)

const isEditingPet = ref(false)
const createPet = async () => {
  const req = {
    Nickname: nickname.value,
    Uid: uid.value,
    Description: description.value,
    Birthday: birthday.value ? dayjs(birthday.value).toISOString() : '',
    Species: species.value,
    Color: color.value ? color.value : '',
    Gender: gender.value ? gender.value : '',
  }
  return await fetchR('/pet/new', req)
}
const updatePet = async () => {
  const req = {
    PetID: petId,
    Nickname: nickname.value,
    Description: description.value,
    Birthday: birthday.value ? dayjs(birthday.value).toISOString() : '',
    Species: species.value,
    Color: color.value ? color.value : '',
    Gender: gender.value ? gender.value : '',
  }
  return await fetchR('/pet/update', req)
}
const onEditPet = async () => {
  isEditingPet.value = true
  let res
  if (isUpdateMode.value) {
    res = await updatePet()
  } else {
    res = await createPet()
  }
  isEditingPet.value = false
  if (res) {
    onDialogCancel()
    petStore.updatePets()
  } else {
  }
}

const clearEditPetData = () => {
  nickname.value = null
  species.value = null
  color.value = null
  gender.value = null
  birthday.value = null
  description.value = null
}

const editDialogTitle = computed(() => {
  return isUpdateMode.value ? '编辑宠物信息' : '添加宠物信息'
})

const onEdit = (index: number) => {
  const pet = pets.value[index]
  isUpdateMode.value = true
  petId = pet['PetID']
  nickname.value = pet['Nickname']
  species.value = pet['Species']
  color.value = pet['Color']
  gender.value = pet['Gender']
  birthday.value = [new Date(pet['Birthday'])]
  description.value = pet['Description']
  isDialogVisible.value = true
}
const onDelete = async (index: number) => {
  const petId = petRows.value[index].petId
  const req = {
    PetID: petId,
  }
  const res = await fetchR('/pet/delete', req)
  if (res) {
    petStore.updatePets()
  }
}
</script>
