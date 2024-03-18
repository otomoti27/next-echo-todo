'use server'

import { cookies } from 'next/headers'
import { FormSchema, FormType } from './schema'
import { postSignin } from '@/features/auth'

export async function signin(data: FormType) {
  const parseData = FormSchema.safeParse(data)

  if (!parseData.success) {
    return {
      success: false,
      errors: parseData.error.format(),
    }
  }

  try {
    const response = await postSignin({
      email: parseData.data.email,
      password: parseData.data.password,
    })
    const token = response.data.token
    cookies().set('token', token, {
      maxAge: 60 * 60 * 24 * 30, // 30 days
      secure: true,
      sameSite: 'none',
      httpOnly: true,
      path: '/',
    })
    console.log(response.data)
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
