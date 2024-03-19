'use client'

import { useEffect } from 'react'
// server-onlyもexportされているため
// eslint-disable-next-line no-restricted-imports
import { fetchCsrfToken } from '@/features/auth/api/fetchCsrfToken'

export const CSRFToken = () => {
  useEffect(() => {
    ;(async () => {
      await fetchCsrfToken()
    })()
  }, [])

  return null
}
