// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';
const { REACT_APP_ENV } = process.env;
export default defineConfig({
    hash: true,
    antd: {},
    dva: {
        hmr: true,
    },
    history: {
        type: 'browser',
    },
    locale: {
        // default zh-CN
        default: 'zh-CN',
        antd: true,
        // default true, when it is true, will use `navigator.language` overwrite default
        baseNavigator: true,
    },
    dynamicImport: {
        loading: '@/components/PageLoading/index',
    },
    targets: {
        ie: 11,
    },
    // umi routes: https://umijs.org/docs/routing
    routes: [
        {
            path: '/',
            component: '../layouts/BlankLayout',
            routes: [
                {
                    path: '/login',
                    component: '../layouts/UserLayout',
                    routes: [
                        {
                            path: '/login/login',
                            name: 'login',
                            component: './login/login',
                        },
                        {
                            path: '/login',
                            redirect: '/login/login',
                        },
                        {
                            name: 'register-result',
                            icon: 'smile',
                            path: '/login/register-result',
                            component: './login/register-result',
                        },
                        {
                            name: 'register',
                            icon: 'smile',
                            path: '/login/register',
                            component: './login/register',
                        },
                        {
                            component: '404',
                        },
                    ],
                },
                {
                    path: '/',
                    component: '../layouts/BasicLayout',
                    Routes: ['src/pages/Authorized'],
                    authority: ['admin', 'user'],
                    routes: [
                        {
                            path: '/',
                            redirect: '/dashboard/analysis',
                        },
                        {
                            path: '/dashboard',
                            name: 'dashboard',
                            icon: 'dashboard',
                            routes: [
                                {
                                    path: '/',
                                    redirect: '/dashboard/analysis',
                                },
                                {
                                    name: 'analysis',
                                    icon: 'smile',
                                    path: '/dashboard/analysis',
                                    component: './dashboard/analysis',
                                },
                                {
                                    name: 'monitor',
                                    icon: 'smile',
                                    path: '/dashboard/monitor',
                                    component: './dashboard/monitor',
                                },
                                {
                                    name: 'workplace',
                                    icon: 'smile',
                                    path: '/dashboard/workplace',
                                    component: './dashboard/workplace',
                                },
                            ],
                        },
                        {
                            path: '/users',
                            icon: 'idcard',
                            name: 'users',
                            routes: [
                                {
                                    path: '/',
                                    redirect: '/users/search',
                                },
                                {
                                    name: 'search',
                                    icon: 'smile',
                                    path: '/users/search',
                                    component: './users/list',
                                },
                            ],
                        },
                        {
                            path: '/card-manage',
                            icon: 'credit-card',
                            name: 'card-manage',
                            routes: [
                                {
                                    path: '/',
                                    redirect: '/card-manage/card-config',
                                },
                                {
                                    name: 'card-config',
                                    icon: 'smile',
                                    path: '/card-manage/card-config',
                                    component: './card-manage/card-config',
                                },
                                {
                                    name: 'card-order',
                                    icon: 'smile',
                                    path: '/card-manage/card-order',
                                    component: './list/table-list',
                                },
                                {
                                    name: 'card-used',
                                    icon: 'smile',
                                    path: '/card-manage/card-used',
                                    component: './list/table-list',
                                }
                            ],
                        },
                        {
                            path: '/delivery-log',
                            icon: 'form',
                            name: 'delivery-log',
                            routes: [
                                {
                                    path: '/',
                                    redirect: '/delivery-log/list',
                                },
                                {
                                    name: 'list',
                                    icon: 'smile',
                                    path: '/delivery-log/list',
                                    component: './list/table-list',
                                },
                            ],
                        },
                        {
                            name: 'account',
                            icon: 'user',
                            path: '/account',
                            routes: [
                                {
                                    path: '/',
                                    redirect: '/account/center',
                                },
                                {
                                    name: 'center',
                                    icon: 'smile',
                                    path: '/account/center',
                                    component: './account/center',
                                },
                                {
                                    name: 'settings',
                                    icon: 'smile',
                                    path: '/account/settings',
                                    component: './account/settings',
                                },
                            ],
                        },
                        {
                            component: '404',
                        },
                    ],
                },
            ],
        },
    ],
    // Theme for antd: https://ant.design/docs/react/customize-theme-cn
    theme: {
        'primary-color': defaultSettings.primaryColor,
    },
    title: false,
    ignoreMomentLocale: true,
    proxy: proxy[REACT_APP_ENV || 'dev'],
    manifest: {
        basePath: '/',
    },
    esbuild: {},
});
