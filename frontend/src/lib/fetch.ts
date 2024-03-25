import { ApiError } from './error'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export async function fetcher<T = any>(
  input: RequestInfo,
  init?: RequestInit & { next?: { revalidate?: number; tag?: Array<string> } },
): Promise<T> {
  const res = await fetch(input, {
    headers: {
      'Content-Type': 'application/json',
    },
    ...init,
  })

  if (!res.ok) {
    try {
      const err: { message: string } = await res.json()
      throw new ApiError(err.message, res.status)
    } catch (e) {
      throw new ApiError(String(res.status), res.status)
    }
  }

  return await res.json()
}
