import { z } from 'zod'

export const FormSchema = z.object({
  email: z.string().email('メールアドレスを入力してください'),
  password: z.string().min(1, 'パスワードを入力してください'),
})

export type FormType = z.infer<typeof FormSchema>
