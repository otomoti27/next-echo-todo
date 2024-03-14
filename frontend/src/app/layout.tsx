import type { Metadata } from 'next'
import { Inter as FontSans } from 'next/font/google'
import './globals.css'
import { CSRFToken } from './CSRFToken'
import { Toaster } from '@/components/ui/toaster'
import { cn } from '@/lib/utils'

process.env['NODE_TLS_REJECT_UNAUTHORIZED'] = '0'

const fontSans = FontSans({
  subsets: ['latin'],
  variable: '--font-sans',
})

export const metadata: Metadata = {
  title: {
    default: 'next-echo-todo',
    template: '%s | next-echo-todo',
  },
  description: 'Next.js + echoの学習用Todoアプリ',
}

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='ja'>
      <body
        className={cn(
          'min-h-screen bg-background font-sans antialiased',
          fontSans.variable,
        )}
      >
        <div className='mx-auto max-w-screen-lg px-4'>{children}</div>
        <Toaster />
        <CSRFToken />
      </body>
    </html>
  )
}
