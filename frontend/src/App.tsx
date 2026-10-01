import { Link, Route, Routes } from 'react-router-dom'

function HomePage() {
  return (
    <main>
      <h1>Warehouse Management</h1>
      <p>
        <Link to="/items">Manage Items</Link>
      </p>
    </main>
  )
}

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/items" element={<div>Item list (TASK-701)</div>} />
    </Routes>
  )
}
