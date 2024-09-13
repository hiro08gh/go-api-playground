import { fetcher } from './utils/fetcher'
import useSWR from 'swr'
import './App.css'

function App() {
   const { data, error, isLoading } = useSWR<string>('/users', fetcher)

   if (isLoading) return <div>Loading...</div>
   if (error) return <div>Error: {error}</div>

  return (
    <>
    {data}
    </>
  )
}

export default App
