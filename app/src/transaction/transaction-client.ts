import type { GetExpenditureRequest, GetExpenditureResponse } from "../common/types";

export class TransactionClient {
    private readonly _endpoint: string;

    constructor() {
        this._endpoint = 'http://localhost:8000/api/v1'
    }

    public async getExpenses(payload: GetExpenditureRequest): Promise<GetExpenditureResponse> {
        const endpoint = `${this._endpoint}/expenses`
        const options: RequestInit = {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: {
                'Content-Type': 'application/json'
            },
        }
        try {
            const response = await fetch(endpoint, options)
            const data: GetExpenditureResponse = await response.json();
            if (data.transactions && data.summary) {
                return data

            }

            return {
                title: '',
                date: '',
                transactions: [],
                summary: []
            }
        }
        catch (error) {
            throw new Error('Failed to get expenses', { cause: error });
        }

    }



}