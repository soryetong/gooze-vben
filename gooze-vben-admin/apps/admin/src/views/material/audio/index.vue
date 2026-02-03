<script lang="ts" setup>
import { h } from 'vue';
import { useVbenVxeGrid, type VxeGridProps } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';
import { ElButton } from 'element-plus';
import MaterialAudioDrawer from './drawer.vue';
import {
  deleteMaterialApi,
  getMaterialListApi,
  updateMaterialApi,
} from '#/api';
import { statusList } from '#/store';
import { useToast, POSITION } from 'vue-toastification';
import { formatDateTime } from '@vben/utils';
import { hasPermission } from '#/directives/permissions';

const toast = useToast();

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  // 控制表单是否显示折叠按钮
  showCollapseButton: false,
  // 按下回车时是否提交表单
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.material.audio.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'status',
      label: $t('ui.table.status'),
      componentProps: {
        options: statusList,
        placeholder: $t('ui.placeholder.select'),
      },
    },
  ],
};

const gridOptions: VxeGridProps = {
  toolbarConfig: {
    custom: true,
    export: true,
    // import: true,
    refresh: true,
    zoom: true,
  },
  height: 'auto',
  exportConfig: {},
  pagerConfig: {},
  rowConfig: {
    isHover: true,
  },
  stripe: true,

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        return await getMaterialListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          keyword: formValues.name,
          status: formValues.status,
          type: 'audio',
        });
      },
    },
  },

  columns: [
    {
      title: $t('ui.table.seq'),
      type: 'seq',
    },
    {
      title: $t('page.material.audio.name'),
      field: 'name',
    },
    {
      title: $t('page.material.audio.resource'),
      field: 'resource',
      slots: { default: 'resource' },
    },
    {
      title: $t('page.material.audio.description'),
      field: 'description',
    },
    {
      title: $t('ui.table.status'),
      field: 'status',
      slots: { default: 'status' },
    },
    {
      title: $t('ui.table.createTime'),
      field: 'createTime',
      slots: { default: 'createdAt' },
    },
    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      width: 120,
    },
  ],
};

const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

async function handleStatusChanged(row: any, checked: boolean) {
  row.pending = true;
  row.status = checked ? 1 : 2;
  row.resource = null;
  row.type = 'audio';
  try {
    await updateMaterialApi(row.id, row);

    toast.success($t('ui.notification.update_success'), {
      timeout: 1000,
      position: POSITION.TOP_RIGHT,
      toastClassName: 'toastification-success',
    });
  } catch {
    // toast.error($t('ui.notification.update_failed'), {
    //   timeout: 2000,
    //   position: POSITION.TOP_CENTER,
    // });
  } finally {
    row.pending = false;
    gridApi.query();
  }
}

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: MaterialAudioDrawer,
  onClosed() {
    const data = drawerApi.getData();
    if (data && data.needRefresh) {
      gridApi.query();
    }
  },
});

function openDrawer(create: boolean, row?: any) {
  drawerApi.setData({
    create,
    row,
  });
  drawerApi.open();
}

/* 创建 */
function handleCreate() {
  openDrawer(true);
}

/* 编辑 */
function handleEdit(row: any) {
  openDrawer(false, row);
}

/* 删除 */
async function handleDelete(row: any) {
  row.pending = true;
  try {
    await deleteMaterialApi(row.id);

    toast.success($t('ui.notification.delete_success'), {
      timeout: 1000,
      position: POSITION.TOP_RIGHT,
      toastClassName: 'toastification-success',
    });
  } catch {
    // toast.error($t('ui.notification.delete_failed'), {
    //   timeout: 2000,
    //   position: POSITION.TOP_CENTER,
    // });
  } finally {
    row.pending = false;
    gridApi.query();
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('page.material.audio.title')">
      <template #toolbar-tools>
        <el-button
          class="mr-2"
          v-permission="['system:material:audio:create']"
          type="primary"
          @click="handleCreate"
        >
          {{ $t('page.material.audio.button.create') }}
        </el-button>
      </template>

      <template #resource="{ row }">
        <audio
          :src="row.resource"
          controls
          class="h-10"
          style="width: 120px"
          preload="metadata"
        />
      </template>

      <template #createdAt="{ row }">
        {{ formatDateTime(row.createTime) }}
      </template>

      <template #status="{ row }">
        <el-switch
          :model-value="row.status === 1"
          :loading="row.pending"
          :inline-prompt="true"
          :active-text="$t('ui.switch.active')"
          :inactive-text="$t('ui.switch.inactive')"
          @change="(checked: boolean) => handleStatusChanged(row, checked)"
          :disabled="!hasPermission(['system:material:audio:update'])"
        />
      </template>

      <template #action="{ row }">
        <ElButton
          type="primary"
          link
          v-permission="['system:material:audio:update']"
          :icon="h(LucideFilePenLine)"
          @click="() => handleEdit(row)"
        />

        <el-popconfirm
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: $t('page.material.audio.module'),
            })
          "
          :confirm-button-text="$t('ui.button.ok')"
          :cancElButton-text="$t('ui.button.cancel')"
          @confirm="() => handleDelete(row)"
        >
          <template #reference>
            <ElButton
              type="danger"
              v-permission="['system:material:audio:delete']"
              link
              :icon="LucideTrash2"
            />
          </template>
        </el-popconfirm>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
<style scoped>
:deep(.vxe-table .vxe-body--row) {
  height: auto !important;
}

:deep(.vxe-table .vxe-body--column) {
  padding: 10px 8px !important;
}

:deep(.vxe-table .vxe-cell),
:deep(.vxe-table .vxe-cell--label),
:deep(.vxe-table .vxe-body--column) {
  height: auto !important;
  max-height: none !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
  line-height: normal !important;
}

:deep(.vxe-table .vxe-body--column:has(.el-image)) {
  overflow: visible !important;
  white-space: normal !important;
}

:deep(.vxe-table .vxe-body--column:has(.el-image) .vxe-cell) {
  overflow: visible !important;
  white-space: normal !important;
}
</style>
