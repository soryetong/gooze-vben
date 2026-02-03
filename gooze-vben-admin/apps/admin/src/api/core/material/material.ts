import { requestClient } from '#/api/request';

/**
 * 获取素材列表
 * @param params 查询参数
 * @return 素材列表
 */
export const getMaterialListApi = async (params: any) => {
  return requestClient.getWithParams('/material/list', params);
};

/**
 * 创建素材
 * @param data 素材数据
 * @return 创建结果
 */
export const createMaterialApi = async (data: any) => {
  return requestClient.upload('/material/add', data);
};

/**
 * 更新素材
 * @param id 素材ID
 * @param data 素材数据
 * @return 更新结果
 */
export const updateMaterialApi = async (id: number, data: any) => {
  return requestClient.upload(`/material/update/${id}`, data);
};

/**
 * 删除素材
 * @param id 素材ID
 * @return 删除结果
 */

export const deleteMaterialApi = async (id: number) => {
  return requestClient.delete(`/material/delete/${id}`);
};
