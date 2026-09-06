import { metaHelper, tableFeatures, type ColumnDef } from "@tanstack/react-table";
import type { Summary } from "../common/types";
export interface MyColumnMeta {
    className?: string;
}
export const features = tableFeatures({
    columnMeta: metaHelper<MyColumnMeta>()
})
export type TableColumns = ColumnDef<typeof features, Summary> & {
    className?: string;
}
export const columns: TableColumns[] = [
    {
        accessorKey: 'category',
        header: 'Category',
        className: 'first-letter:uppercase',
        cell: (info) => info.getValue(),
        meta: {
            className: 'text-left'
        }
    },
    {
        accessorKey: 'amount',
        header: 'Amount',
        className: 'text-stone-900 font-bold text-right',
        cell: (info) => `$${info.getValue()}`,
        meta: {
            className: 'text-right'
        }
    },
]