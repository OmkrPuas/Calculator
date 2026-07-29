import { api, calculate } from './calculator'

describe('calculator api', () => {
  it('returns result when response includes result', async () => {
    vi.spyOn(api, 'post').mockResolvedValue({ data: { result: 10 } } as any)

    const res = await calculate({ operation: 'add', a: 5, b: 5 })
    expect(res).toBe(10)
  })

  it('throws error when server returns error', async () => {
    vi.spyOn(api, 'post').mockRejectedValue({
      isAxiosError: true,
      response: { data: { error: 'invalid' } },
    } as any)

    await expect(calculate({ operation: 'add', a: 3, b: 3 })).rejects.toThrow('invalid')
  })

  it('throws backend unavailable when request could not reach backend', async () => {
    vi.spyOn(api, 'post').mockRejectedValue({
      isAxiosError: true,
      request: {},
    } as any)

    await expect(calculate({ operation: 'add', a: 3, b: 3 })).rejects.toThrow('backend unavailable')
  })
})
