import type {ColumnType} from "ant-design-vue/es/table";

export interface AntTableCell {
    text: any;
    value: any;
    record: any;
    index: number;
    column: ColumnType<any>;
}

export interface AntColumnsType extends ColumnType, Record<string, any> {
    key?: string;
    align: any;
}