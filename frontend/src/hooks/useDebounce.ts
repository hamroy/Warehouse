import { useEffect, useState } from 'react'

/**
 * useDebouncedValue returns a debounced copy of `value`.
 * The returned value only updates after `delay` ms of inactivity,
 * so rapid typing does not fire a request per keypress.
 */
export function useDebouncedValue<T>(value: T, delay = 300): T {
  const [debounced, setDebounced] = useState(value)

  useEffect(() => {
    const timer = setTimeout(() => setDebounced(value), delay)
    return () => clearTimeout(timer)
  }, [value, delay])

  return debounced
}
