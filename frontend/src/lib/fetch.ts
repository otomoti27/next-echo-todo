import { ApiError } from './error'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export async function fetcher<T = any>(
  input: RequestInfo,
  init?: RequestInit & { next?: { revalidate?: number; tag?: Array<string> } },
): Promise<T> {
  const res = await fetch(input, init)

  if (!res.ok) {
    try {
      const err: { message: string } = await res.json()
      throw new ApiError(err.message, res.status)
    } catch {
      throw new ApiError('レスポンス解析に失敗しました', res.status)
    }
  }

  return await res.json()
}
