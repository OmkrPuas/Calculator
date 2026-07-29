import axios from 'axios'

export type CalculateRequest = {
  operation: string
  a: number
  b: number
}

const api = axios.create({
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
    if (err.response && err.response.data && err.response.data.error) {
      throw new Error(err.response.data.error)
    }
    if (err.code === 'ECONNABORTED') {
      throw new Error('request timeout')
    }
    throw new Error(err.message || 'network error')
  }
}
