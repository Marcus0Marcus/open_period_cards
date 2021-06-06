import request from '@/utils/request';

export type LoginParamsType = {
    name: string;
    password: string;
    mobile: string;
    captcha: string;
};

export async function AccountLogin(params: LoginParamsType) {
    return request('login', {
        method: 'POST',
        data: params,
    });
}

export async function getFakeCaptcha(mobile: string) {
    return request(`/api/login/captcha?mobile=${mobile}`);
}
