import request from './request';

// 获取任务结果详情
export function getTaskDetail(data: object) {
  return request({
    url: '/plugin/atune/result',
    method: 'get',
    params: data,
  });
}

// 获取调优结果详情
export function getTuneResult(data: object) {
  return request({
    url: '/plugin/atune/tune_result',
    method: 'get',
    params: data,
  });
}