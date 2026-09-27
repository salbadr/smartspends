import { useState } from 'react';
import { SummaryTable } from './SummaryTable';
import { TransactionsTable } from './TransactionsTable';
import type { Summary, Transaction } from '../common/types';

export function TransactionsResult({summary, transactions}: {summary: Summary[], transactions: Transaction[]}) {

  const [display, setDisplay] = useState<'transactions' | 'summary'>('transactions')

  const getDisplayClassnames = (activeDisplay: 'transactions' | 'summary') => {
    let displayClassnames = 'pb-1 cursor-pointer';
    if (display === activeDisplay) {
      displayClassnames += ' text-blue-700 font-bold border-b-2 border-blue-700';
    }

    return displayClassnames
  }
  return (
    <section className='w-full'>
      <div className='grid grid-cols-[0.5fr_0.25fr_3.25fr] gap-4 text-sm w-full mt-10 border-b-1 border-stone-300'>
        <div className={getDisplayClassnames('transactions')} onClick={() => setDisplay('transactions')}>
          Transactions
        </div>
        <div className={getDisplayClassnames('summary')} onClick={() => setDisplay('summary')}>
          Summary
        </div>
      </div>
      {display === 'summary' && <SummaryTable summary={summary} />}
      {display === 'transactions' && <TransactionsTable transactions={transactions} />}
    </section>
  )
}