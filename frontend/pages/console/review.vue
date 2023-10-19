<template>
  <div>
    <ElTableV2
      class="[&>*]:text-base"
      :width="1000"
      :height="2000"
      :columns="columns"
      :data="adoptionRows"
      :sort-by="sortState"
    >
    </ElTableV2>
  </div>
</template>

<script setup lang="tsx">
import { Column, ElButton, SortBy, TableV2SortOrder, dayjs } from 'element-plus'
import { storeToRefs } from 'pinia'
import { useReviewerAdoptionStore } from '~/utils/store/reviewerAdoptionStore'

definePageMeta({
  layout: 'console',
})

const columns: Column[] = [
  {
    dataKey: 'petId',
    title: '宠物 ID',
    width: 200,
    sortable: true,
  },
  {
    dataKey: 'nickname',
    title: '宠物昵称',
    width: 200,
    sortable: true,
  },
  {
    dataKey: 'species',
    title: '物种',
    width: 200,
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
    dataKey: 'uid',
    title: '用户 ID',
    width: 200,
    sortable: true,
    cellRenderer: ({ cellData: uid }) => (
      <a
        class="after:underline after:underline-offset-4 after:text-slate-400 after:text-xs after:ml-1 after:content-['查看详情']"
        href={`/user/${uid}`}
        target="_blank"
      >
        {uid}
      </a>
    ),
  },
  {
    dataKey: 'username',
    title: '用户名',
    width: 200,
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
          onClick={async () => {
            await onReview(rowIndex, true)
          }}
        >
          批准
        </ElButton>
        <ElButton
          type="danger"
          size="small"
          round
          onClick={async () => {
            await onReview(rowIndex, false)
          }}
        >
          拒绝
        </ElButton>
      </>
    ),
  },
]

const sortState = ref<SortBy>({
  key: 'publishDate',
  order: TableV2SortOrder.ASC,
})

const adoptionStore = useReviewerAdoptionStore()
const { adoptions } = storeToRefs(adoptionStore)
onMounted(() => {
  adoptionStore.updateAdoptions()
})

const adoptionRows = computed(() => {
  return adoptions.value.map((adoption, idx) => {
    return {
      petId: adoption['PetID'],
      nickname: adoption['Nickname'],
      species: adoption['Species'],
      publishDate: adoption['PublishDate'],
      uid: adoption['Uid'],
      username: adoption['Username'],
    }
  })
})

const onReview = async (index: number, isApproved: boolean) => {
  const req = {
    AdoptionID: adoptions.value[index]['AdoptionID'],
    IsApproved: isApproved,
  }
  const res = await fetchR('/adoption/review', req)
  if (res) {
    adoptionStore.updateAdoptions()
  }
}
</script>
