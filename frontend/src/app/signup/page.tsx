import { Metadata } from 'next'
import { SignUpForm } from './SignUpForm'

export const metadata: Metadata = {
  title: 'Sign Up',
}

export default function SignUp() {
  return <SignUpForm />
}
