import { PlusOutlined } from '@ant-design/icons';
import { Button, Divider, message, Input, Drawer } from 'antd';
import React, { useState, useRef } from 'react';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import ProDescriptions from '@ant-design/pro-descriptions';
import CreateForm from './components/CreateForm';
import type { FormValueType } from './components/UpdateForm';
import UpdateForm from './components/UpdateForm';
import type { CardConfigListItem } from './data.d';
import { queryCardTypeList, updateCardType, addCardType, removeCardType } from './service';

/**
 * 添加节点
 *
 * @param fields
 */
const handleAdd = async (fields: CardConfigListItem) => {
    const hide = message.loading('正在添加');
    try {
        await addCardType({ ...fields });
        hide();
        message.success('添加成功');
        return true;
    } catch (error) {
        hide();
        message.error('添加失败请重试！');
        return false;
    }
};

/**
 * 更新节点
 *
 * @param fields
 */
const handleUpdate = async (fields: FormValueType) => {
    const hide = message.loading('正在配置');
    try {
        await updateCardType({
            name: fields.name,
            desc: fields.desc,
            key: fields.key,
        });
        hide();

        message.success('配置成功');
        return true;
    } catch (error) {
        hide();
        message.error('配置失败请重试！');
        return false;
    }
};

/**
 * 删除节点
 *
 * @param selectedRows
 */
const handleRemove = async (selectedRows: CardConfigListItem[]) => {
    const hide = message.loading('正在删除');
    if (!selectedRows) return true;
    try {
        await removeCardType({
            key: selectedRows.map((row) => row.key),
        });
        hide();
        message.success('删除成功，即将刷新');
        return true;
    } catch (error) {
        hide();
        message.error('删除失败，请重试');
        return false;
    }
};

const CardConfigList: React.FC<{}> = () => {
    const [createModalVisible, handleModalVisible] = useState<boolean>(false);
    const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
    const [stepFormValues, setStepFormValues] = useState({});
    const actionRef = useRef<ActionType>();
    const [row, setRow] = useState<CardConfigListItem>();
    const [selectedRowsState, setSelectedRows] = useState<CardConfigListItem[]>([]);
    const columns: ProColumns<CardConfigListItem>[] = [
        {
            title: 'ID',
            dataIndex: 'name',
            tip: '卡片ID是唯一的',
            hideInForm: true,
            render: (dom, entity) => {
                return <a onClick={() => setRow(entity)}>{dom}</a>;
            },
        },
        {
            title: '周期内次数',
            formItemProps: {
                rules: [
                    {
                        required: true,
                        message: '周期内次数为必填项',
                    },
                ],
            },
            dataIndex: 'period_times',
        },
        {
            title: '总次数',
            dataIndex: 'total_times',
            sorter: true,
            hideInForm: false,
            formItemProps: {
                rules: [
                    {
                        required: true,
                        message: '总次数为必填项',
                    },
                ],
            },
        },
        {
            title: '描述',
            dataIndex: 'describe',
            sorter: true,
            hideInForm: false,
            valueType: 'textarea',
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record) => [
                <a
                    onClick={() => {
                        handleUpdateModalVisible(true);
                        setStepFormValues(record);
                    }}
                >
                    配置
        </a>,
                <Divider type="vertical" />,
                <a href="">订阅警报</a>,
            ],
        },
    ];

    return (
        <PageContainer>
            <ProTable<CardConfigListItem>
                headerTitle="查询表格"
                actionRef={actionRef}
                rowKey="key"
                search={{
                    labelWidth: 120,
                }}
                toolBarRender={() => [
                    <Button type="primary" onClick={() => handleModalVisible(true)}>
                        <PlusOutlined /> 新建
          </Button>,
                ]}
                request={(params, sorter, filter) => queryCardTypeList({ ...params, sorter, filter })}
                columns={columns}
                rowSelection={{
                    onChange: (_, selectedRows) => setSelectedRows(selectedRows),
                }}
            />
            {selectedRowsState?.length > 0 && (
                <FooterToolbar
                    extra={
                        <div>
                            已选择 <a style={{ fontWeight: 600 }}>{selectedRowsState.length}</a> 项&nbsp;&nbsp;
              <span>
                                服务调用次数总计 {selectedRowsState.reduce((pre, item) => pre + item.callNo, 0)} 万
              </span>
                        </div>
                    }
                >
                    <Button
                        onClick={async () => {
                            await handleRemove(selectedRowsState);
                            setSelectedRows([]);
                            actionRef.current?.reloadAndRest?.();
                        }}
                    >
                        批量删除
          </Button>
                    <Button type="primary">批量审批</Button>
                </FooterToolbar>
            )}
            <CreateForm onCancel={() => handleModalVisible(false)} modalVisible={createModalVisible}>
                <ProTable<CardConfigListItem, CardConfigListItem>
                    onSubmit={async (value) => {
                        const success = await handleAdd(value);
                        if (success) {
                            handleModalVisible(false);
                            if (actionRef.current) {
                                actionRef.current.reload();
                            }
                        }
                    }}
                    rowKey="key"
                    type="form"
                    columns={columns}
                />
            </CreateForm>
            {stepFormValues && Object.keys(stepFormValues).length ? (
                <UpdateForm
                    onSubmit={async (value) => {
                        const success = await handleUpdate(value);
                        if (success) {
                            handleUpdateModalVisible(false);
                            setStepFormValues({});
                            if (actionRef.current) {
                                actionRef.current.reload();
                            }
                        }
                    }}
                    onCancel={() => {
                        handleUpdateModalVisible(false);
                        setStepFormValues({});
                    }}
                    updateModalVisible={updateModalVisible}
                    values={stepFormValues}
                />
            ) : null}

            <Drawer
                width={600}
                visible={!!row}
                onClose={() => {
                    setRow(undefined);
                }}
                closable={false}
            >
                {row?.key && (
                    <ProDescriptions<CardConfigListItem>
                        column={2}
                        title={row?.key}
                        request={async () => ({
                            data: row || {},
                        })}
                        params={{
                            id: row?.key,
                        }}
                        columns={columns}
                    />
                )}
            </Drawer>
        </PageContainer>
    );
};

export default CardConfigList;
