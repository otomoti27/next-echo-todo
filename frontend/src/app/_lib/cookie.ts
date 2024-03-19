'server-only'

import { cookies } from 'next/headers'

export const setServerCookie = () => {
  return cookies()
    .getAll()
    .map((cookie) => `${cookie.name}=${cookie.value}`)
    .join('; ')
}
