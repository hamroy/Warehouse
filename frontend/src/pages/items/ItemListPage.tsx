import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { useItems, useDeleteItem } from '../../hooks/useItems'
import { useDebouncedValue } from '../../hooks/useDebounce'
import type { Item } from '../../services/itemApi'

const CATEGORIES = ['Electronics', 'Furniture']

function ItemRow({ item, onEdit, onDelete }: { item: Item; onEdit: (id: string) => void; onDelete: (id: string) => void }) {
  return (
    <tr>
      <td>{item.sku}</td>
      <td>{item.name}</td>
      <td>{item.category}</td>
      <td>{item.unit}</td>
      <td>
        <button type="button" onClick={() => onEdit(item.id)}>
          Edit
        </button>
        <button type="button" onClick={() => onDelete(item.id)}>
          Delete
        </button>
      </td>
    </tr>
  )
}

export default function ItemListPage() {
  const [search, setSearch] = useState('')
  const [category, setCategory] = useState('')
  const [page, setPage] = useState(1)
  const limit = 10

  const debouncedSearch = useDebouncedValue(search, 300)

  // reset page when filters change
  useEffect(() => {
    setPage(1)
  }, [debouncedSearch, category])

  const { data, isLoading, isError, error } = useItems({
    search: debouncedSearch,
    category,
    page,
    limit,
  })

  const deleteItem = useDeleteItem()

  const handleEdit = (id: string) => {
    window.location.hash = `#/items/${id}/edit`
  }

  const handleDelete = (id: string) => {
    if (!window.confirm('Delete this item?')) return
    deleteItem.mutate(id)
  }

  const items = data?.data ?? []
  const meta = data?.meta
  const totalPages = meta ? Math.max(1, Math.ceil(meta.total / limit)) : 1

  return (
    <main>
      <h1>Warehouse Items</h1>

      <div className="toolbar">
        <input
          type="search"
          placeholder="Search SKU / Name"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          aria-label="Search SKU or name"
        />
        <select
          value={category}
          onChange={(e) => setCategory(e.target.value)}
          aria-label="Category filter"
        >
          <option value="">All Categories</option>
          {CATEGORIES.map((c) => (
            <option key={c} value={c}>
              {c}
            </option>
          ))}
        </select>
        <Link to="/items/new">
          <button type="button">+ Add Item</button>
        </Link>
      </div>

      {isLoading && <p>Loading...</p>}
      {isError && <p role="alert">Error: {String((error as Error)?.message ?? error)}</p>}

      {!isLoading && !isError && (
        <>
          {items.length === 0 ? (
            <p>No items found.</p>
          ) : (
            <table>
              <thead>
                <tr>
                  <th>SKU</th>
                  <th>Name</th>
                  <th>Category</th>
                  <th>Unit</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                {items.map((item) => (
                  <ItemRow key={item.id} item={item} onEdit={handleEdit} onDelete={handleDelete} />
                ))}
              </tbody>
            </table>
          )}

          <nav className="pagination" aria-label="Pagination">
            <button type="button" disabled={page <= 1} onClick={() => setPage((p) => p - 1)}>
              Previous
            </button>
            <span>
              Page {page} of {totalPages}
            </span>
            <button
              type="button"
              disabled={page >= totalPages}
              onClick={() => setPage((p) => p + 1)}
            >
              Next
            </button>
          </nav>
        </>
      )}
    </main>
  )
}
