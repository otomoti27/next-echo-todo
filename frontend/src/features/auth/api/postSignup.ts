import { axios } from '@/app/_lib/axios'

type Params = {
  name: string
  email: string
  password: string
}

export const postSignup = (params: Params) => {
  return axios.post('/signup', params)
}
