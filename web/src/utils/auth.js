// import Cookies from 'js-cookie'

const TokenKey = 'Authorization'

export function getToken () {
  // return Cookies.get(TokenKey)
  return sessionStorage.getItem(TokenKey)
}

export function setToken (token) {
  // return Cookies.set(TokenKey, token)
  return sessionStorage.setItem(TokenKey, token)
}

export function removeToken () {
  // return Cookies.remove(TokenKey)
  return sessionStorage.removeItem(TokenKey)
}
