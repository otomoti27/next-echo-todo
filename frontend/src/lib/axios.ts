'client-only'

import axiosBase from 'axios'
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

axios.interceptors.response.use(
  (response) => response,
  (error) => {
    console.log('axios error', error)
    return Promise.reject(error.response?.data)
  },
)
