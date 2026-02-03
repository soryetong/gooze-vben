<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import { ElIcon } from 'element-plus';
import { Plus, X, FileAudio } from 'lucide-vue-next';
import { useToast, POSITION } from 'vue-toastification';
import { $t } from '@vben/locales';

const toast = useToast();

interface FileItem {
  file?: File;
  url: string;
  uid: number;
  isExisting?: boolean;
  type?: 'image' | 'audio' | 'video';
}

interface Props {
  modelValue?: File[] | string[];
  /** 文件类型：image/audio/video */
  type?: 'image' | 'audio' | 'video';
  limit?: number;
  maxSize?: number;
  sizeUnit?: 'MB' | 'KB';
  accept?: string;
  autoReplace?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  type: 'image',
  limit: 1,
  maxSize: 10,
  sizeUnit: 'MB',
  autoReplace: false,
});

const emit = defineEmits<{
  'update:modelValue': [value: File[]];
  'update:model-value': [value: File[]];
  input: [value: File[]];
  change: [value: File[]];
}>();

const fileItems = ref<FileItem[]>([]);
const fileInputRef = ref<HTMLInputElement | null>(null);

// 默认 accept 类型
const defaultAccept = computed(() => {
  if (props.accept) return props.accept;
  switch (props.type) {
    case 'image':
      return 'image/jpeg,image/png,image/gif,image/webp';
    case 'audio':
      return 'audio/mpeg,audio/mp3,audio/wav,audio/ogg,audio/aac,audio/m4a';
    case 'video':
      return 'video/mp4,video/webm,video/ogg,video/avi,video/mov';
    default:
      return '*/*';
  }
});

// 文件类型名称
const typeName = computed(() => {
  switch (props.type) {
    case 'image':
      return '图片';
    case 'audio':
      return '音频';
    case 'video':
      return '视频';
    default:
      return '文件';
  }
});

// 是否显示上传按钮
const showUploadButton = computed(() => {
  return fileItems.value.length < props.limit;
});

// 触发文件选择
const handleClick = () => {
  if (fileItems.value.length >= props.limit && !props.autoReplace) {
    toast.warning(`最多只能上传 ${props.limit} 个${typeName.value}`, {
      timeout: 2000,
      position: POSITION.TOP_CENTER,
    });
    return;
  }
  fileInputRef.value?.click();
};

// 文件选择变化
const handleFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement;
  const files = Array.from(input.files || []);

  files.forEach((file) => {
    if (!beforeUpload(file)) {
      return;
    }

    // 自动替换：删除最后一个文件
    if (props.autoReplace && fileItems.value.length >= props.limit) {
      fileItems.value.pop();
    }

    // 达到上限：不添加
    if (fileItems.value.length >= props.limit) {
      toast.warning(`最多只能上传 ${props.limit} 个${typeName.value}`, {
        timeout: 2000,
        position: POSITION.TOP_CENTER,
      });
      return;
    }

    const uid = Date.now() + Math.random();
    fileItems.value.push({
      file,
      url: URL.createObjectURL(file),
      uid,
      type: props.type,
    });
  });

  emitUpdate();
  input.value = '';
};

// 上传前校验
const beforeUpload = (file: File): boolean => {
  // 类型校验
  const acceptTypes = defaultAccept.value.split(',').map((t) => t.trim());
  const fileType = file.type;
  const isValidType = acceptTypes.some((accept) => {
    if (accept === '*/*') return true;
    if (accept.endsWith('/*')) {
      const prefix = accept.split('/')[0];
      return fileType.startsWith(prefix + '/');
    }
    return fileType === accept;
  });

  if (!isValidType) {
    toast.error(
      `${typeName.value}格式不正确！支持格式：${defaultAccept.value}`,
      {
        timeout: 2000,
        position: POSITION.TOP_CENTER,
      },
    );
    return false;
  }

  // 大小校验
  const maxSizeBytes =
    props.maxSize * (props.sizeUnit === 'MB' ? 1024 * 1024 : 1024);
  if (file.size > maxSizeBytes) {
    toast.error(
      `${typeName.value}大小不能超过 ${props.maxSize}${props.sizeUnit}！`,
      {
        timeout: 2000,
        position: POSITION.TOP_CENTER,
      },
    );
    return false;
  }

  return true;
};

// 删除文件
const handleRemove = (uid: number) => {
  const index = fileItems.value.findIndex((item) => item.uid === uid);
  if (index > -1) {
    const item = fileItems.value[index];
    if (item && item.url && !item.isExisting) {
      URL.revokeObjectURL(item.url);
    }
    fileItems.value.splice(index, 1);
    emitUpdate();
  }
};

