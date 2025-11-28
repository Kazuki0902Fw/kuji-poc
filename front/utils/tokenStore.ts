import dayjs from "dayjs"

export type Token = {
  accessToken: string
  accessTokenExpiresAt: string
  refreshToken: string
  refreshTokenExpiresAt: string
}

const localStorageAuthTokenKey = "AuthToken"

export const getTokenFromLocalStorage = () => {
  if (typeof window === 'undefined' || !window.localStorage) {
    return null
  }
  const tokenStr = localStorage.getItem(localStorageAuthTokenKey)
  if (!tokenStr) {
    return null
  }
  return JSON.parse(tokenStr) as Token
}

export const setTokenToLocalStorage = (token: Token) => {
  if (typeof window === 'undefined' || !window.localStorage) {
    return
  }
  localStorage.setItem(localStorageAuthTokenKey, JSON.stringify(token))
}

export const clearTokenFromLocalStorage = () => {
  if (typeof window === 'undefined' || !window.localStorage) {
    return
  }
  localStorage.removeItem(localStorageAuthTokenKey)
}

// Backward compatibility aliases
export const getTokenFromSessionStorage = getTokenFromLocalStorage
export const setTokenToSessionStorage = setTokenToLocalStorage
export const clearTokenFromSessionStorage = clearTokenFromLocalStorage

export const isLoggedIn = () => {
  const token = getTokenFromLocalStorage()
  if (!token) {
    return false
  }

  const current = dayjs()
  const expiresAt = dayjs(token.refreshTokenExpiresAt)
  if (expiresAt.isBefore(current)) {
    return false
  }

  return true
}
