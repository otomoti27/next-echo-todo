import { BASE_URL } from '@/config'

export const fetchCsrfToken = async () => {
  return await fetch(BASE_URL + '/api/csrf', {
    credentials: 'include',
    cache: 'no-store',
  })
}
