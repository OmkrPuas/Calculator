import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import TraditionalCalculator from './TraditionalCalculator'

describe('TraditionalCalculator', () => {
  it('handles basic input and addition', async () => {
    const user = userEvent.setup()
    render(<TraditionalCalculator />)

    await user.click(screen.getByRole('button', { name: '1' }))
    await user.click(screen.getByRole('button', { name: '+' }))
    await user.click(screen.getByRole('button', { name: '2' }))
    await user.click(screen.getByRole('button', { name: '=' }))

    expect(await screen.findByRole('status')).toHaveTextContent('3')
  })

  it('shows ERR on invalid sqrt input', async () => {
    const user = userEvent.setup()
    render(<TraditionalCalculator />)

    await user.click(screen.getByRole('button', { name: '-' }))
    await user.click(screen.getByRole('button', { name: '1' }))
    await user.click(screen.getByRole('button', { name: '=' }))
    await user.click(screen.getByRole('button', { name: '√' }))

    expect(await screen.findByRole('status')).toHaveTextContent('ERR')
  })

  it('calculates percentage correctly', async () => {
    const user = userEvent.setup()
    render(<TraditionalCalculator />)

    await user.click(screen.getByRole('button', { name: '5' }))
    await user.click(screen.getByRole('button', { name: '0' }))
    await user.click(screen.getByRole('button', { name: '%' }))

    expect(await screen.findByRole('status')).toHaveTextContent('0.5')
  })
})
