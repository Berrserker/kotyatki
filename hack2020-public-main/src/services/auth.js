import callApi from '@/helpers/api-wrapper'

import { AUTH_TOKEN_NAME } from '@/config'

export const logIn = async (data) => {
  const response = await callApi({
    method: 'post',
    url: 'auth',
    data,
    authRequired: false
  })
  if (response.data && response.data.account && response.data.account.token) {
    localStorage.setItem(AUTH_TOKEN_NAME, response.data.account.token)
  }
  return mapLogInResponse(response)
}

export const logOut = () => {
  localStorage.setItem(AUTH_TOKEN_NAME, '')
}

const mapLogInResponse = (resp) => {
  if (!resp.data) return resp
  let data = null
  let errors = null
  if (resp.data && resp.data.status === 1) {
    errors = {
      message: resp.data.message
    }
  } else {
    data = {
      message: resp.data.message
    }
  }
  const response = {
    data,
    errors
  }
  return response
}
