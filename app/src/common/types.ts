export type Transaction = {
  date: string,
  description: string,
  category: string,
  amount: number,
  confidence: number,
}

export type Summary = {
  category: string,
  amount: number
}


export type GetExpenditureRequest = {
  expenses: string,
}


export type GetExpenditureResponse = {
  title: string,
  date: string,
  transactions: Transaction[],
  summary: Summary[]
  confidence: number,
}

