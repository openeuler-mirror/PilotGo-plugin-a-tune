import request from './request';

// 获取任务结果详情
export function getTaskDetail(data: object) {
  return request({
    url: '/plugin/atune/result',
    method: 'get',
    params: data,
  });
}
