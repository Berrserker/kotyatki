import axios from 'axios'
import to from 'await-to-js'

import { API_BASE_URL, AUTH_TOKEN_NAME } from '@/config'

const fetcher = axios.create({})

export default async ({ method, url, data, params, remote = false, authRequired = true, formData = false }) => {
  const authToken = localStorage.getItem(AUTH_TOKEN_NAME) || ''
  const response = {
    data: null,
    errors: null
  }
  const requestConfig = {
    headers: {},
    method,
    url: remote ? url : `${API_BASE_URL}${url}`
  }
  if (authRequired && authToken) requestConfig.headers.Authorization = `Bearer ${authToken}`
  if (formData) requestConfig.headers['Content-Type'] = 'multipart/form-data'
  if (data) requestConfig.data = data
  if (params) requestConfig.params = params
  const [reqErrors, reqData] = await to(fetcher(requestConfig))
  if (reqData && reqData.data !== null) response.data = reqData.data
  if (reqErrors && reqErrors.response) response.errors = reqErrors.response.data
  return response
}
