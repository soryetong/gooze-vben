<script setup lang="ts">
import '@wangeditor/editor/dist/css/style.css';
import {
  onBeforeUnmount,
  ref,
  shallowRef,
  watch,
  onMounted,
  nextTick,
} from 'vue';
// @ts-ignore - wangEditor 类型定义问题
import { Editor, Toolbar } from '@wangeditor/editor-for-vue';
import type {
  IDomEditor,
  IEditorConfig,
  IToolbarConfig,
} from '@wangeditor/editor';
import { ElButton } from 'element-plus';
import { Plus } from 'lucide-vue-next';
import MaterialPicker from './MaterialPicker.vue';

interface Props {
  modelValue?: string;
  height?: string;
  placeholder?: string;
  disabled?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: '',
  height: '400px',
  placeholder: '请输入内容...',
  disabled: false,
});

const emit = defineEmits<{
  'update:modelValue': [value: string];
  'update:model-value': [value: string];
  change: [value: string];
}>();

const editorRef = shallowRef<IDomEditor>();
const valueHtml = ref('');

const materialPickerRef = ref<InstanceType<typeof MaterialPicker>>();
const tempImageUrls = ref<string[]>([]);

const openImagePicker = () => {
  materialPickerRef.value?.openPicker();
};

const insertImages = async (urls: string[]) => {
  if (!editorRef.value || urls.length === 0) return;

  const editor = editorRef.value;

  const imagesHtml = urls
    .map(
      (url) =>
        `<p><img src="${url}" alt="图片" style="max-width: 100%;" /></p>`,
    )
    .join('');

  try {
    const currentHtml = editor.getHtml();
    if (
      !currentHtml ||
      currentHtml === '<p><br></p>' ||
      currentHtml.trim() === '' ||
      currentHtml === '<p></p>'
    ) {
      editor.setHtml(imagesHtml);
    } else {
      editor.setHtml(currentHtml + imagesHtml);
    }

    await nextTick();

    const html = editor.getHtml();
    emit('update:modelValue', html);
    emit('update:model-value', html);
    emit('change', html);
  } catch (error) {
    console.error('插入图片失败:', error);
  }
};

watch(tempImageUrls, (newUrls, oldUrls) => {
  if (newUrls.length > 0 && oldUrls.length === 0) {
    insertImages(newUrls);
    setTimeout(() => {
      tempImageUrls.value = [];
    }, 100);
  }
});

// 工具栏配置
const toolbarConfig: Partial<IToolbarConfig> = {
  excludeKeys: ['uploadImage', 'insertImage'], // 排除默认的图片上传和插入按钮
  toolbarKeys: [
    'headerSelect',
    '|',
    'bold',
    'italic',
    'underline',
    'through',
    'clearStyle',
    '|',
    'color',
    'bgColor',
    '|',
    'fontSize',
    'fontFamily',
    'lineHeight',
    '|',
    'bulletedList',
    'numberedList',
    'todo',
    {
      key: 'group-justify',
      title: '对齐',
      iconSvg:
        '<svg viewBox="0 0 1024 1024"><path d="M768 793.6v102.4H51.2v-102.4h716.8z m204.8-230.4v102.4H51.2v-102.4h921.6z m-204.8-230.4v102.4H51.2v-102.4h716.8zM972.8 102.4v102.4H51.2V102.4h921.6z"></path></svg>',
      menuKeys: [
        'justifyLeft',
        'justifyCenter',
        'justifyRight',
        'justifyJustify',
      ],
    },
    {
      key: 'group-indent',
      title: '缩进',
      iconSvg:
        '<svg viewBox="0 0 1024 1024"><path d="M0 64h1024v128H0z m384 192h640v128H384z m0 192h640v128H384z m0 192h640v128H384zM0 832h1024v128H0z m0-128V320l256 192z"></path></svg>',
      menuKeys: ['indent', 'delIndent'],
    },
    '|',
    'emotion',
    'insertLink',
    'insertTable',
    'codeBlock',
    'divider',
    '|',
    'undo',
    'redo',
    '|',
    'fullScreen',
  ],
};

const editorConfig: Partial<IEditorConfig> = {
  placeholder: props.placeholder,
  readOnly: props.disabled,
  MENU_CONF: {},
};

if (editorConfig.MENU_CONF) {
  editorConfig.MENU_CONF['uploadImage'] = {
    server: '',
    maxFileSize: 0,
    allowedFileTypes: [],
  };
}

onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});

const handleCreated = (editor: IDomEditor) => {
  editorRef.value = editor;
};

const handleChange = (editor: IDomEditor) => {
  const html = editor.getHtml();
  emit('update:modelValue', html);
  emit('update:model-value', html);
  emit('change', html);
};

watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal !== valueHtml.value) {
      valueHtml.value = newVal || '';
    }
  },
  { immediate: true },
);

onMounted(() => {
  valueHtml.value = props.modelValue || '';
});
</script>

<template>
  <div class="rich-editor-container">
    <div class="border-border rounded-md border">
      <div
        class="border-border flex items-center justify-between border-b bg-gray-50 px-2 py-1"
      >
        <Toolbar
          class="flex-1"
          :editor="editorRef"
          :defaultConfig="toolbarConfig"
          mode="default"
        />
        <ElButton
          type="primary"
          size="small"
          @click="openImagePicker"
          :disabled="!editorRef || props.disabled"
        >
          <Plus :size="14" class="mr-1" />
          插入图片
        </ElButton>
      </div>
      <Editor
        v-model="valueHtml"
        :style="{ height: height }"
        :defaultConfig="editorConfig"
        mode="default"
        @onCreated="handleCreated"
        @onChange="handleChange"
      />
    </div>

    <div style="display: none">
      <MaterialPicker
        ref="materialPickerRef"
        v-model="tempImageUrls"
        type="image"
        multiple
      />
    </div>
  </div>
</template>

<style scoped>
.rich-editor-container {
  width: 100%;
}

/* 编辑器样式调整 */
:deep(.w-e-text-container) {
  background-color: #fff;
}

:deep(.w-e-toolbar) {
  background-color: #f8f9fa;
}

:deep(.w-e-text-placeholder) {
  color: #999;
  font-style: normal;
}

/* 禁用状态 */
:deep(.w-e-text-container[contenteditable='false']) {
  background-color: #f5f5f5;
  cursor: not-allowed;
}
</style>
