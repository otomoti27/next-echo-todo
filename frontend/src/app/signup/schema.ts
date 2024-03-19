import { z } from 'zod'

export const UserSchema = z
  .object({
    name: z
      .string()
      .min(1, 'ユーザー名を入力してください')
      .max(32, 'ユーザー名は32文字以内で入力してください'),
    email: z.string().email('メールアドレスを入力してください'),
    password: z
      .string()
      .min(6, 'パスワードは6文字以上で入力してください')
      .max(32, 'パスワードは32文字以内で入力してください'),
    passwordConfirm: z.string().min(1, '確認用のパスワードを入力してください'),
  })
  .refine((data) => data.password === data.passwordConfirm, {
    message: 'パスワードが一致しません',
    path: ['passwordConfirm'],
  })
