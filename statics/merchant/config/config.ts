import { defineConfig, utils } from 'umi';
const { winPath } = utils;
export default defineConfig({
    // 通过 package.json 自动挂载 umi 插件，不需再次挂载
    // plugins: [],
    singular: true,
    routes: [
        {
            path: '/',
            exact: true,
            component: 'HelloWorld',
        }
    ],
    antd: {},
    dva: {
        hmr: true,
    },
    locale: {
        default: 'zh-CN',
        baseNavigator: true,
    },
    dynamicImport: {
    },
    pwa: false,
    lessLoader: { javascriptEnabled: true },
    cssLoader: {
        modules: {
            getLocalIdent: (
                context: {
                    resourcePath: string;
                },
                _: string,
                localName: string,
            ) => {
                if (
                    context.resourcePath.includes('node_modules') ||
                    context.resourcePath.includes('ant.design.pro.less') ||
                    context.resourcePath.includes('global.less')
                ) {
                    return localName;
                }
                const match = context.resourcePath.match(/src(.*)/);
                if (match && match[1]) {
                    const antdProPath = match[1].replace('.less', '');
                    const arr = winPath(antdProPath)
                        .split('/')
                        .map((a: string) => a.replace(/([A-Z])/g, '-$1'))
                        .map((a: string) => a.toLowerCase());
                    return `antd-pro${arr.join('-')}-${localName}`.replace(/--/g, '-');
                }
                return localName;
            },
        }
    }
})