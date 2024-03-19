'server-only'

import axiosBase from 'axios'
import { cookies } from 'next/headers'
import { API_URL } from '@/config'

export const axios = axiosBase.create({
  baseURL: API_URL,
  withCredentials: true,
  withXSRFToken: true,
  headers: {
    'Content-Type': 'application/json',
  },
  xsrfCookieName: '_csrf',
  xsrfHeaderName: 'X-CSRF-Token',
})

axios.interceptors.request.use(
  (config) => {
    const csrfToken = cookies().get('_csrf')?.value
    config.headers['Cookie'] = cookies()
      .getAll()
      .map((cookie) => `${cookie.name}=${cookie.value}`)
      .join('; ')
    config.headers['X-CSRF-Token'] = csrfToken
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

axios.interceptors.response.use(
  (response) => response,
  (error) => {
    console.log('axios error', error)
    return Promise.reject(error.response?.data)
  },
)
