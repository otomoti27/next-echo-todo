import { User } from '../types'
import { setServerCookie } from '@/app/_lib/cookie'
import { API_URL } from '@/config'
import { fetcher } from '@/lib/fetch'

export const getMe = (cache?: RequestCache) => {
  return fetcher<User>(API_URL + '/auth/me', {
    credentials: 'include',
    headers: {
      Cookie: setServerCookie(),
    },
    cache,
  })
}
