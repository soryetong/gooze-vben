<script setup lang="ts">
import { ref, computed, defineAsyncComponent, watch } from 'vue';
import {
  ElDialog,
  ElButton,
  ElImage,
  ElEmpty,
  ElInput,
  ElPagination,
} from 'element-plus';
import { Plus, Check, Search, X } from 'lucide-vue-next';
import { getMaterialListApi } from '#/api';
import { useVbenDrawer } from '@vben/common-ui';

interface Props {
  modelValue?: string | string[];
  type?: 'image' | 'audio' | 'video';
  multiple?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  type: 'image',
  multiple: false,
});

const emit = defineEmits<{
  'update:modelValue': [value: string | string[]];
  'update:model-value': [value: string | string[]];
  change: [value: string | string[]];
}>();

const dialogVisible = ref(false);
const materialList = ref<any[]>([]);
const loading = ref(false);
const selectedUrl = ref<string>('');
const selectedUrls = ref<string[]>([]);

const currentPage = ref(1);
const pageSize = ref(12);
const total = ref(0);
const searchKeyword = ref('');

let searchTimer: NodeJS.Timeout | null = null;

watch(searchKeyword, () => {
  if (searchTimer) {
    clearTimeout(searchTimer);
  }
  searchTimer = setTimeout(() => {
    handleSearch();
  }, 500);
});

const typeName = computed(() => {
  switch (props.type) {
    case 'image':
      return '图片';
    case 'audio':
      return '音频';
    case 'video':
      return '视频';
    default:
      return '素材';
  }
});

const AsyncImageDrawer = defineAsyncComponent(
  () => import('#/views/material/image/drawer.vue'),
);
const AsyncAudioDrawer = defineAsyncComponent(
  () => import('#/views/material/audio/drawer.vue'),
);
const AsyncVideoDrawer = defineAsyncComponent(
  () => import('#/views/material/video/drawer.vue'),
);

const [ImageDrawer, imageDrawerApi] = useVbenDrawer({
  connectedComponent: AsyncImageDrawer,
  zIndex: 3000,
  onClosed() {
    const data = imageDrawerApi.getData();
    if (data?.needRefresh) {
      currentPage.value = 1;
      loadMaterials();
    }
  },
});

const [AudioDrawer, audioDrawerApi] = useVbenDrawer({
  connectedComponent: AsyncAudioDrawer,
  zIndex: 3000,
  onClosed() {
    const data = audioDrawerApi.getData();
    if (data?.needRefresh) {
      currentPage.value = 1;
      loadMaterials();
    }
  },
});

const [VideoDrawer, videoDrawerApi] = useVbenDrawer({
  connectedComponent: AsyncVideoDrawer,
  zIndex: 3000,
  onClosed() {
    const data = videoDrawerApi.getData();
    if (data?.needRefresh) {
      currentPage.value = 1;
      loadMaterials();
    }
  },
});

const loadMaterials = async () => {
  loading.value = true;
  try {
    const result = await getMaterialListApi({
      page: currentPage.value,
      pageSize: pageSize.value,
      keyword: searchKeyword.value,
      type: props.type,
      status: 1, // 只显示正常的素材
    });
    materialList.value = result.items || [];
    total.value = result.total || 0;
  } catch (error) {
    console.error('加载素材失败:', error);
  } finally {
    loading.value = false;
  }
};

const handleSearch = () => {
  currentPage.value = 1;
  loadMaterials();
};

const handlePageChange = (page: number) => {
  currentPage.value = page;
  loadMaterials();
};

const openPicker = () => {
  dialogVisible.value = true;
  if (props.multiple) {
    selectedUrls.value = Array.isArray(props.modelValue)
      ? [...props.modelValue]
      : [];
  } else {
    selectedUrl.value = (props.modelValue as string) || '';
  }
  searchKeyword.value = '';
  currentPage.value = 1;
  loadMaterials();
};

const selectMaterial = (url: string) => {
  if (props.multiple) {
    const index = selectedUrls.value.indexOf(url);
    if (index > -1) {
      selectedUrls.value.splice(index, 1);
    } else {
      selectedUrls.value.push(url);
    }
  } else {
    selectedUrl.value = url;
  }
};

