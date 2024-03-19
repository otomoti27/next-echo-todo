import { NextRequest, NextResponse } from 'next/server'
// import { getMe } from './features/auth'
// import { ApiError } from './lib/error'

export async function middleware(request: NextRequest) {
  if (process.env.NODE_ENV === 'development') {
    process.env['NODE_TLS_REJECT_UNAUTHORIZED'] = '0'
  }

  // 認証チェック
  if (request.nextUrl.pathname === '/') {
    return NextResponse.next()
  }
  const token = request.cookies.get('token')
  if (token && !request.nextUrl.pathname.startsWith('/u')) {
    return NextResponse.redirect(new URL('/u/todos', request.url))
  }

  if (
    request.nextUrl.pathname === '/signin' ||
    request.nextUrl.pathname === '/signup'
  ) {
    return NextResponse.next()
  }

  if (!token) {
    return NextResponse.redirect(new URL('/signin', request.url))
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico|static).*)'],
}

// const verifyToken = async () => {
//   try {
//     const user = await getMe('no-store')
//     console.log('verifyToken', user)
//     return true
//   } catch (error) {
//     if (error instanceof ApiError) {
//       console.error(error.message)
//       return false
//     }
//   }
// }
