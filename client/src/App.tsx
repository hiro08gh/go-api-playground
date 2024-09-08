import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [data, setData] = useState<string | null>(null)

  useEffect(() => {
    fetch('http://localhost:3001/users').then((res) => res.json()).then((data: string) => setData(data))
  }, [])

  return (
    <>
      {data && (
        <h1>{data}</h1>
      )}
    </>
  )
}

export default App