// 触发更新事件
const emitUpdate = () => {
  const files = fileItems.value
    .filter((item) => !item.isExisting && item.file)
    .map((item) => item.file!);
  emit('update:modelValue', files);
  emit('update:model-value', files);
  emit('input', files);
  emit('change', files);
};

// 监听 modelValue 变化
watch(
  () => props.modelValue,
  (files) => {
    if (!files || files.length === 0) {
      if (fileItems.value.length > 0) {
        fileItems.value = [];
      }
      return;
    }
    const currentFiles = fileItems.value
      .map((item) => item?.file)
      .filter(Boolean);
    if (
      files.length === currentFiles.length &&
      files.every((f, i) => f === currentFiles[i])
    ) {
      return;
    }
    fileItems.value = files.map((file, index) => {
      if (typeof file === 'string') {
        return {
          url: file,
          uid: Date.now() + index,
          isExisting: true,
          type: props.type,
        };
      }
      return {
        file,
        url: URL.createObjectURL(file),
        uid: Date.now() + index,
        isExisting: false,
        type: props.type,
      };
    });
  },
  { immediate: true },
);

defineExpose({
  clear: () => {
    fileItems.value = [];
    emitUpdate();
  },
});
</script>

<template>
  <div class="material-upload">
    <div class="flex flex-wrap gap-2">
      <!-- 已上传的文件列表 -->
      <div
        v-for="item in fileItems"
        :key="item.uid"
        class="upload-item border-border relative overflow-hidden rounded-md border"
      >
        <!-- 图片预览 -->
        <div v-if="type === 'image'" class="upload-preview">
          <img :src="item.url" class="h-full w-full object-cover" />
        </div>

        <!-- 音频预览 -->
        <div
          v-else-if="type === 'audio'"
          class="upload-preview bg-muted flex-col"
        >
          <ElIcon class="text-primary" :size="40">
            <FileAudio />
          </ElIcon>
          <audio
            v-if="item.url"
            :src="item.url"
            controls
            class="mt-2 w-full"
            style="max-width: 200px"
          />
        </div>

        <!-- 视频预览 -->
        <div
          v-else-if="type === 'video'"
          class="upload-preview bg-muted flex-col"
        >
          <video
            v-if="item.url"
            :src="item.url"
            controls
            class="h-full w-full object-cover"
          />
        </div>

        <!-- 删除按钮 -->
        <button
          class="upload-remove"
          type="button"
          @click="() => handleRemove(item.uid)"
        >
          <ElIcon :size="16">
            <X />
          </ElIcon>
        </button>
      </div>

      <!-- 上传按钮 -->
      <div v-if="showUploadButton" class="upload-trigger" @click="handleClick">
        <ElIcon :size="28" class="text-muted-foreground">
          <Plus />
        </ElIcon>
        <span class="upload-text">{{ $t('ui.button.upload') }}</span>
      </div>
    </div>

    <!-- 隐藏的文件输入框 -->
    <input
      ref="fileInputRef"
      type="file"
      :accept="defaultAccept"
      :multiple="limit > 1"
      style="display: none"
      @change="handleFileChange"
    />
  </div>
</template>

<style scoped>
.material-upload {
  width: 100%;
}

.upload-item {
  position: relative;
  width: 180px;
  height: 180px;
  transition: all 0.3s;
  background: var(--el-bg-color-overlay);
}

.upload-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.12);
  transform: translateY(-2px);
}

.upload-preview {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  background: var(--el-fill-color-light);
}

.upload-remove {
  position: absolute;
  top: 8px;
  right: 8px;
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0, 0, 0, 0.6);
  border: none;
  border-radius: 50%;
  cursor: pointer;
  color: white;
  transition: all 0.2s;
  z-index: 10;
  opacity: 0;
}

.upload-item:hover .upload-remove {
  opacity: 1;
}

.upload-remove:hover {
  background: rgba(0, 0, 0, 0.8);
  transform: scale(1.1);
}

.upload-trigger {
  width: 180px;
  height: 180px;
  border: 2px dashed var(--el-border-color);
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s;
  background: var(--el-fill-color-blank);
}

.upload-trigger:hover {
  border-color: var(--el-color-primary);
  background: var(--el-color-primary-light-9);
}

.upload-trigger:hover .upload-text {
  color: var(--el-color-primary);
}

.upload-text {
  margin-top: 8px;
  font-size: 14px;
  color: var(--el-text-color-secondary);
  transition: color 0.3s;
}
</style>
