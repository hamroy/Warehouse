import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { itemApi } from '../services/itemApi'
import type { Item, ListItemsParams } from '../services/itemApi'

export function useItems(params: ListItemsParams) {
  return useQuery({
    queryKey: ['items', params],
    queryFn: () => itemApi.getItems(params),
  })
}

export function useItem(id: string) {
  return useQuery({
    queryKey: ['item', id],
    queryFn: () => itemApi.getItem(id),
    enabled: Boolean(id),
  })
}

export function useCreateItem() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (payload: Parameters<typeof itemApi.createItem>[0]) =>
      itemApi.createItem(payload),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['items'] })
    },
  })
}

export function useUpdateItem() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: Parameters<typeof itemApi.updateItem>[1] }) =>
      itemApi.updateItem(id, payload),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['items'] })
    },
  })
}

export function useDeleteItem() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) => itemApi.deleteItem(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['items'] })
    },
  })
}

export type { Item }
