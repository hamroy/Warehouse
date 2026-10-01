import { apiFetch } from '../lib/api'

export interface Item {
  id: string
  sku: string
  name: string
  category: string
  unit: string
  created_at: string
  updated_at: string
  deleted_at: string | null
}

export interface ItemListMeta {
  page: number
  limit: number
  total: number
}

export interface ItemListResponse {
  success: boolean
  message: string
  data: Item[]
  meta: ItemListMeta
}

export interface ItemResponse {
  success: boolean
  message: string
  data: Item
}

export interface ErrorEnvelope {
  success: boolean
  message: string
  errors?: string[]
}

export interface ListItemsParams {
  search?: string
  category?: string
  page?: number
  limit?: number
}

function buildQuery(params: ListItemsParams): string {
  const query = new URLSearchParams()
  if (params.search) query.set('search', params.search)
  if (params.category) query.set('category', params.category)
  if (params.page) query.set('page', String(params.page))
  if (params.limit) query.set('limit', String(params.limit))
  const qs = query.toString()
  return qs ? `?${qs}` : ''
}

export const itemApi = {
  getItems(params: ListItemsParams = {}): Promise<ItemListResponse> {
    return apiFetch<ItemListResponse>(`/api/v1/items${buildQuery(params)}`)
  },

  getItem(id: string): Promise<ItemResponse> {
    return apiFetch<ItemResponse>(`/api/v1/items/${id}`)
  },

  createItem(payload: Omit<Item, 'id' | 'created_at' | 'updated_at' | 'deleted_at'>): Promise<ItemResponse> {
    return apiFetch<ItemResponse>('/api/v1/items', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
  },

  updateItem(
    id: string,
    payload: Omit<Item, 'id' | 'created_at' | 'updated_at' | 'deleted_at'>,
  ): Promise<ItemResponse> {
    return apiFetch<ItemResponse>(`/api/v1/items/${id}`, {
      method: 'PUT',
      body: JSON.stringify(payload),
    })
  },

  deleteItem(id: string): Promise<{ success: boolean; message: string }> {
    return apiFetch(`/api/v1/items/${id}`, { method: 'DELETE' })
  },
}
