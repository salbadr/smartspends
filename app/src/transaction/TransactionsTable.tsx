import { useTable,  } from "@tanstack/react-table";
import type { Transaction } from "../common/types";
import { columns, features, type TableColumns } from "./TranscationsTableSettings";

export function TransactionsTable({ transactions }: { transactions: Transaction[] }) {
    const table = useTable({
        columns: columns,
        features: features,
        data: transactions
    })

    if(transactions.length === 0){
        return (<></>)
    }

    return (<section className="w-full my-10">
        <table className="w-full p-2">
            <thead>
                {table.getHeaderGroups().map((headerGroup) => (
                    <tr key={headerGroup.id} className="border-b-2 text-sm uppercase">
                        {headerGroup.headers.map((header) => (
                            <th key={header.id} className={`py-2 px-4 ${header.column.columnDef.meta?.className ?? ''}`}>
                                <table.FlexRender header={
                                    header
                                } />
                            </th>))}
                    </tr>
                ))}
            </thead>
            <tbody>
                {table.getRowModel().rows.map((row) => (
                    <tr key={row.id} className="border-b-1 border-stone-200 text-sm even:bg-stone-50">
                        {row.getAllCells().map((cell) => {
                            const c = cell.column.columnDef as TableColumns
                            return (
                                <td key={cell.id} className={`p-4 ${c.className ?? ''}`}>
                                    <table.FlexRender cell={cell} />
                                </td>
                            )
                        }
                        )}
                    </tr>
                ))}
            </tbody>
        </table>
    </section>)
}