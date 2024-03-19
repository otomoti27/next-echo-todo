'use server'

import { z } from 'zod'
import { UserSchema } from './schema'
import { postSignup } from '@/features/auth'

export async function signup(data: z.infer<typeof UserSchema>) {
  const parseData = UserSchema.safeParse(data)

  if (!parseData.success) {
    return {
      success: false,
      errors: parseData.error.format(),
    }
  }

  try {
    const response = await postSignup({
      name: parseData.data.name,
      email: parseData.data.email,
      password: parseData.data.password,
    })
    console.log(response)
  } catch (error) {
    console.error(error)
    return {
      success: false,
      errors: ['エラーが発生しました'],
    }
  }

  return {
    success: true,
    errors: [],
  }
}
