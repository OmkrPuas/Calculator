import axios from 'axios'

export type CalculateRequest = {
  operation: string
  a: number
  b: number
}

export const api = axios.create({
  baseURL: (import.meta.env.VITE_API_BASE_URL as string) || 'http://localhost:8080',
  timeout: 5000,
  headers: { 'Content-Type': 'application/json' },
})

export async function calculate(payload: CalculateRequest): Promise<number> {
  try {
    const res = await api.post('/calculate', payload)
    if (res.data && typeof res.data.result === 'number') {
      return res.data.result
    }
    throw new Error(res.data?.error || 'invalid response from server')
  } catch (err: any) {
    if (axios.isAxiosError(err)) {
      if (err.response?.data?.error) {
        throw new Error(err.response.data.error)
      }
      if (err.code === 'ECONNABORTED') {
        throw new Error('request timeout')
      }
      if (err.request) {
        throw new Error('backend unavailable')
      }
      throw new Error('network error')
    }
    throw new Error(err?.message || 'network error')
  }
}
