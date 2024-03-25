'use server'

import { cookies } from 'next/headers'

export const logout = async () => {
  cookies().delete('token')
}
