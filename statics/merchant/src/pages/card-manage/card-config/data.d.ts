export interface CardConfigListItem {
    key: number;
    period_times: number;
    total_times: string;
    describe: string;
}

export interface CardConfigListPagination {
    total: number;
    pageSize: number;
    current: number;
}

export interface CardTypeListData {
    list: CardConfigListItem[];
    pagination: Partial<CardConfigListPagination>;
}

export interface CardTypeListParams {
    status?: string;
    name?: string;
    desc?: string;
    key?: number;
    pageSize?: number;
    currentPage?: number;
    filter?: Record<string, any[]>;
    sorter?: Record<string, any>;
}