const confirmSelect = () => {
  if (props.multiple) {
    emit('update:modelValue', selectedUrls.value);
    emit('update:model-value', selectedUrls.value);
    emit('change', selectedUrls.value);
  } else {
    emit('update:modelValue', selectedUrl.value);
    emit('update:model-value', selectedUrl.value);
    emit('change', selectedUrl.value);
  }
  dialogVisible.value = false;
};

const openUploadDrawer = () => {
  const drawerData = { create: true };

  switch (props.type) {
    case 'image':
      imageDrawerApi.setData(drawerData);
      imageDrawerApi.open();
      break;
    case 'audio':
      audioDrawerApi.setData(drawerData);
      audioDrawerApi.open();
      break;
    case 'video':
      videoDrawerApi.setData(drawerData);
      videoDrawerApi.open();
      break;
  }
};

const removeImage = (index: number) => {
  if (Array.isArray(props.modelValue)) {
    const newUrls = [...props.modelValue];
    newUrls.splice(index, 1);
    emit('update:modelValue', newUrls);
    emit('update:model-value', newUrls);
    emit('change', newUrls);
  }
};

defineExpose({
  openPicker,
});
</script>

<template>
  <div class="w-full">
    <!-- 单选预览 -->
    <div
      v-if="!multiple"
      class="border-border hover:border-primary relative h-[150px] w-[150px] cursor-pointer overflow-hidden rounded-lg border-2 border-dashed transition-all"
      @click="openPicker"
    >
      <div v-if="modelValue" class="relative h-full w-full">
        <div
          v-if="type === 'image'"
          class="bg-muted flex h-full w-full items-center justify-center"
        >
          <el-image
            :src="modelValue as string"
            fit="contain"
            class="h-full w-full object-contain"
          />
        </div>
        <div
          v-else-if="type === 'audio'"
          class="bg-muted flex h-full w-full items-center justify-center p-5"
        >
          <audio :src="modelValue as string" controls class="w-full" />
        </div>
        <video
          v-else-if="type === 'video'"
          :src="modelValue as string"
          class="h-full w-full object-cover"
        />
        <div
          class="absolute inset-0 flex items-center justify-center bg-black/50 opacity-0 transition-opacity hover:opacity-100"
        >
          <span class="text-white">点击更换</span>
        </div>
      </div>
      <div
        v-else
        class="bg-background flex h-full w-full flex-col items-center justify-center"
      >
        <Plus :size="30" class="text-muted-foreground" />
        <span class="text-muted-foreground mt-2">选择{{ typeName }}</span>
      </div>
    </div>

    <!-- 多选预览 -->
    <div v-else class="flex flex-wrap gap-2">
      <div
        v-for="(url, index) in modelValue as string[]"
        :key="index"
        class="border-border group relative h-[150px] w-[150px] overflow-hidden rounded-lg border-2 border-dashed"
      >
        <div
          v-if="type === 'image'"
          class="bg-muted flex h-full w-full items-center justify-center"
        >
          <el-image
            :src="url"
            fit="contain"
            class="h-full w-full object-contain"
          />
        </div>
        <!-- 删除按钮 -->
        <div
          class="absolute right-1 top-1 flex h-7 w-7 cursor-pointer items-center justify-center rounded-full bg-red-500 text-white opacity-0 transition-opacity hover:bg-red-600 group-hover:opacity-100"
          @click="removeImage(index)"
        >
          <X :size="16" />
        </div>
      </div>

      <!-- 添加按钮 -->
      <div
        class="border-border hover:border-primary flex h-[150px] w-[150px] cursor-pointer items-center justify-center rounded-lg border-2 border-dashed transition-all"
        @click="openPicker"
      >
        <div class="flex flex-col items-center justify-center">
          <Plus :size="30" class="text-muted-foreground" />
          <span class="text-muted-foreground mt-2">选择{{ typeName }}</span>
        </div>
      </div>
    </div>

    <!-- 选择弹窗 -->
    <ElDialog
      v-model="dialogVisible"
      :title="`选择${typeName}${multiple ? '（可多选）' : ''}`"
      width="800px"
      :close-on-click-modal="false"
      :append-to-body="true"
      :lock-scroll="true"
      top="5vh"
    >
      <div class="mb-4 space-y-3">
        <div class="bg-muted flex items-center justify-between rounded-lg p-3">
          <span class="text-muted-foreground text-sm">
            没有合适的素材？点击按钮上传新素材
          </span>
          <ElButton type="primary" size="default" @click="openUploadDrawer">
            <Plus :size="16" class="mr-1" />
            上传新素材
          </ElButton>
        </div>

        <div class="flex items-center gap-3">
          <ElInput
            v-model="searchKeyword"
            placeholder="搜索素材名称，输入自动搜索..."
            clearable
            size="large"
            class="flex-1"
          >
            <template #prefix>
              <Search :size="18" class="text-muted-foreground" />
            </template>
          </ElInput>
        </div>
      </div>

      <div
        v-if="materialList.length > 0"
        class="grid max-h-[55vh] grid-cols-4 gap-4 overflow-y-auto p-2"
        v-loading="loading"
      >
        <div
          v-for="item in materialList"
          :key="item.id"
          class="border-border hover:border-primary relative cursor-pointer overflow-hidden rounded-lg border-2 transition-all hover:shadow-md"
          :class="{
            'border-primary shadow-[0_0_0_2px_var(--el-color-primary-light-7)]':
              multiple
                ? selectedUrls.includes(item.resource)
                : selectedUrl === item.resource,
          }"
          @click="selectMaterial(item.resource)"
        >
          <div
            v-if="type === 'image'"
            class="bg-muted flex h-[150px] w-full items-center justify-center"
          >
            <el-image
              :src="item.resource"
              fit="contain"
              class="h-full w-full object-contain"
            />
          </div>
          <div
            v-else-if="type === 'audio'"
            class="bg-muted flex h-[150px] w-full items-center justify-center p-2"
          >
            <audio :src="item.resource" controls class="w-full" />
          </div>
          <video
            v-else-if="type === 'video'"
            :src="item.resource"
            class="bg-muted h-[150px] w-full object-cover"
          />
          <div class="bg-white p-2">
            <div
              class="text-foreground overflow-hidden text-ellipsis whitespace-nowrap text-sm"
            >
              {{ item.name }}
            </div>
          </div>
          <div
            v-if="
              multiple
                ? selectedUrls.includes(item.resource)
                : selectedUrl === item.resource
            "
            class="bg-primary absolute right-2 top-2 flex h-8 w-8 items-center justify-center rounded-full text-white"
          >
            <Check :size="20" />
          </div>
        </div>
      </div>

      <div
        v-else-if="!loading"
        class="flex h-[55vh] items-center justify-center"
      >
        <ElEmpty description="暂无素材，请先上传">
          <ElButton type="primary" size="default" @click="openUploadDrawer">
            <Plus :size="16" class="mr-1" />
            上传素材
          </ElButton>
        </ElEmpty>
      </div>

      <div v-else class="flex h-[55vh] items-center justify-center">
        <div v-loading="true" class="h-20 w-20"></div>
      </div>

      <div v-if="total > 0" class="mt-4 flex justify-center">
        <ElPagination
          :current-page="currentPage"
          :page-size="pageSize"
          :total="total"
          layout="total, prev, pager, next"
          @current-change="handlePageChange"
        />
      </div>

      <template #footer>
        <div class="flex items-center justify-between">
          <span v-if="multiple" class="text-muted-foreground text-sm">
            已选择 {{ selectedUrls.length }} 个{{ typeName }}
          </span>
          <span v-else></span>
          <div class="flex gap-2">
            <ElButton @click="dialogVisible = false">取消</ElButton>
            <ElButton
              type="primary"
              @click="confirmSelect"
              :disabled="multiple ? selectedUrls.length === 0 : !selectedUrl"
            >
              确定
            </ElButton>
          </div>
        </div>
      </template>
    </ElDialog>

    <ImageDrawer v-if="type === 'image'" />
    <AudioDrawer v-if="type === 'audio'" />
    <VideoDrawer v-if="type === 'video'" />
  </div>
</template>
