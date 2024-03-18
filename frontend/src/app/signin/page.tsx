import { Metadata } from 'next'
import Link from 'next/link'
import { SignInForm } from './SignInForm'

export const metadata: Metadata = {
  title: 'サインイン',
}

export default function SignUp() {
  return (
    <>
      <SignInForm />
      <div className='mt-8 space-y-2'>
        <p>
          ご登録がまだの方は
          <Link href='/signup' className='underline'>
            こちら
          </Link>
        </p>
      </div>
    </>
  )
}
