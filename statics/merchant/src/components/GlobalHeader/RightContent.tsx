import { Tooltip, Tag } from 'antd';
import type { Settings as ProSettings } from '@ant-design/pro-layout';
import { QuestionCircleOutlined } from '@ant-design/icons';
import React from 'react';
import type { ConnectProps } from 'umi';
import { connect, SelectLang } from 'umi';
import type { ConnectState } from '@/models/connect';
import Avatar from './AvatarDropdown';
import HeaderSearch from '../HeaderSearch';
import styles from './index.less';
import NoticeIconView from './NoticeIconView';

export type GlobalHeaderRightProps = {
	theme?: ProSettings['navTheme'] | 'realDark';
} & Partial<ConnectProps> &
	Partial<ProSettings>;
const ENVTagColor = {
	dev: 'orange',
	test: 'green',
	pre: '#87d068',
};

const GlobalHeaderRight: React.SFC<GlobalHeaderRightProps> = (props) => {
	const { theme, layout } = props;
	let className = styles.right;

	if (theme === 'dark' && layout === 'top') {
		className = `${styles.right}  ${styles.dark}`;
	}

	return (
		<div className={className}>
			<HeaderSearch
				className={`${styles.action} ${styles.search}`}
				placeholder="站内搜索"
				defaultValue="如何使用周期卡"
				options={[
					{
						label: <a href="https://umijs.org/zh/guide/umi-ui.html">umi ui</a>,
						value: '如何使用周期卡？',
					},
					{
						label: <a href="next.ant.design">如何添加新的周期卡类型？</a>,
						value: '如何添加新的周期卡类型？',
					},
					{
						label: <a href="https://protable.ant.design/">Pro Table</a>,
						value: 'Pro Table',
					},
					{
						label: <a href="https://prolayout.ant.design/">Pro Layout</a>,
						value: 'Pro Layout',
					},
				]}
				onSearch={value => {
					console.log('input', value);
				}}
			/>
			<Tooltip title="帮助文档">
				<a
					style={{
						color: 'inherit',
					}}
					target="_blank"
					href="https://pro.ant.design/docs/getting-started"
					rel="noopener noreferrer"
					className={styles.action}
				>
					<QuestionCircleOutlined />
				</a>
			</Tooltip>
			<NoticeIconView />
			<Avatar menu />
			{REACT_APP_ENV && (
				<span>
					<Tag color={ENVTagColor[REACT_APP_ENV]}>{REACT_APP_ENV}</Tag>
				</span>
			)}
			<SelectLang className={styles.action} />
		</div>
	);
};

export default connect(({ settings }: ConnectState) => ({
	theme: settings.navTheme,
	layout: settings.layout,
}))(GlobalHeaderRight);
