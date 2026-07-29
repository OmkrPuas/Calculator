import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import * as api from '../api/calculator'
import Calculator from './Calculator'

describe('Calculator page', () => {
  it('renders form controls and calculates values', async () => {
    vi.spyOn(api, 'calculate').mockResolvedValue(5)

    render(<Calculator />)

    const valueAInput = screen.getByLabelText('value-a')
    const valueBInput = screen.getByLabelText('value-b')
    const operationSelect = screen.getByRole('combobox')
    const calculateButton = screen.getByRole('button', { name: /calculate/i })

    await userEvent.type(valueAInput, '2')
    await userEvent.type(valueBInput, '3')
    await userEvent.selectOptions(operationSelect, 'add')
    await userEvent.click(calculateButton)

    await waitFor(() => {
      expect(screen.getByText('Result: 5')).toBeInTheDocument()
    })
  })

  it('shows validation error when valueB is required but empty', async () => {
    render(<Calculator />)

    await userEvent.type(screen.getByLabelText('value-a'), '9')
    const calculateButton = screen.getByRole('button', { name: /calculate/i })
    await userEvent.click(calculateButton)

    expect(await screen.findByText(/Value B is required/i)).toBeInTheDocument()
  })

  it('shows backend unavailable message when api call fails due to server unreachability', async () => {
    vi.spyOn(api, 'calculate').mockRejectedValue(new Error('backend unavailable'))

    render(<Calculator />)

    await userEvent.type(screen.getByLabelText('value-a'), '4')
    await userEvent.type(screen.getByLabelText('value-b'), '2')
    const calculateButton = screen.getByRole('button', { name: /calculate/i })
    await userEvent.click(calculateButton)

    expect(await screen.findByText(/backend unavailable/i)).toBeInTheDocument()
  })
})
