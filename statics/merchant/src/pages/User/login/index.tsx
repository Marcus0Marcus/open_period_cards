import {
	LockOutlined,
	UserOutlined,
} from '@ant-design/icons';
import { Alert, Space, Tabs } from 'antd';
import React, { useState } from 'react';
import ProForm, { ProFormText } from '@ant-design/pro-form';
import { useIntl, connect, FormattedMessage } from 'umi';
import type { Dispatch } from 'umi';
import type { StateType } from '@/models/login';
import type { LoginParamsType } from '@/services/login';
import type { ConnectState } from '@/models/connect';

import styles from './index.less';

export type LoginProps = {
	dispatch: Dispatch;
	userLogin: StateType;
	submitting?: boolean;
};

const LoginMessage: React.FC<{
	content: string;
}> = ({ content }) => (
	<Alert
		style={{
			marginBottom: 24,
		}}
		message={content}
		type="error"
		showIcon
	/>
);

const Login: React.FC<LoginProps> = (props) => {
	const { userLogin = {}, submitting } = props;
	const { status, type: loginType } = userLogin;
	const [type, setType] = useState<string>('account');
	const intl = useIntl();

	const handleSubmit = (values: LoginParamsType) => {
		const { dispatch } = props;
		dispatch({
			type: 'login/login',
			payload: { ...values, type },
		});
	};
	return (
		<div className={styles.main}>
			<ProForm
				initialValues={{
					autoLogin: true,
				}}
				submitter={{
					render: (_, dom) => dom.pop(),
					submitButtonProps: {
						loading: submitting,
						size: 'large',
						style: {
							width: '100%',
						},
					},
				}}
				onFinish={(values) => {
					handleSubmit(values as LoginParamsType);
					return Promise.resolve();
				}}
			>
				<Tabs activeKey={type} onChange={setType}>
					<Tabs.TabPane
						key="account"
						tab={intl.formatMessage({
							id: 'pages.login.accountLogin.tab',
							defaultMessage: '账户密码登录',
						})}
					/>
				</Tabs>

				{status === 'error' && loginType === 'account' && !submitting && (
					<LoginMessage
						content={intl.formatMessage({
							id: 'pages.login.accountLogin.errorMessage',
							defaultMessage: '账户或密码错误',
						})}
					/>
				)}
				{type === 'account' && (
					<>
						<ProFormText
							name="userName"
							fieldProps={{
								size: 'large',
								prefix: <UserOutlined className={styles.prefixIcon} />,
							}}
							placeholder={intl.formatMessage({
								id: 'pages.login.username.placeholder',
								defaultMessage: 'q用户名: admin or user',
							})}
							rules={[
								{
									required: true,
									message: (
										<FormattedMessage
											id="pages.login.username.required"
											defaultMessage="请输入用户名!"
										/>
									),
								},
							]}
						/>
						<ProFormText.Password
							name="password"
							fieldProps={{
								size: 'large',
								prefix: <LockOutlined className={styles.prefixIcon} />,
							}}
							placeholder={intl.formatMessage({
								id: 'pages.login.password.placeholder',
								defaultMessage: '请输入密码！',
							})}
							rules={[
								{
									required: true,
									message: (
										<FormattedMessage
											id="pages.login.password.required"
											defaultMessage="请输入密码！"
										/>
									),
								},
							]}
						/>
					</>
				)}
			</ProForm>
			<Space>
				<a>{intl.formatMessage({
					id: 'pages.login.registerAccount',
					defaultMessage: '申请试用',
				})}</a>
			</Space>
		</div>
	);
};

export default connect(({ login, loading }: ConnectState) => ({
	userLogin: login,
	submitting: loading.effects['login/login'],
}))(Login);
