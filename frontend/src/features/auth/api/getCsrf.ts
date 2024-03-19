import { axios } from '@/app/_lib/axios'

type Res = {
  csrf_token: string
}

export const getCsrf = async () => {
  const res = await axios.get<Res>('/csrf')
  return res.data.csrf_token
}
