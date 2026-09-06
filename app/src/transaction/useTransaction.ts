import { useState } from 'react';
import type { GetExpenditureResponse } from '../common/types';
import { TransactionClient } from './transaction-client';
import type { Transaction, Summary } from '../common/types';

export function useTransaction() {

    const [transactions, setTransactions] = useState<Transaction[]>([]);
    const [summary, setSummary] = useState<Summary[]>([]);
    const getExpenditure = async (prevState: unknown, formData: FormData) => {

        try {
            const payload = {
                expenses: formData.get('expenses')?.toString() ?? ''
            }

            const transactionClient = new TransactionClient();
            const data: GetExpenditureResponse = await transactionClient.getExpenses(payload)
            if (data.transactions && data.summary) {
                setTransactions(data.transactions)
                setSummary(data.summary)

            }
        }
        catch (error) {
            console.error(error);
        }
    }

    return {
        isLoaded: transactions.length > 0 && summary.length > 0,
        transactions,
        summary,
        getExpenditure
    }

}