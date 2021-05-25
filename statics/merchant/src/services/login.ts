import request from '@/utils/request';

export type LoginParamsType = {
	phone: string;
	password: string;
};

export async function AccountLogin(params: LoginParamsType) {
	return request('/api/login/account', {
		method: 'POST',
		data: params,
	});
}
