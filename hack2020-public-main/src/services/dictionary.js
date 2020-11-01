import callApi from '@/helpers/api-wrapper'

export const fetchDictionary = async (keys) => {
  const data = JSON.stringify({
    request: keys
  })
  const response = await callApi({
    method: 'post',
    url: 'get_voc',
    data
  })
  return response
}

export const updateDictionary = async (login, name, voc) => {
  const data = JSON.stringify({
    login,
    Voc: [
      {
        Login: login,
        Name: name,
        Voc: voc
      }
    ]
  })
  const response = await callApi({
    method: 'post',
    url: 'set_voc',
    data
  })
  return response
}
