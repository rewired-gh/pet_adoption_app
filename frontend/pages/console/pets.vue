<template>
  <div class="flex flex-col items-start space-y-4">
    <ElButton size="large" round @click="onOpenDialog()">
      添加宠物信息
    </ElButton>
    <ElTableV2
      class="[&>*]:text-base"
      :width="1100"
      :height="80"
      :columns="columns"
      :data="data"
    >
    </ElTableV2>
  </div>
  <ElDialog
    append-to-body
    title="添加宠物信息"
    width="50%"
    v-model="isDialogVisible"
  >
    <div class="grid auto-rows-auto grid-cols-3 grid-flow-row gap-x-6 gap-y-6">
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
          class="w-44"
        />
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span> 生日 </span>
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
        <ElInput type="textarea" :rows="8" v-model="description" size="large" />
      </BaseWithLabel>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="onDialogCancel()">取消</el-button>
        <el-button
          type="primary"
          :disabled="!isInfoValid"
          @click="onDialogCancel"
        >
          添加
        </el-button>
      </span>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'console',
})

const columns = [
  {
    dataKey: 'nickname',
    title: '昵称',
    width: 200,
  },
  {
    dataKey: 'species',
    title: '物种',
    width: 200,
  },
  {
    dataKey: 'color',
    title: '颜色',
    width: 200,
  },
  {
    dataKey: 'gender',
    title: '性别',
    width: 200,
  },
  {
    dataKey: 'description',
    title: '描述',
    width: 200,
  },
  {
    dataKey: 'publishDate',
    title: '发布日期',
    width: 200,
  },
  {
    dataKey: 'isAdopted',
    title: '是否已领养',
    width: 200,
  },
  {
    dataKey: 'isAdopted',
    title: '操作',
    width: 400,
  },
]

const data = []

const isDialogVisible = ref(false)

const onOpenDialog = () => {
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

const isInfoValid = computed(() => {
  return nickname.value && species.value && description.value
})
</script>
