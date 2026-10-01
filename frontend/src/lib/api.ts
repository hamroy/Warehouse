const API_BASE = import.meta.env.VITE_API_URL ?? 'http://localhost:8080'

export async function apiFetch<T>(
  path: string,
  init?: RequestInit,
): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`, {
    headers: { 'Content-Type': 'application/json' },
    ...init,
  })

  if (!res.ok) {
    let message = `Request failed with status ${res.status}`
    try {
      const body = (await res.json()) as { message?: string }
      if (body.message) message = body.message
    } catch {
      // keep default message when body is not JSON
    }
    throw new Error(message)
  }

  return (await res.json()) as T
}
