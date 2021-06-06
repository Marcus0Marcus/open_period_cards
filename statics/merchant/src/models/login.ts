import { stringify } from 'querystring';
import type { Reducer, Effect } from 'umi';
import { history } from 'umi';

import { AccountLogin } from '@/services/login';
import { setAuthority } from '@/utils/authority';
import { getPageQuery } from '@/utils/utils';
import { message } from 'antd';

export type StateType = {
    status?: 'ok' | 'error';
    type?: string;
    currentAuthority?: 'user' | 'guest' | 'admin';
};

export type LoginModelType = {
    namespace: string;
    state: StateType;
    effects: {
        login: Effect;
        logout: Effect;
    };
    reducers: {
        changeLoginStatus: Reducer<StateType>;
    };
};

const Model: LoginModelType = {
    namespace: 'login',

    state: {
        status: undefined,
    },

    effects: {
        *login({ payload }, { call, put }) {
            const response = yield call(AccountLogin, payload);
            yield put({
                type: 'changeLoginStatus',
                payload: response,
            });
            console.log(response)
            debugger
            // Login successfully
            if (response.code === 0) {
                const urlParams = new URL(window.location.href);
                const params = getPageQuery();
                message.success('🎉 🎉 🎉  登录成功！');
                let { redirect } = params as { redirect: string };
                if (redirect) {
                    const redirectUrlParams = new URL(redirect);
                    if (redirectUrlParams.origin === urlParams.origin) {
                        redirect = redirect.substr(urlParams.origin.length);
                        if (redirect.match(/^\/.*#/)) {
                            redirect = redirect.substr(redirect.indexOf('#') + 1);
                        }
                    } else {
                        window.location.href = '/';
                        return;
                    }
                }
                history.replace(redirect || '/');
            }
            else {
                message.error(response.message)
            }
        },

        logout() {

            const { redirect } = getPageQuery();
            // console.log(window.location.pathname !== '/login/login' && !redirect)
            // Note: There may be security issues, please note
            if (window.location.pathname !== '/login/login' && !redirect) {
                history.replace({
                    pathname: '/login/login',
                    search: stringify({
                        redirect: window.location.href,
                    }),
                });
            }
        },
    },

    reducers: {
        changeLoginStatus(state, { payload }) {
            setAuthority(payload.currentAuthority);
            return {
                ...state,
                status: payload.status,
                type: payload.type,
            };
        },
    },
};

export default Model;
