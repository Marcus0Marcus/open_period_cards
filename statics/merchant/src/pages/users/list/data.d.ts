export interface TableListItem {
    key: number;
    name: string;
    phone: string;
}

export interface TableListPagination {
    total: number;
    pageSize: number;
    current: number;
}

export interface TableListData {
    list: TableListItem[];
    pagination: Partial<TableListPagination>;
}

export interface TableListParams {
    status?: string;
    name?: string;
    desc?: string;
    key?: number;
    pageSize?: number;
    currentPage?: number;
    filter?: Record<string, any[]>;
    sorter?: Record<string, any>;
}
