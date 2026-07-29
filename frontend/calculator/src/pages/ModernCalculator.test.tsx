import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import ModernCalculator from './ModernCalculator'

describe('ModernCalculator', () => {
  it('performs exponent operation and displays result', async () => {
    const user = userEvent.setup()
    render(<ModernCalculator />)

    await user.click(screen.getByRole('button', { name: '2' }))
    await user.click(screen.getByRole('button', { name: '^' }))
    await user.click(screen.getByRole('button', { name: '2' }))
    await user.click(screen.getByRole('button', { name: '=' }))

    expect(await screen.findByRole('status')).toHaveTextContent('4')
  })

  it('shows ERR for invalid square root value', async () => {
    const user = userEvent.setup()
    render(<ModernCalculator />)

    await user.click(screen.getByRole('button', { name: '-' }))
    await user.click(screen.getByRole('button', { name: '1' }))
    await user.click(screen.getByRole('button', { name: '=' }))
    await user.click(screen.getByRole('button', { name: '√' }))

    expect(await screen.findByRole('status')).toHaveTextContent('ERR')
  })
})
