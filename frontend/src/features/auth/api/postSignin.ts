import { axios } from '@/app/_lib/axios'

type Params = {
  email: string
  password: string
}

export const postSignin = (params: Params) => {
  return axios.post('/login', params)
}
