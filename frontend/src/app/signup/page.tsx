import { Metadata } from 'next'
import Link from 'next/link'
import { SignUpForm } from './SignUpForm'

export const metadata: Metadata = {
  title: 'サインアップ',
}

export default function SignUp() {
  return (
    <>
      <SignUpForm />
      <div className='mt-8 space-y-2'>
        <p>
          ご登録済みの方は
          <Link href='/signin' className='underline'>
            こちら
          </Link>
        </p>
      </div>
    </>
  )
}
