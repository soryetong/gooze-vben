<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { useVbenForm } from '#/adapter/form';
import { createMaterialApi, updateMaterialApi } from '#/api';
import { statusList } from '#/store';
import { useToast, POSITION } from 'vue-toastification';

const toast = useToast();
const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('ui.modal.add', { moduleName: $t('page.material.video.module') })
    : $t('ui.modal.update', { moduleName: $t('page.material.video.module') }),
);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  // 所有表单项共用，可单独在表单内覆盖
  commonConfig: {
    // 所有表单项
    componentProps: {
      class: 'w-full',
    },
  },
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.material.video.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'MaterialUpload',
      fieldName: 'resource',
      label: $t('page.material.video.resource'),
      help: '支持 MP4、WebM 等格式，大小不超过 50MB',
      componentProps: {
        type: 'video',
        limit: 1,
        maxSize: 50,
        sizeUnit: 'MB',
        autoReplace: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'description',
      label: $t('page.material.video.description'),
      defaultValue: '',
      componentProps: {
        type: 'textarea',
        autosize: { minRows: 5, maxRows: 10 },
        rows: 5,
        maxlength: 500,
        showWordLimit: true,
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'RadioGroup',
      fieldName: 'status',
      defaultValue: 1,
      label: $t('ui.table.status'),
      rules: 'selectRequired',
      componentProps: {
        optionType: 'button',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: statusList,
      },
    },
  ],
});

const [Drawer, drawerApi] = useVbenDrawer({
  onCancel() {
    drawerApi.close();
  },

  async onConfirm() {
    // 校验输入的数据
    const validate = await baseFormApi.validate();
    if (!validate.valid) {
      return;
    }

    setLoading(true);

    // 获取表单数据
    const values = await baseFormApi.getValues();

    // 构建提交数据
    const submitData: any = {
      name: values.name,
      resource: values.resource,
      description: values.description,
      status: values.status,
      type: 'video',
    };

    if (values.resource && values.resource.length > 0) {
      const file = values.resource[0];
      if (file instanceof File) {
        submitData.resource = file;
      }
    }

    try {
      await (data.value?.create
        ? createMaterialApi(submitData)
        : updateMaterialApi(data.value.row.id, submitData));

      toast.success(
        data.value?.create
          ? $t('ui.notification.create_success')
          : $t('ui.notification.update_success'),
        {
          timeout: 1000,
          position: POSITION.TOP_RIGHT,
          toastClassName: 'toastification-success',
        },
      );

      drawerApi.close();
      drawerApi.setData({ needRefresh: true });
    } catch {
      // toast.error(
      //   data.value?.create
      //     ? $t('ui.notification.create_failed')
      //     : $t('ui.notification.update_failed'),
      //   {
      //     timeout: 2000,
      //     position: POSITION.TOP_CENTER,
      //   },
      // );
    } finally {
      setLoading(false);
    }
  },

  onOpenChange(isOpen) {
    if (isOpen) {
      // 获取传入的数据
      data.value = drawerApi.getData<Record<string, any>>();

      if (data.value?.row?.meta && data.value?.row?.meta?.authority) {
        const authority = data.value.row.meta.authority;
        data.value.row.meta.authority = authority.join(',');
      }

      // 为表单赋值
      const formData = { ...data.value?.row };

      // 如果是编辑模式且有视频 URL，转换为数组格式供 MaterialUpload 显示
      if (!data.value?.create && formData.resource) {
        formData.resource = [formData.resource]; // 将 URL 字符串包装成数组
      }

      baseFormApi.setValues(formData);

      setLoading(false);
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ loading });
}
</script>

<template>
  <Drawer :title="getTitle">
    <BaseForm />
  </Drawer>
</template>
