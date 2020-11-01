import callApi from '@/helpers/api-wrapper'

export const fetchUser = async () => {
  const response = await callApi({
    method: 'get',
    url: 'profile'
  })
  return response
}
