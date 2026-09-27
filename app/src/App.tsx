import { useActionState, useState } from 'react';
import './App.css'
import logo from './assets/logo.svg'
import { useTransaction } from './transaction/useTransaction';
import { TransactionsResult } from './transaction/TransactionsResult';

function App() {
  const [expenses, setExpenses] = useState('');
  const { summary, transactions, getExpenditure, isLoaded } = useTransaction()
  const [, formAction, isPending] = useActionState(getExpenditure, null);

  return (
    <>
      <header className='border-b-1 border-stone-300 px-10 flex justify-between items-center'>
        <div className='w-60'>
          <img src={logo} />

        </div>
        <div className='text-xs uppercase'>
          csv categorizer
        </div>
      </header>
      <main className='flex flex-col items-center justify-center mx-30 my-10'>
        <section className='border-b-1 border-stone-300 mx-10 w-full'>
          <div className='mb-10 w-full'>
            <h1 className='text-[32px] font-bold text-stone- 900'>Categorize your expenses</h1>
            <p>Paste your expenses below. We'll classify each row and surface insights across categories.</p>
          </div>
          <form action={formAction} className='flex flex-col gap-2 align-center justify-center mb-10 w-full'>
            <label htmlFor="csvInput" className='text-xs uppercase'>Expenses</label>
            <textarea id="csvInput" name='expenses' className='w-full h-[200px] border-1 border-stone-300 p-2 hover:border-blue-500 bg-stone-100' placeholder='Paste your CSV here' defaultValue={expenses} onChange={(event) => setExpenses(event.target.value)}>
            </textarea>

            <button type='submit' className='bg-blue-500 self-end p-4 border-1 text-white w-[200px] disabled:bg-stone-300 hover:cursor-pointer hover:disabled:cursor-not-allowed' disabled={!isPending && expenses === ''}>Submit</button>
          </form>


        </section>
        {isPending && <section className='w-full'>
          <h2>Getting results...</h2>
        </section>}
        {isLoaded && <TransactionsResult summary={summary} transactions={transactions} />}
      </main>
    </>


  )
}

export default App
