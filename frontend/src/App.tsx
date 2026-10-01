import { HashRouter, Link, Route, Routes } from 'react-router-dom'
import ItemListPage from './pages/items/ItemListPage'
import ItemFormPage from './pages/items/ItemFormPage'

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
    <HashRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/items" element={<ItemListPage />} />
        <Route path="/items/new" element={<ItemFormPage />} />
        <Route path="/items/:id/edit" element={<ItemFormPage />} />
      </Routes>
    </HashRouter>
  )
}
