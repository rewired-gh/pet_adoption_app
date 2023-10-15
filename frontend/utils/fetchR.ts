export const fetchR = async (apiPath: string, req: Object) => {
  let res = null
  try {
    res = await fetch(`${conf.baseApi}${apiPath}`, {
      method: 'POST',
      body: JSON.stringify(req),
    })
  } catch (e) {}
  if (res && res.ok) {
    return await res.json()
  }
  return null
}
