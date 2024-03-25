'use client'

import { ExclamationTriangleIcon } from '@radix-ui/react-icons'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
// eslint-disable-next-line no-restricted-imports
import { logout } from '@/features/auth/actions/logout'

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string }
  reset: () => void
}) {
  const [alert, setAlert] = useState<{
    title: string
    description: string
    onClick: () => void
    buttonText: string
  }>({
    title: '',
    description: '',
    onClick: () => {},
    buttonText: '',
  })
  const router = useRouter()

  useEffect(() => {
    console.error(error)
    // カスタムエラーが受け取れないため、messageでエラーの種類を判別
    switch (error.message) {
      case '401':
      case '403':
        setAlert({
          title: '認証エラー',
          description: '認証に失敗しました。再度ログインしてください。',
          onClick: async () => {
            await logout()
            router.push('/signin')
          },
          buttonText: 'ログイン',
        })
        break
      default:
        setAlert({
          title: 'エラー',
          description: 'エラーが発生しました。リロードしてください。',
          onClick: () => reset(),
          buttonText: 'リロード',
        })
        break
    }
  }, [error, reset, router])

  return (
    <div>
      <Alert variant='destructive'>
        <ExclamationTriangleIcon className='size-4' />
        <AlertTitle>{alert.title}</AlertTitle>
        <AlertDescription>
          {alert.description}
          <Button onClick={alert.onClick} variant={'destructive'}>
            {alert.buttonText}
          </Button>
        </AlertDescription>
      </Alert>
    </div>
  )
}
