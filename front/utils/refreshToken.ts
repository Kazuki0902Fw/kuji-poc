import dayjs from "dayjs"
import {
  getTokenFromSessionStorage,
  setTokenToSessionStorage,
} from "./tokenStore"

export const noTokenError = new Error(
  "No token found. needs authenticate by login",
)

export const refreshTokenExpiredError = new Error(
  "Refresh token expired. needs authenticate by login",
)

export const failedToRefreshTokenError = new Error(
  "Failed to refresh token. needs authenticate by login",
)

export const isNeedNewLoginError = (error: unknown) => {
  if (error instanceof Error) {
    return (
      error.message === noTokenError.message ||
      error.message === refreshTokenExpiredError.message ||
      error.message === failedToRefreshTokenError.message
    )
  }
  return false
}

const refreshToken = async (): Promise<Token> => {
  const token = getTokenFromSessionStorage()
  if (!token) {
    return Promise.reject(noTokenError)
  }

  try {
    const { data } = await useAsyncGql("refreshToken", {
      input: { refreshToken: token.refreshToken }
    }, {
      getCachedData: () => undefined,
    });

    const newToken = data.value.refreshAccessToken
    setTokenToSessionStorage(newToken)
    return newToken
  } catch (error) {
    console.error(`failed to refresh token. src error: ${error}`)
    return Promise.reject(failedToRefreshTokenError)
  }
}

export const fetchTokenWithRefreshIfNeeds = (): Promise<Token> => {
  const token = getTokenFromSessionStorage()
  if (!token) {
    return Promise.reject(noTokenError)
  }

  const current = dayjs()

  const refreshTokenExpiresAt = dayjs(token.refreshTokenExpiresAt)
  if (refreshTokenExpiresAt.isBefore(current)) {
    return Promise.reject(refreshTokenExpiredError)
  }

  const accessTokenExpiresAt = dayjs(token.accessTokenExpiresAt)
  if (accessTokenExpiresAt.isBefore(current)) {
    return refreshToken()
  }

  return Promise.resolve(token)
}
