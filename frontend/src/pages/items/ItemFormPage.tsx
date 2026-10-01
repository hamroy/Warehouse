import { useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { useItem, useCreateItem, useUpdateItem } from '../../hooks/useItems'

interface FormState {
  sku: string
  name: string
  category: string
  unit: string
}

const emptyForm: FormState = { sku: '', name: '', category: '', unit: '' }

interface FormErrors {
  sku?: string
  name?: string
  category?: string
  unit?: string
}

function validate(form: FormState): FormErrors {
  const errors: FormErrors = {}
  if (!form.sku.trim()) errors.sku = 'SKU is required'
  if (!form.name.trim()) errors.name = 'Name is required'
  if (!form.category.trim()) errors.category = 'Category is required'
  if (!form.unit.trim()) errors.unit = 'Unit is required'
  return errors
}

export default function ItemFormPage() {
  const { id } = useParams<{ id: string }>()
  const isEdit = Boolean(id)
  const navigate = useNavigate()

  const [form, setForm] = useState<FormState>(emptyForm)
  const [errors, setErrors] = useState<FormErrors>({})
  const [submitError, setSubmitError] = useState<string | null>(null)

  const itemQuery = useItem(id ?? '')
  const createItem = useCreateItem()
  const updateItem = useUpdateItem()

  // populate form when editing
  useEffect(() => {
    if (isEdit && itemQuery.data?.data) {
      const it = itemQuery.data.data
      setForm({ sku: it.sku, name: it.name, category: it.category, unit: it.unit })
    }
  }, [isEdit, itemQuery.data])

  const set = (field: keyof FormState) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm((f) => ({ ...f, [field]: e.target.value }))
    setErrors((err) => ({ ...err, [field]: undefined }))
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitError(null)

    const validation = validate(form)
    if (Object.keys(validation).length > 0) {
      setErrors(validation)
      return
    }

    const payload = {
      sku: form.sku.trim(),
      name: form.name.trim(),
      category: form.category.trim(),
      unit: form.unit.trim(),
    }

    const onSuccess = () => {
      navigate('/items')
    }
    const onFailure = (err: unknown) => {
      const message = err instanceof Error ? err.message : 'Request failed'
      setSubmitError(message)
    }

    if (isEdit && id) {
      updateItem.mutate({ id, payload }, { onSuccess, onError: onFailure })
    } else {
      createItem.mutate(payload, { onSuccess, onError: onFailure })
    }
  }

  if (isEdit && itemQuery.isLoading) {
    return <main>Loading item...</main>
  }
  if (isEdit && itemQuery.isError) {
    return <main>Failed to load item.</main>
  }

  return (
    <main>
      <h1>{isEdit ? 'Edit Item' : 'Create Item'}</h1>

      <form onSubmit={handleSubmit} noValidate>
        <div className="field">
          <label htmlFor="sku">SKU</label>
          <input id="sku" type="text" value={form.sku} onChange={set('sku')} />
          {errors.sku && <span className="error">{errors.sku}</span>}
        </div>

        <div className="field">
          <label htmlFor="name">Name</label>
          <input id="name" type="text" value={form.name} onChange={set('name')} />
          {errors.name && <span className="error">{errors.name}</span>}
        </div>

        <div className="field">
          <label htmlFor="category">Category</label>
          <input id="category" type="text" value={form.category} onChange={set('category')} />
          {errors.category && <span className="error">{errors.category}</span>}
        </div>

        <div className="field">
          <label htmlFor="unit">Unit</label>
          <input id="unit" type="text" value={form.unit} onChange={set('unit')} />
          {errors.unit && <span className="error">{errors.unit}</span>}
        </div>

        {submitError && <p role="alert" className="error">{submitError}</p>}

        <div className="actions">
          <button type="submit" disabled={createItem.isPending || updateItem.isPending}>
            {isEdit ? 'Update' : 'Create'}
          </button>
          <Link to="/items">
            <button type="button">Cancel</button>
          </Link>
        </div>
      </form>
    </main>
  )
}
