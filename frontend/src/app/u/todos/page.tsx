import { getMe } from '@/features/auth'

export default async function Todos() {
  const user = await getMe()
  console.log(user)
  return (
    <div>
      <h1>Todoリスト</h1>
    </div>
  )
}
