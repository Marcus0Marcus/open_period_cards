import request from 'umi-request';
import type { TableListParams } from './data.d';

export async function queryCardTypeList(params?: TableListParams) {
    return request('/api/card-type/list', {
        method: 'POST',
        data: params,
    });
}

export async function removeCardType(params: { key: number[] }) {
    return request('/api/rule', {
        method: 'POST',
        data: {
            ...params,
            method: 'delete',
        },
    });
}

export async function addCardType(params: TableListParams) {
    return request('/api/rule', {
        method: 'POST',
        data: {
            ...params,
            method: 'post',
        },
    });
}

export async function updateCardType(params: TableListParams) {
    return request('/api/rule', {
        method: 'POST',
        data: {
            ...params,
            method: 'update',
        },
    });
}
